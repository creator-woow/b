package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

type accessTokenClaims struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarUrl string `json:"avatar_url"`
	jwt.RegisteredClaims
}

func createAccessToken(claims accessTokenClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims{
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		AvatarUrl: claims.AvatarUrl,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   claims.Subject,
			ID:        uuid.New().String(),
		},
	})
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

func createRefreshToken(claims jwt.RegisteredClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   claims.Subject,
		ID:        uuid.New().String(),
	})
	tokenString, _ := token.SignedString(secret)
	return tokenString
}
