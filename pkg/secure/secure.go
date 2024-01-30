package secure

type Secure interface {
	Encrypt(message string) ([]byte, string, error)
	Decrypt(key string, message string) (string, error)
}
