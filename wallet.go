package mypolymarketapi

import (
	"crypto/ecdsa"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type WalletType string

const (
	WalletTypeEOA         WalletType = "EOA"
	WalletTypePOLY_PROXY  WalletType = "POLY_PROXY"
	WalletTypeGNOSIS_SAFE WalletType = "GNOSIS_SAFE"
)

var WalletTypeMap = map[WalletType]string{
	WalletTypeEOA:         "0",
	WalletTypePOLY_PROXY:  "1",
	WalletTypeGNOSIS_SAFE: "2",
}

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Signer     common.Address // 钱包地址 (Signer)
	Type       WalletType     // 钱包类型：EOA、POLY_PROXY、GNOSIS_SAFE
}

func NewWallet(privateKeyHex string) (*Wallet, error) {
	// 1. 解析私钥
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	// 2. 从私钥派生 Signer 地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key")
	}
	signer := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Wallet{
		PrivateKey: privateKey,
		Signer:     signer,
	}, nil
}

type ApiKeyCreds struct {
	APIKey     string
	Secret     string
	Passphrase string
}

func (w *Wallet) createL1Headers(nonce uint64) (map[string]string, error) {
	ts := time.Now().Unix()
	signature, err := signClobAuthEIP712(w.PrivateKey, w.Signer, POLYGON_MAINNET_CHAIN_ID, ts, nonce)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"POLY_ADDRESS":   w.Signer.String(),
		"POLY_SIGNATURE": signature,
		"POLY_TIMESTAMP": strconv.FormatInt(ts, 10),
		"POLY_NONCE":     strconv.FormatUint(nonce, 10),
	}, nil
}

// nonce string 用于保证APIKey的唯一性，如果所填 nonce 之前已经使用过，则之前生成的凭证会自动失效并创建新的凭证
func (w *Wallet) CreateApiKey(nonce uint64) (*ApiKeyCreds, error) {
	headers, err := w.createL1Headers(nonce)
	if err != nil {
		return nil, err
	}
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBCreateApiKey])
	res, err := pmCallAPIWithHeaders[CLOBApiCredentialsRes](url, NIL_REQBODY, POST, headers)
	if err != nil {
		return nil, err
	}
	apiKeyCreds := &ApiKeyCreds{
		APIKey:     res.Data.APIKey,
		Secret:     res.Data.Secret,
		Passphrase: res.Data.Passphrase,
	}
	return apiKeyCreds, nil
}

// nonce string 用于找回之前的APIKey，如果所填 nonce 之前没有使用过，则返回空字符串
func (w *Wallet) DeriveApiKey(nonce uint64) (*ApiKeyCreds, error) {
	headers, err := w.createL1Headers(nonce)
	if err != nil {
		return nil, err
	}
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, CLOB_REST, CLOBAPITypeMap[CLOBDeriveApiKey])
	res, err := pmCallAPIWithHeaders[CLOBApiCredentialsRes](url, NIL_REQBODY, GET, headers)
	if err != nil {
		return nil, err
	}
	apiKeyCreds := &ApiKeyCreds{
		APIKey:     res.Data.APIKey,
		Secret:     res.Data.Secret,
		Passphrase: res.Data.Passphrase,
	}
	return apiKeyCreds, nil
}
