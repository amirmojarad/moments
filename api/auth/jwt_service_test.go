package auth

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setSecretKeyInEnv() {
	os.Setenv("SECRET_KEY", "secret")
}

func TestGenerateToken(t *testing.T) {
	setSecretKeyInEnv()
	jwtService := New()

	email := "test@email.com"

	token := jwtService.GenerateToken(email)

	jwtToken, err := jwtService.ValidateToken(token)
	if err != nil {
		t.Log(err)
	}

	claims := jwtService.GetMapClaims(jwtToken)
	assert.Equal(t, email, claims["email"])
}
