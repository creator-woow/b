package user

import (
	"errors"
	"gorm.io/gorm"
	"happy-server/shared"
)

// GetUsers gets all users from database
func GetUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

// GetUser gets one user from database by primary key (id)
func GetUser(db *gorm.DB, id shared.ID) (*User, error) {
	var foundUser User
	db.First(&foundUser, id)
	if foundUser.ID == 0 {
		return nil, errors.New("user not found")
	}
	return &foundUser, nil
}

// CreateUser creates new user record in database
func CreateUser(db *gorm.DB, d createUserData) (*User, error) {
	var foundUser User
	db.First(&foundUser, "email = ?", d.Email)
	if foundUser.ID != 0 {
		return nil, errors.New("user already exists")
	}
	newUser := &User{
		Email:     d.Email,
		Password:  d.Password,
		FirstName: d.FirstName,
		LastName:  d.LastName,
	}
	db.Create(newUser)
	return newUser, nil
}

// DeleteUser deletes user record from database
func DeleteUser(db *gorm.DB, id shared.ID) error {
	_, err := GetUser(db, id)
	if err != nil {
		return err
	}
	db.Delete(&User{}, id)
	return nil
}

// UpdateUser updates user record in database
func UpdateUser(db *gorm.DB, id shared.ID, d updateUserData) (*User, error) {
	foundUser, err := GetUser(db, id)
	if err != nil {
		return nil, err
	}
	db.Model(foundUser).Where("id = ?", id).Updates(d)
	return foundUser, nil
}
