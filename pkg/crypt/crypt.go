package crypt

type Cryptor interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

type NoopCryptor struct {}

func (n *NoopCryptor) Encrypt(in []byte) ([]byte, error) {
	return in, nil
}

func (n *NoopCryptor) Decrypt(in []byte) ([]byte, error) {
	return in, nil
}