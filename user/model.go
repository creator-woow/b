package user

import "happy-server/lib"

type (
	CreateUserData struct {
		Email     string `json:"email" binding:"required,email,max=255"`
		Password  string `json:"password" binding:"required,max=255"`
		FirstName string `json:"firstName" binding:"required,max=255"`
		LastName  string `json:"lastName" binding:"required,max=255"`
	}
	updateUserData struct {
		Email     string `json:"email" binding:"omitempty,email,max=255"`
		Password  string `json:"password" binding:"max=255"`
		FirstName string `json:"firstName" binding:"max=255"`
		LastName  string `json:"lastName" binding:"max=255"`
	}
	User struct {
		lib.BaseModel
		Email     string `json:"email" gorm:"unique;size:500;not null"`
		Password  string `json:"-" gorm:"size:255;not null"`
		FirstName string `json:"firstName" gorm:"size:255;not null"`
		LastName  string `json:"lastName" gorm:"size:255;not null"`
		AvatarUrl string `json:"avatarUrl" gorm:"size:500"`
	}
)
