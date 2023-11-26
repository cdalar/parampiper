package secure

type None struct{}

func (e *None) Encrypt(data []byte) (string, error) {
	return string(data), nil
}

func (e *None) Decrypt(data string) ([]byte, error) {
	return []byte(data), nil
}
