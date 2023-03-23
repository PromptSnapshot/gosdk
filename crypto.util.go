package gosdk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type CryptoUtils interface {
	Encrypt(secret []byte, content string) (encoded string, err error)
	Decrypt(secret []byte, secure string) (decoded string, err error)
	Hash(bv []byte) string
	Bcrypt(bv []byte) string
	CompareHashAndPassword(hash string, password string) error
}

func NewCryptoUtils() CryptoUtils {
	return &cryptoUtils{}
}

type cryptoUtils struct {
}

func (u *cryptoUtils) Encrypt(secret []byte, content string) (encoded string, err error) {
	plainText := []byte(content)

	block, err := aes.NewCipher(secret)
	if err != nil {
		return
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText), nil
}

func (u *cryptoUtils) Decrypt(secret []byte, secure string) (decoded string, err error) {
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)

	if err != nil {
		return
	}

	block, err := aes.NewCipher(secret)

	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("invalid ciphertext")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}

func (u *cryptoUtils) Hash(bv []byte) string {
	h := sha256.New()
	h.Write(bv)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (u *cryptoUtils) Bcrypt(bv []byte) string {
	hashed, _ := bcrypt.GenerateFromPassword(bv, bcrypt.MinCost)
	return string(hashed)
}

func (u *cryptoUtils) CompareHashAndPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
