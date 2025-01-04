package models

type User struct {
	BaseModel
	Email     string `json:"email" gorm:"unique;size:500;not null"`
	Password  string `json:"-" gorm:"size:255;not null"`
	FirstName string `json:"firstName" gorm:"size:255;not null"`
	LastName  string `json:"lastName" gorm:"size:255;not null"`
	AvatarUrl string `json:"avatarUrl" gorm:"size:500"`
}
