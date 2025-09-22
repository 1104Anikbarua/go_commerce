package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"encoding/base64"
	"net/http"
	"strings"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		payloadToken := r.Header.Get("authorization")
		if payloadToken == "" {
			http.Error(w, "Unauthorized Access", http.StatusUnauthorized)
			return
		}
		accessTokenSplit := strings.Split(payloadToken, " ")
		tokenLength := len(accessTokenSplit)
		if tokenLength != 2 {
			http.Error(w, "Not Acceptable", http.StatusNotAcceptable)
			return
		}
		accessToken := accessTokenSplit[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Forbidden Request", http.StatusForbidden)
			return
		}
		tokenHeader := tokenParts[0]
		tokenPayload := tokenParts[1]
		previousSignature := tokenParts[2]

		message := tokenHeader + "." + tokenPayload
		byteArrMessage := []byte(message)
		oldSecret := []byte(config.GetConfig().JwtSecretKey)
		h := hmac.New(sha256.New, oldSecret)
		h.Write(byteArrMessage)
		hash := h.Sum(nil)
		currentSignature := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash)
		if previousSignature != currentSignature {
			http.Error(w, "Unauthorize Access", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
