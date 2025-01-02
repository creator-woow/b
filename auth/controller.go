package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"happy-server/user"
)

// todo: move somewhere
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func compareTextAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(db *gorm.DB, d user.CreateUserData) (*user.User, error) {
	hashedPassword, hashingErr := hashPassword(d.Password)
	if hashingErr != nil {
		return nil, hashingErr
	}
	d.Password = hashedPassword
	newUser, creationErr := user.CreateUser(db, d)
	if creationErr != nil {
		return nil, creationErr
	}
	return newUser, nil
}

func Login(db *gorm.DB, d loginData) (*TokensPair, error) {
	foundUser, notFoundErr := user.ReadUserByEmail(db, d.Email)
	if notFoundErr != nil {
		return nil, errors.New("wrong credentials")
	}
	if !compareTextAndHash(d.Password, foundUser.Password) {
		return nil, errors.New("wrong credentials")
	}
	sub := fmt.Sprintf("%v", foundUser.ID)
	accessToken := createAccessToken(
		accessTokenClaims{
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			AvatarUrl: foundUser.AvatarUrl,
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: sub,
			},
		},
	)
	refreshToken := createRefreshToken(
		jwt.RegisteredClaims{Subject: sub},
	)
	tokens := &TokensPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokens, nil
}
