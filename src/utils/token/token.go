package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"github.com/gin-gonic/gin"

	gouuid "github.com/nu7hatch/gouuid"
)

var secret = []byte("mysecret90123456")

func New() cipher.Block {
	secretLen := len(secret)
	if secretLen != 16 && secretLen != 24 && secretLen != 32 {
		panic("Secret must be 16, 24 or 32 bytes long")
	}
	block, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}
	return block
}

func Encrypt(id string) (string, error) {

	u, err := gouuid.ParseHex(id)
	if err != nil {
		panic(err)
	}

	block := New()

	b := base64.StdEncoding.EncodeToString(u[:])
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	token := base64.StdEncoding.EncodeToString(ciphertext)
	return token, nil
}

func Decrypt(g *gin.Context) (string, error) {

	token := ExtractToken(g)
	block := New()

	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	if len(data) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	content, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return "", err
	}

	uuid, err := gouuid.Parse(content[:16])
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

func ExtractToken(g *gin.Context) string {
	token := g.Query("token")
	if token != "" {
		return token
	}
	bearerToken := g.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
