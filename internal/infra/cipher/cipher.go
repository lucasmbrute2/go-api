package cipher

import "golang.org/x/crypto/bcrypt"

type Cipher struct {
	Salts int
}

func NewCipher(salts int) *Cipher {
	return &Cipher{}
}

func (c *Cipher) Encrypt(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), c.Salts)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (c *Cipher) Compare(plainText string, cipherText string) (bool,error) {
	err := bcrypt.CompareHashAndPassword([]byte(cipherText), []byte(plainText))
	if err != nil {
		return false, err
	}

	return true, nil
}