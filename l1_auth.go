package mypolymarketapi

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// clobAuth712Message 为 CLOB L1 EIP-712（ClobAuth）固定文案，与官方文档一致。
const clobAuth712Message = "This message attests that I control the given wallet"

// signClobAuthEIP712 对 Polymarket CLOB L1 的 ClobAuth 结构做 EIP-712 签名，返回 0x 前缀十六进制签名（供 POLY_SIGNATURE header）。
func signClobAuthEIP712(privateKey *ecdsa.PrivateKey, signer common.Address, chainID ChainType, timestampUnix int64, nonce uint64) (string, error) {
	td := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
			"ClobAuth": {
				{Name: "address", Type: "address"},
				{Name: "timestamp", Type: "string"},
				{Name: "nonce", Type: "uint256"},
				{Name: "message", Type: "string"},
			},
		},
		PrimaryType: "ClobAuth",
		Domain: apitypes.TypedDataDomain{
			Name:    "ClobAuthDomain",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(chainID.Int64()),
		},
		Message: apitypes.TypedDataMessage{
			"address":   signer.Hex(),
			"timestamp": strconv.FormatInt(timestampUnix, 10),
			"nonce":     new(big.Int).SetUint64(nonce),
			"message":   clobAuth712Message,
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
