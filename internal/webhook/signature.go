package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// Generate a new SHA256 signature for the given webhook payload.
//
// This must be a raw binary buffer due to inconsistencies across different platforms
// & languages, which may result in different signatures being generated.
func NewSignature(payload []byte, secretKey string) (string, error) {
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write(payload)
	signature := hash.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(signature), nil
}
