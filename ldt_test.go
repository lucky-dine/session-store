package token

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("LDT_SECRET", `ZAJ6PpUtEySe^=Yad-nhH-3%9xh6g78E#?xhZx?M7z5e*qqUA57!9g28rKn+K6gxH+P$5UAh#Zs%PZ?!%h=RPjGU6**at&UYdcFAV9j#N3!NJwh&j!!Au6aQ*Ag*6u8xkh+ehUBQ#6QTa7u!+tpvAXgk7xp%DDaa6jV-BXE&e^3U7y$PzAc*smQv#CXZ=KpyekwYmx2rU2qDh^DKsukzcJuLMq!EX*m-XmNeL!WwDnkpSs$+jt?K?ja-Nfxpz7a@`)
	code := m.Run()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	tokenString := New(map[string]string{"username": "bsmith", "admin": "false"})

	fmt.Println(tokenString)
}

func TestGetLdtTokenFromRequest(t *testing.T) {
	request, err := http.NewRequest("POST", "www.luckydine.com", nil)
	request.Header.Add("Authorization", "Bearer SkRKaEpERXlKSGd1WWtzdVdISkRObmhLTjFka2FtTmpWMUo2UlU5NFNsRnBNVVZ0TGpOaVNXYzVVak5yVW5ob1IwUjBjVWxLVUdkWWEwTXUuOVlzODhubDAzUlFPbG1PWTlyejNrdG1XbElyRlJXYk1kVjUyc2luSXl5YzNib2JpK2N5RWRKVG1HRWFzRTFKc3ViQ1BqbmdNblQzd1FSZm1tOUZTVm95ejYxaFNlL3pGWWg2N2EzNzRRTHRQUElSdA==")

	isExp, token, err := GetLdtTokenFromRequest(request)

	if token["username"] != "bsmith" {
		t.Error(err)
	}

	fmt.Println(token, isExp)
}

func TestGetLdtToken(t *testing.T) {
	tokenString := `SkRKaEpERXlKSGd1WWtzdVdISkRObmhLTjFka2FtTmpWMUo2UlU5NFNsRnBNVVZ0TGpOaVNXYzVVak5yVW5ob1IwUjBjVWxLVUdkWWEwTXUuOVlzODhubDAzUlFPbG1PWTlyejNrdG1XbElyRlJXYk1kVjUyc2luSXl5YzNib2JpK2N5RWRKVG1HRWFzRTFKc3ViQ1BqbmdNblQzd1FSZm1tOUZTVm95ejYxaFNlL3pGWWg2N2EzNzRRTHRQUElSdA==`

	isExp, token, err := GetLdtToken(tokenString)

	if token["username"] != "bsmith" {
		t.Error(err)
	}

	fmt.Println(token, isExp)
}

func TestRenewLdtTokenFromRequest(t *testing.T) {
	request, err := http.NewRequest("POST", "www.luckydine.com", nil)
	request.Header.Add("Authorization", "Bearer SkRKaEpERXlKSGd1WWtzdVdISkRObmhLTjFka2FtTmpWMUo2UlU5NFNsRnBNVVZ0TGpOaVNXYzVVak5yVW5ob1IwUjBjVWxLVUdkWWEwTXUuOVlzODhubDAzUlFPbG1PWTlyejNrdG1XbElyRlJXYk1kVjUyc2luSXl5YzNib2JpK2N5RWRKVG1HRWFzRTFKc3ViQ1BqbmdNblQzd1FSZm1tOUZTVm95ejYxaFNlL3pGWWg2N2EzNzRRTHRQUElSdA==")

	isExp, token, err := RenewLdtTokenFromRequest(request)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(isExp, token)
}

func TestRenewLdtToken(t *testing.T) {
	tokenString := `SkRKaEpERXlKSGd1WWtzdVdISkRObmhLTjFka2FtTmpWMUo2UlU5NFNsRnBNVVZ0TGpOaVNXYzVVak5yVW5ob1IwUjBjVWxLVUdkWWEwTXUuOVlzODhubDAzUlFPbG1PWTlyejNrdG1XbElyRlJXYk1kVjUyc2luSXl5YzNib2JpK2N5RWRKVG1HRWFzRTFKc3ViQ1BqbmdNblQzd1FSZm1tOUZTVm95ejYxaFNlL3pGWWg2N2EzNzRRTHRQUElSdA==`

	isExp, token, err := RenewLdtToken(tokenString)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(isExp, token)
}
