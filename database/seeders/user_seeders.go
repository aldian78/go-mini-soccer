package seeders

import (
	"mini-soccer/constants"
	"mini-soccer/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	user := models.User{
		UUID:        uuid.New(),
		Name:        "Administator",
		Username:    "admin",
		Password:    string(password),
		PhoneNumber: "08712345678",
		Email:       "admin@gmail.com",
		RoleID:      constants.Admin,
	}

	err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("Failed to seed user: %v", err)
		panic(err)
	}
	logrus.Infof("User %s successfully seeded", user.Username)
}
