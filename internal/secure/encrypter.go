package secure

// Define the Encrypter interface
type Encrypter interface {
	Encrypt(plaintext []byte) (string, error)
	Decrypt(asciiCiphertext string) ([]byte, error)
}
