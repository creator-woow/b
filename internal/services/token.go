package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	AccessTokenExpireTime  = time.Hour * 24
	RefreshTokenExpireTime = time.Hour * 24 * 30
)

type accessTokenClaims struct {
	jwt.RegisteredClaims
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarUrl string `json:"avatar_url"`
}

type TokenService struct {
	DB        *gorm.DB
	JWTSecret string
}

func (ts *TokenService) CreateAccess(claims accessTokenClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims{
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		AvatarUrl: claims.AvatarUrl,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   claims.Subject,
			ID:        uuid.New().String(),
		},
	})
	tokenString, _ := token.SignedString(ts.JWTSecret)
	return tokenString
}

func (ts *TokenService) CreateRefresh(claims jwt.RegisteredClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireTime)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   claims.Subject,
		ID:        uuid.New().String(),
	})
	tokenString, _ := token.SignedString(ts.JWTSecret)
	return tokenString
}
