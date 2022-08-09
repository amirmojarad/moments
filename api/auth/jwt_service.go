package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetMapClaims(token *jwt.Token) jwt.MapClaims
}

func New() JWTService {
	secretKey := os.Getenv("SECRET_KEY")
	return JwtService{
		secretKey: secretKey,
		issure:    "moments",
	}
}

type JwtService struct {
	secretKey string
	issure    string
}

type claims struct {
	Email string `json:"email"`
	jwt.Claims
}

func (service JwtService) GenerateToken(email string) string {
	claims := &claims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := t.SignedString([]byte(service.secretKey)); err != nil {
		log.Error().Err(err).Msg("error while generating token")
		return ""
	} else {
		return token
	}
}

func (service JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, service.tokenValidation)
}

func (service JwtService) tokenValidation(t *jwt.Token) (interface{}, error) {
	if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, fmt.Errorf("invalid token %+v", t.Header["alg"])
	}
	return []byte(service.secretKey), nil
}

func (service JwtService) GetMapClaims(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}
