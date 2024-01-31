package secure

type Secure interface {
	Encrypt(message []byte) (string, string, error)
	Decrypt(key []byte, message string) (string, error)
}
