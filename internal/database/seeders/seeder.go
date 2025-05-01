package seeders

import (
	"core.yuyuid.id/internal/app/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {

	// hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("@Admin123"), 4)
	if err != nil {
		return err
	}
	users := []model.User{
		{
			FirstName: "Mochamad Yudi",
			LastName:  "Sobari",
			Email:     "admin@gmail.com",
			Status:    "ACTIVE",
			Password:  string(hashedPassword),
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.doe@example.com",
			Status:    "ACTIVE",
			Password:  string(hashedPassword),
		},
		{
			FirstName: "John",
			LastName:  "Smith",
			Email:     "john.smith@example.com",
			Status:    "ACTIVE",
			Password:  string(hashedPassword),
		},
	}

	// for
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
