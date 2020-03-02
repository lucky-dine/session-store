package token

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("LDT_SECRET", "test key")
	code := m.Run()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	tokenString := New(map[string]string{"username": "bsmith"})

	fmt.Println(tokenString)
}

func TestGetLdtToken(t *testing.T) {
	request, err := http.NewRequest("POST", "www.luckydine.com", nil)
	request.Header.Add("Authorization", "Bearer JDJhJDEyJEVvT1BhMDZIUmk4Zjcya3BxS3lrR2VkNGdENlJydzFMRWtvSkt1VVhmTEdmMk1KdU16YURLLrY0HK0xmaiUGRlugk2pRmWcJ+RlTemD2Ow35Ytac8dHzy/xHfVY/aKKzW8jSTiaRgg=")

	token, err := GetLdtToken(request)

	if token["username"] != "bsmith" {
		t.Error(err)
	}

	fmt.Println(token["username"])
}
