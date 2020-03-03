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

func TestGetLdtToken(t *testing.T) {
	request, err := http.NewRequest("POST", "www.luckydine.com", nil)
	request.Header.Add("Authorization", "Bearer SkRKaEpERXlKRWxHUld0U2VuRnZlVEV6TVVwcFdXZ3lRMU0zUVU4MVVXbGpSWE5EVTBSRmJGQjZiVGwyVTJGak9UZE1aR3RYWW1GVGJHNXQuM2FVVEdhY2NVdlZFc042Wlk5WVZNY0ZpYW8rMVpuQURBb2d2WVpwbFJNQUV5bkRwdEFXRVhxOVRsUHh1UVN4NWxaTGtuS0svN2pNMWJNb2NkTndBcGMycnozd2pldVp4ZXMxVlNKcWE1aFNRYzcxag==")

	isExp, token, err := GetLdtToken(request)

	if isExp {
		t.Fail()
	}

	if token["username"] != "bsmith" {
		t.Error(err)
	}

	fmt.Println(token, isExp)
}

func TestRenewLdtToken(t *testing.T) {
	request, err := http.NewRequest("POST", "www.luckydine.com", nil)
	request.Header.Add("Authorization", "Bearer SkRKaEpERXlKR05ZTm00MlJrdHRabkJZTDBwM1lrTXpNRnBCYkhWSFdteElSRTFyVWpsalVHRjJXa00wT0ZOTVYzQkROVEUyVjB0QlNtOXguWnhxYWtjUjkzWjkxenBHZjZTQWFrMDFrYkUxcHI5aDRRNHhxOGs3S3ZuVkxsSnZieHRlUVR2UVQyT2ZFd042NEt2UmR2aDBBM242WjdyLzVCN1kyR2tIK004MD0=")

	isExp, token, err := RenewLdtToken(request)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(isExp, token)
}
