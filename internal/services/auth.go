package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"happy-server/internal/models"
	"happy-server/pkg/password"
)

type LoginData struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,max=255"`
}

type TokensPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthService struct {
	DB           *gorm.DB
	UserService  *UserService
	TokenService *TokenService
}

func (s *AuthService) Register(d CreateUserData) (*models.User, error) {
	hashedPassword, hashingErr := pass.GenerateHash(d.Password)
	if hashingErr != nil {
		return nil, hashingErr
	}
	d.Password = hashedPassword
	newUser, creationErr := s.UserService.CreateUser(d)
	if creationErr != nil {
		return nil, creationErr
	}
	return newUser, nil
}

func (s *AuthService) Login(d LoginData) (*TokensPair, error) {
	foundUser, notFoundErr := s.UserService.ReadUserByEmail(d.Email)
	if notFoundErr != nil {
		return nil, errors.New("wrong credentials")
	}
	if !pass.ComparePasswordAndHash(d.Password, foundUser.Password) {
		return nil, errors.New("wrong credentials")
	}
	sub := fmt.Sprintf("%v", foundUser.ID)
	accessToken := s.TokenService.CreateAccess(
		accessTokenClaims{
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			AvatarUrl: foundUser.AvatarUrl,
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: sub,
			},
		},
	)
	refreshToken := s.TokenService.CreateRefresh(
		jwt.RegisteredClaims{Subject: sub},
	)
	tokens := &TokensPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokens, nil
}
