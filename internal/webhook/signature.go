package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// Generate a new SHA256 signature
func NewSignature(payload Payload, secretKey string) (string, error) {
	p := bytes.Buffer{}
	err := json.NewEncoder(&p).Encode(payload)

	if err != nil {
		return "", err
	}

	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write(p.Bytes())
	signature := hash.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(signature), nil
}
