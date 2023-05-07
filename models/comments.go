package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// Commment represents the model for an user
type Comment struct {
	GORMModel
	Message string `gorm:"not null" json:"message" valid:"required~Message is required"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	userData, exists := tx.Statement.Context.Value("userData").(jwt.MapClaims)
	if !exists {
		return errors.New("userData is missing in the context")
	}

	u.UserID = uint(userData["id"].(float64))

	// _, err = govalidator.ValidateStruct(u)

	// if err != nil {
	// 	return
	// }

	// photoData, exists := tx.Statement.Context.Value("photoData").(jwt.MapClaims)
	// if !exists {
	// 	return errors.New("photoData is missing in the context")
	// }
	// u.PhotoID = uint(photoData["id"].(float64))

	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}
	return
}
