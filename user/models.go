package user

import (
	"happy-server/shared"
)

type createUserData struct {
	Email     string `json:"email" binding:"required,email,max=255"`
	Password  string `json:"password" binding:"required,max=255"`
	FirstName string `json:"firstName" binding:"required,max=255"`
	LastName  string `json:"lastName" binding:"required,max=255"`
}

type updateUserData struct {
	Email     string `json:"email" binding:"omitempty,email,max=255"`
	Password  string `json:"password" binding:"max=255"`
	FirstName string `json:"firstName" binding:"max=255"`
	LastName  string `json:"lastName" binding:"max=255"`
}

type User struct {
	shared.BaseModel
	Email     string `json:"email" gorm:"unique;size:500;not null"`
	Password  string `json:"-" gorm:"size:255;not null"`
	FirstName string `json:"firstName" gorm:"size:255;not null"`
	LastName  string `json:"lastName" gorm:"size:255;not null"`
}
