package user

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"happy-server/lib"
)

// ReadUsers gets all users from database
func ReadUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ReadUserById reads one user from database by id (pk)
func ReadUserById(db *gorm.DB, id lib.ID) (*User, error) {
	var foundUser User
	if err := db.First(&foundUser, id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &foundUser, nil
}

// ReadUserByEmail reads one user by email field
func ReadUserByEmail(db *gorm.DB, email string) (*User, error) {
	var foundUser User
	if err := db.First(&foundUser, "email = ?", email).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &foundUser, nil
}

// CreateUser creates new user record in database
func CreateUser(db *gorm.DB, d CreateUserData) (*User, error) {
	if _, alreadyExistErr := ReadUserByEmail(db, d.Email); alreadyExistErr == nil {
		return nil, errors.New("user already exists")
	}
	newUser := User{
		Email:     d.Email,
		Password:  d.Password,
		FirstName: d.FirstName,
		LastName:  d.LastName,
	}
	if creationErr := db.Create(&newUser).Error; creationErr != nil {
		return nil, creationErr
	}
	return &newUser, nil
}

// UpdateUser updates user record in database
func UpdateUser(db *gorm.DB, id lib.ID, d updateUserData) (*User, error) {
	targetUser, err := ReadUserById(db, id)
	if err != nil {
		return nil, err
	}
	if updateErr := db.Model(targetUser).Clauses(clause.Returning{}).Updates(d).Error; updateErr != nil {
		return nil, updateErr
	}
	return targetUser, nil
}

// DeleteUser deletes user record from database
func DeleteUser(db *gorm.DB, id lib.ID) error {
	_, err := ReadUserById(db, id)
	if err != nil {
		return err
	}
	db.Delete(&User{}, id)
	return nil
}
