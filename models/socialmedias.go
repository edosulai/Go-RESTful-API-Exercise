package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// Social Media represents the model for an user
type SocialMedia struct {
	GORMModel
	Name             string `gorm:"not null" json:"name" valid:"required~Name is required and cannot be empty"`
	Social_Media_Url string `gorm:"not null" json:"social_media_url" valid:"required~Social Media Url is required and cannot be empty"`
	UserID           uint   `gorm:"not null" json:"user_id"`
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	userData, exists := tx.Statement.Context.Value("userData").(jwt.MapClaims)
	if !exists {
		return errors.New("userData is missing in the context")
	}

	u.UserID = uint(userData["id"].(float64))

	_, err = govalidator.ValidateStruct(u)

	if err != nil {
		return
	}

	return
}
