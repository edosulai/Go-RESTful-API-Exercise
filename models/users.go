package models

import (
	"gorest/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GORMModel
	UserName string `gorm:"not null" json:"username" valid:"required~User Name is required and cannot be empty"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      uint8  `gorm:"not null" json:"age" valid:"required,range(8|2147483647)~Age must be greater than or equal to 8 and cannot be empty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass

	return
}
