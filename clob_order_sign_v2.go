package mypolymarketapi

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// 以下 EIP-712 定义与 Polymarket clob-client-v2 保持一致：
// - src/order-utils/model/ctfExchangeV2TypedData.ts（域名字段、Order 成员顺序与类型）
// - src/order-utils/exchangeOrderBuilderV2.ts（message 中 side / signatureType 编码）
// - src/config.ts（Polygon / Amoy 上 exchangeV2 / negRiskExchangeV2）

const (
	ctfExchangeV2DomainName    = "Polymarket CTF Exchange"
	ctfExchangeV2DomainVersion = "2"
)

// CLOBOrderV2Signer 由 *CLOBOrderReq 实现：基于当前载荷做 CLOB V2 EIP-712 签名并写入 Signature。
type CLOBOrderV2Signer interface {
	SignCLOBOrderV2EIP712(privateKey *ecdsa.PrivateKey, chainID ChainType, negRisk bool) error
}

var _ CLOBOrderV2Signer = (*CLOBOrderReq)(nil)

// Polygon / Amoy CTF Exchange V2 合约地址（与 clob-client-v2 src/config.ts 一致）。
var (
	polygonExchangeV2        = common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B")
	polygonNegRiskExchangeV2 = common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59")
	amoyExchangeV2           = common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B")
	amoyNegRiskExchangeV2    = common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59")
)

// VerifyingContractCLOBV2 返回用于 EIP-712 domain.verifyingContract 的合约地址（chainId + negRisk）。
func VerifyingContractCLOBV2(chainID int64, negRisk bool) (common.Address, error) {
	switch chainID {
	case 137:
		if negRisk {
			return polygonNegRiskExchangeV2, nil
		}
		return polygonExchangeV2, nil
	case 80002:
		if negRisk {
			return amoyNegRiskExchangeV2, nil
		}
		return amoyExchangeV2, nil
	default:
		return common.Address{}, fmt.Errorf("unsupported chainId for CLOB V2: %d (use 137 or 80002)", chainID)
	}
}

// SignCLOBOrderV2EIP712 使用私钥对当前订单载荷做 CLOB V2 EIP-712 签名，结果写入 Signature。
// 必填：Maker、Signer、TokenID、MakerAmount、TakerAmount、Side、Timestamp、SignatureType。
// Salt 若为 nil，会自动生成并写回（与 SDK 行为类似，数值落在 int64 可表示范围内以便 JSON 序列化）。
// Builder / Metadata 空或全零 hex 时按 EIP-712 使用 bytes32(0)；与 TS metadata ?? bytes32Zero 一致。
// privateKey 必须对应 Signer 链上地址。
func (o *CLOBOrderReq) SignCLOBOrderV2EIP712(privateKey *ecdsa.PrivateKey, chainID ChainType, negRisk bool) error {
	if o == nil {
		return errors.New("CLOBOrderReq is nil")
	}
	if privateKey == nil {
		return errors.New("privateKey is nil")
	}
	if o.Maker == nil || o.Signer == nil || o.TokenID == nil || o.MakerAmount == nil || o.TakerAmount == nil ||
		o.Side == nil || o.Timestamp == nil || o.SignatureType == nil {
		return errors.New("CLOBOrderReq: missing field required for V2 signing (maker, signer, tokenId, makerAmount, takerAmount, side, timestamp, signatureType)")
	}
	maker := common.HexToAddress(strings.TrimSpace(*o.Maker))
	signer := common.HexToAddress(strings.TrimSpace(*o.Signer))
	signerAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	if signerAddr != signer {
		return fmt.Errorf("privateKey does not match signer: derived %s, want %s", signerAddr.Hex(), signer.Hex())
	}
	tokenID, ok := new(big.Int).SetString(strings.TrimSpace(*o.TokenID), 10)
	if !ok {
		return errors.New("tokenId: invalid decimal uint256")
	}
	makerAmt, ok := new(big.Int).SetString(strings.TrimSpace(*o.MakerAmount), 10)
	if !ok {
		return errors.New("makerAmount: invalid decimal uint256")
	}
	takerAmt, ok := new(big.Int).SetString(strings.TrimSpace(*o.TakerAmount), 10)
	if !ok {
		return errors.New("takerAmount: invalid decimal uint256")
	}
	sideStr := strings.TrimSpace(*o.Side)
	var sideBuy bool
	switch strings.ToUpper(sideStr) {
	case "BUY":
		sideBuy = true
	case "SELL":
		sideBuy = false
	default:
		return fmt.Errorf("side: want BUY or SELL, got %q", sideStr)
	}
	tsMs, err := strconv.ParseInt(strings.TrimSpace(*o.Timestamp), 10, 64)
	if err != nil {
		return fmt.Errorf("timestamp: %w", err)
	}
	st := *o.SignatureType
	if st < 0 || st > 3 {
		return fmt.Errorf("signatureType: out of range 0..3, got %d", st)
	}
	var salt *big.Int
	if o.Salt == nil {
		saltI64, genErr := generateCLOBOrderSaltV2Int64()
		if genErr != nil {
			return genErr
		}
		o.Salt = GetPointer(saltI64)
	}
	salt = big.NewInt(*o.Salt)

	metaBytes, err := clobOrderMetadataBuilderBytes32(o.Metadata, o.Builder)
	if err != nil {
		return err
	}

	sig, err := signCLOBOrderV2EIP712Core(
		privateKey,
		chainID,
		negRisk,
		salt,
		maker,
		signer,
		tokenID,
		makerAmt,
		takerAmt,
		sideBuy,
		uint8(st),
		tsMs,
		metaBytes.metadata,
		metaBytes.builder,
	)
	if err != nil {
		return err
	}
	o.Signature = GetPointer(sig)
	return nil
}

type metadataBuilderBytes32 struct {
	metadata [32]byte
	builder  [32]byte
}

func clobOrderMetadataBuilderBytes32(metadata, builder *string) (metadataBuilderBytes32, error) {
	var out metadataBuilderBytes32
	mStr := ""
	if metadata != nil {
		mStr = strings.TrimSpace(*metadata)
	}
	bStr := ""
	if builder != nil {
		bStr = strings.TrimSpace(*builder)
	}
	md, err := parseBytes32WireOrZero(mStr)
	if err != nil {
		return out, fmt.Errorf("metadata: %w", err)
	}
	bd, err := parseBytes32WireOrZero(bStr)
	if err != nil {
		return out, fmt.Errorf("builder: %w", err)
	}
	out.metadata = md
	out.builder = bd
	return out, nil
}

// parseBytes32WireOrZero：Wire 上 metadata 常为 ""；builder 常为 0x..64；空串按 bytes32(0) 参与 EIP-712。
func parseBytes32WireOrZero(s string) ([32]byte, error) {
	if s == "" {
		return [32]byte{}, nil
	}
	return ParseBytes32Hex(s)
}

// GenerateCLOBOrderSaltV2Int64 生成 order.salt（int64，与 JSON integer 及常见 SDK 数量级兼容）。
func GenerateCLOBOrderSaltV2Int64() (int64, error) {
	return generateCLOBOrderSaltV2Int64()
}

func generateCLOBOrderSaltV2Int64() (int64, error) {
	var rbuf [8]byte
	if _, err := rand.Read(rbuf[:]); err != nil {
		return 0, err
	}
	u := binary.BigEndian.Uint64(rbuf[:])
	// 近似 TS：Math.round(Math.random() * Date.now())
	x := time.Now().UnixMilli()
	return int64(u%1000000) * x, nil
}

// ParseBytes32Hex 解析 0x + 64 hex 的 bytes32；空串视为全零。
func ParseBytes32Hex(s string) ([32]byte, error) {
	var out [32]byte
	s = strings.TrimSpace(s)
	if s == "" {
		return out, nil
	}
	s = strings.TrimPrefix(strings.ToLower(s), "0x")
	if len(s) != 64 {
		return out, errors.New("bytes32 must be 64 hex chars")
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return out, err
	}
	copy(out[:], b)
	return out, nil
}

func signCLOBOrderV2EIP712Core(
	privateKey *ecdsa.PrivateKey,
	chainID ChainType,
	negRisk bool,
	salt *big.Int,
	maker, signer common.Address,
	tokenID, makerAmount, takerAmount *big.Int,
	sideBuy bool,
	signatureType uint8,
	timestampMs int64,
	metadata, builder [32]byte,
) (string, error) {
	vContract, err := VerifyingContractCLOBV2(chainID.Int64(), negRisk)
	if err != nil {
		return "", err
	}
	side := big.NewInt(0)
	if !sideBuy {
		side = big.NewInt(1)
	}
	sigType := new(big.Int).SetUint64(uint64(signatureType))
	tsStr := strconv.FormatInt(timestampMs, 10)

	td := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Order": {
				{Name: "salt", Type: "uint256"},
				{Name: "maker", Type: "address"},
				{Name: "signer", Type: "address"},
				{Name: "tokenId", Type: "uint256"},
				{Name: "makerAmount", Type: "uint256"},
				{Name: "takerAmount", Type: "uint256"},
				{Name: "side", Type: "uint8"},
				{Name: "signatureType", Type: "uint8"},
				{Name: "timestamp", Type: "uint256"},
				{Name: "metadata", Type: "bytes32"},
				{Name: "builder", Type: "bytes32"},
			},
		},
		PrimaryType: "Order",
		Domain: apitypes.TypedDataDomain{
			Name:              ctfExchangeV2DomainName,
			Version:           ctfExchangeV2DomainVersion,
			ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
			VerifyingContract: vContract.Hex(),
		},
		Message: apitypes.TypedDataMessage{
			"salt":          salt,
			// go-ethereum apitypes 的 address 分支只接受 string / []byte(len 20) / [20]byte，
			// 不接受命名类型 common.Address（与 [20]byte 在 type-switch 下不等价），否则报 dataMismatchError。
			"maker":         maker.Hex(),
			"signer":        signer.Hex(),
			"tokenId":       tokenID,
			"makerAmount":   makerAmount,
			"takerAmount":   takerAmount,
			"side":          side,
			"signatureType": sigType,
			"timestamp":     uint256FromDecimalString(tsStr),
			"metadata":      metadata[:],
			"builder":       builder[:],
		},
	}
	digest, _, err := apitypes.TypedDataAndHash(td)
	if err != nil {
		return "", err
	}
	sig, err := crypto.Sign(digest, privateKey)
	if err != nil {
		return "", err
	}
	if sig[64] < 27 {
		sig[64] += 27
	}
	return "0x" + hex.EncodeToString(sig), nil
}

func uint256FromDecimalString(s string) *big.Int {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return big.NewInt(0)
	}
	return i
}
