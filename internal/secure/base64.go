package secure

import (
	"encoding/base64"
)

type Base64 struct{}

func (e *Base64) Encrypt(data []byte) (string, error) {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}

func (e *Base64) Decrypt(data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
