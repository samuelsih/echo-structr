package jwt

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	SetSecretKey("SUPER_SECRET_BLABLABLA")

	code := m.Run()
	os.Exit(code)
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	claims := Claims{
		"id":   "132123123",
		"name": "idk",
	}

	_, err := Generate(claims, time.Now().Add(150000*time.Hour).Unix())

	if err != nil {
		t.Fatalf("err should not nil: %v", err)
	}
}

func TestVerify(t *testing.T) {
	t.Parallel()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjIyMTYwNDE0NjksImlkIjoiMTMyMTIzMTIzIiwibmFtZSI6ImlkayJ9.VEDHvf4_-JhshVaTAZPlMfXVwRlVN2pWooXoOLIyZHc"

	claims, err := Verify(token)
	if err != nil {
		t.Fatalf("err should not nil: %v", err)
	}

	if _, ok := claims["id"]; !ok {
		t.Fatalf("no id in payload")
	}

	if claims["id"] != "132123123" {
		t.Fatalf("wrong id in payload, expected 132123123 got %s", claims["id"])
	}

	if _, ok := claims["name"]; !ok {
		t.Fatalf("no name in payload")
	}

	if claims["name"] != "idk" {
		t.Fatalf("wrong name in payload, expected idk got %s", claims["name"])
	}
}
