package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type TSHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type TSPayload struct {
	Sub          int    `json:"_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	IsShopOwner  bool   `json:"is_shop_owner"`
	JwtSecretKey string `json:"jwt_secret_key"`
}

func CreateJwt(payload TSPayload, secret string) (string, error) {
	header := TSHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerBase64 := CreateBase64UrlEncode(byteArrHeader)
	byteArrPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadBase64 := CreateBase64UrlEncode(byteArrPayload)

	message := headerBase64 + "." + payloadBase64
	byteArrMessage := []byte(message)
	byteArrSecret := []byte(secret)
	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	signature := h.Sum(nil)
	signatureBase64 := CreateBase64UrlEncode(signature)
	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64
	return jwt, nil
}

func CreateBase64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)

}
