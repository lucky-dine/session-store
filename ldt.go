package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"os"
	"strings"
)

type TokenObject map[string]string

func New(customFields map[string]string) (ts string) {
	to := TokenObject{}
	to = customFields
	ts = to.Encrypt()
	return
}

func (t *TokenObject) Encrypt() (ts string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("LDT_SECRET")), 12)
	customFields, _ := json.Marshal(t)

	encryptedCustomFields := encrypt(customFields, os.Getenv("LDT_SECRET"))
	ts = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s.%s", string(hash), string(encryptedCustomFields))))
	return
}

func GetLdtToken(req *http.Request) (to TokenObject, err error) {
	token := req.Header.Get("Authorization")
	firstSplitToken := strings.Split(token, " ")

	if len(firstSplitToken) == 2 {
		token = firstSplitToken[1]
	}

	tokenDecoded, _ := base64.StdEncoding.DecodeString(token)
	token = string(tokenDecoded)

	splitToken := strings.Split(token, ".")

	if len(splitToken) != 2 {
		err = errors.New("Unauthorized: Invalid Token")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(splitToken[0]), []byte(os.Getenv("LDT_SECRET")))

	if err != nil {
		err = errors.New("Unauthorized: Invalid Token")
		return
	}

	tokenBody := decrypt([]byte(splitToken[1]), os.Getenv("LDT_SECRET"))
	err = json.Unmarshal(tokenBody, &to)
	return
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
