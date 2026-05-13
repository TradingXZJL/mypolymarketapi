package mypolymarketapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

func HmacSha256(secret, data string) ([]byte, error) {
	key, err := base64.URLEncoding.DecodeString(secret)
	if err != nil {
		return nil, fmt.Errorf("failed to decode secret: %v", err)
	}
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil), nil
}

func signL2(secret, method, timestamp string, url url.URL, reqBody []byte) (string, error) {
	path := url.Path
	hmacSha256Data := timestamp + method + path
	if reqBody != nil {
		hmacSha256Data += string(reqBody)
	}

	hmacBytes, err := HmacSha256(secret, hmacSha256Data)
	if err != nil {
		return "", err
	}
	sign := base64.URLEncoding.EncodeToString(hmacBytes)

	// log.Warn(hmacSha256Data)
	// log.Warn("timestamp: ", timestamp)
	// log.Warn("method: ", method)
	// log.Warn("requestPath: ", path)
	// log.Warn("reqBody: ", string(reqBody))
	// log.Warn("hmacSha256Data: ", hmacSha256Data)
	// log.Warn("sign: ", sign)

	return sign, nil
}
