package user

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"happy-server/shared"
)

// GetUsers gets all users from database
func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser gets one user from database by primary key (id)
func GetUser(db *gorm.DB, id shared.ID) (*User, error) {
	var foundUser User
	if err := db.First(&foundUser, id).Error; err != nil {
		return nil, err
	}
	return &foundUser, nil
}

// CreateUser creates new user record in database
func CreateUser(db *gorm.DB, d createUserData) (*User, error) {
	if err := db.First(&User{}, "email = ?", d.Email); err == nil {
		return nil, errors.New("user already exists")
	}
	newUser := User{
		Email:     d.Email,
		Password:  d.Password,
		FirstName: d.FirstName,
		LastName:  d.LastName,
	}
	db.Create(&newUser)
	return &newUser, nil
}

// UpdateUser updates user record in database
func UpdateUser(db *gorm.DB, id shared.ID, d updateUserData) (*User, error) {
	targetUser, err := GetUser(db, id)
	if err != nil {
		return nil, err
	}
	if updateErr := db.Model(targetUser).Clauses(clause.Returning{}).Updates(d).Error; updateErr != nil {
		return nil, updateErr
	}
	return targetUser, nil
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
