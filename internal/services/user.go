package services

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"happy-server/internal/models"
)

type (
	CreateUserData struct {
		Email     string `json:"email" binding:"required,email,max=255"`
		Password  string `json:"password" binding:"required,max=255"`
		FirstName string `json:"firstName" binding:"required,max=255"`
		LastName  string `json:"lastName" binding:"required,max=255"`
	}

	UpdateUserData struct {
		Email     string `json:"email" binding:"omitempty,email,max=255"`
		Password  string `json:"password" binding:"max=255"`
		FirstName string `json:"firstName" binding:"max=255"`
		LastName  string `json:"lastName" binding:"max=255"`
	}

	UserService struct {
		DB *gorm.DB
	}
)

// ReadUsers gets all users from database
func (us *UserService) ReadUsers() ([]models.User, error) {
	var users []models.User
	if err := us.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ReadUserById reads one user from database by id (pk)
func (us *UserService) ReadUserById(id models.ID) (*models.User, error) {
	var foundUser models.User
	if err := us.DB.First(&foundUser, id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &foundUser, nil
}

// ReadUserByEmail reads one user by email field
func (us *UserService) ReadUserByEmail(email string) (*models.User, error) {
	var foundUser models.User
	if err := us.DB.First(&foundUser, "email = ?", email).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &foundUser, nil
}

// CreateUser creates new user record in database
func (us *UserService) CreateUser(d CreateUserData) (*models.User, error) {
	if _, alreadyExistErr := us.ReadUserByEmail(d.Email); alreadyExistErr == nil {
		return nil, errors.New("user already exists")
	}
	newUser := models.User{
		Email:     d.Email,
		Password:  d.Password,
		FirstName: d.FirstName,
		LastName:  d.LastName,
	}
	if creationErr := us.DB.Create(&newUser).Error; creationErr != nil {
		return nil, creationErr
	}
	return &newUser, nil
}

// UpdateUser updates user record in database
func (us *UserService) UpdateUser(id models.ID, d UpdateUserData) (*models.User, error) {
	targetUser, err := us.ReadUserById(id)
	if err != nil {
		return nil, err
	}
	if updateErr := us.DB.Model(targetUser).Clauses(clause.Returning{}).Updates(d).Error; updateErr != nil {
		return nil, updateErr
	}
	return targetUser, nil
}

// DeleteUser deletes user record from database
func (us *UserService) DeleteUser(id models.ID) error {
	_, err := us.ReadUserById(id)
	if err != nil {
		return err
	}
	us.DB.Delete(&models.User{}, id)
	return nil
}
