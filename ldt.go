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
	"strconv"
	"strings"
	"time"
)

// Token Unecrypted raw token in the form of a map[string]string
type Token map[string]string

// New Creates new token using the custom fields and will add a 15 min expiration
func New(customFields map[string]string) (ts string) {
	to := Token{}

	expTime := time.Now()
	expTime = expTime.Add(15 * time.Minute)

	customFields["exp"] = strconv.FormatInt(expTime.Unix(), 10)

	to = customFields
	ts = to.encrypt()
	return
}

func (t Token) encrypt() (ts string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("LDT_SECRET")), 12)
	customFields, _ := json.Marshal(t)
	encryptedCustomFields := encrypt(customFields, os.Getenv("LDT_SECRET"))
	hashString := base64.StdEncoding.EncodeToString(hash)
	fieldsString := base64.StdEncoding.EncodeToString(encryptedCustomFields)
	ts = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s.%s", hashString, fieldsString)))
	return
}

func (t Token) isExpired() (isExpired bool, err error) {
	timeNow := time.Now()
	expiration, err := strconv.ParseInt(t.GetValue("exp"), 10, 64)
	if timeNow.After(time.Unix(expiration, 0)) {
		isExpired = true
		return
	}
	return
}

// GetValue Returns value from an unencypted token or if no value exists will return empty string
func (t Token) GetValue(key string) (value string) {
	for k, v := range t {
		if k == key {
			return v
		}
	}
	return
}

// RenewLdtTokenFromRequest renew the token received in a *http.Request by
// resetting the expiration field and return if the token was expired
func RenewLdtTokenFromRequest(req *http.Request) (isExpired bool, ts string, err error) {
	token := req.Header.Get("Authorization")
	return RenewLdtToken(token)
}

// RenewLdtToken renew the token from a string by
// resetting the expiration field and return if the token was expired
func RenewLdtToken(tokenString string) (isExpired bool, ts string, err error) {
	isExpired, to, err := GetLdtToken(tokenString)

	if err != nil {
		return
	}

	expTime := time.Now()
	expTime = expTime.Add(15 * time.Minute)

	to["exp"] = strconv.FormatInt(expTime.Unix(), 10)

	ts = to.encrypt()
	return
}

// GetLdtTokenFromRequest Returns a token map[string]string from the request and return whether it is expired
func GetLdtTokenFromRequest(req *http.Request) (isExpired bool, to Token, err error) {
	token := req.Header.Get("Authorization")
	return GetLdtToken(token)
}

// GetLdtTokenFromRequest Returns a token map[string]string from a string and return whether it is expired
func GetLdtToken(tokenString string) (isExpired bool, to Token, err error) {
	firstSplitToken := strings.Split(tokenString, " ")

	if len(firstSplitToken) == 2 {
		tokenString = firstSplitToken[1]
	}

	tokenDecoded, _ := base64.StdEncoding.DecodeString(tokenString)
	tokenString = string(tokenDecoded)

	splitToken := strings.Split(tokenString, ".")

	if len(splitToken) != 2 {
		err = errors.New("Unauthorized: Invalid Token")
		return
	}

	hash, _ := base64.StdEncoding.DecodeString(splitToken[0])

	err = bcrypt.CompareHashAndPassword(hash, []byte(os.Getenv("LDT_SECRET")))

	if err != nil {
		err = errors.New("Unauthorized: Invalid Token")
		return
	}

	encyptedBody, _ := base64.StdEncoding.DecodeString(splitToken[1])

	tokenBody := decrypt(encyptedBody, os.Getenv("LDT_SECRET"))
	err = json.Unmarshal(tokenBody, &to)

	if err != nil {
		return
	}

	isExpired, err = to.isExpired()

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
