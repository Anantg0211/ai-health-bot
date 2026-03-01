package services

import (
	"ai-powered-health-bot/db"
	"ai-powered-health-bot/models"
	"errors"

	"gorm.io/gorm"
)

func GetOrCreateUser(mobile string) (*models.User, error) {

	var user models.User

	err := db.DB.Where("mobile = ?", mobile).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		user = models.User{
			Mobile: mobile,
		}

		err = db.DB.Create(&user).Error
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}