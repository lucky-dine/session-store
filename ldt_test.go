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
	request.Header.Add("Authorization", "Bearer SkRKaEpERXlKR05ZTm00MlJrdHRabkJZTDBwM1lrTXpNRnBCYkhWSFdteElSRTFyVWpsalVHRjJXa00wT0ZOTVYzQkROVEUyVjB0QlNtOXguWnhxYWtjUjkzWjkxenBHZjZTQWFrMDFrYkUxcHI5aDRRNHhxOGs3S3ZuVkxsSnZieHRlUVR2UVQyT2ZFd042NEt2UmR2aDBBM242WjdyLzVCN1kyR2tIK004MD0=")

	isExp, token, err := GetLdtToken(request)

	if isExp {
		t.Fail()
	}

	if token["username"] != "bsmith" {
		t.Error(err)
	}

	fmt.Println(token["username"], isExp)
}

func TestRenewLdtToken(t *testing.T) {
	request, err := http.NewRequest("POST", "www.luckydine.com", nil)
	request.Header.Add("Authorization", "Bearer SkRKaEpERXlKRmxNTjI1SGQwTnFORVoxYm5wNEwwWXZOa2RPUjNWSlFrY3VXRE54ZEVGa09XbEhXbUZrV2tVNVVIZEdOV0ZLVTBsQ2JrcEQuUGg1Q0w4OFk2cmRaTWN3YmJzaFFydWorSG9VdnVpNWZwblVLVGNjL0JyaGczYUVOMUE4RmZBZkVuQU9zTlJnYWZVbi9ya21GN09BcmtjU01vYzBvKytQSVpTRT0=")

	isExp, token, err := RenewLdtToken(request)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(isExp, token)
}
