package config

import (
	"fmt"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup Database
func InitDatabase() (*gorm.DB, error) {
	config := Config
	// uri := fmt.Sprintf("postgreqsl://%s:%s@%s:%d/%s?sslmode=disable",
	// 	config.Database.Username,
	// 	encodedPassword,
	// 	config.Database.Host,
	// 	config.Database.Port,
	// 	config.Database.Name,
	// )

	encodedPassword := url.QueryEscape(config.Database.Password)
	uri := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%v sslmode=disable",
		config.Database.Username,
		encodedPassword,
		config.Database.Name,
		config.Database.Port,
	)

	logrus.Infof("check uri db : %s", uri)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.MaxLifetimeConnection) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Database.MaxIdleTime) * time.Second)

	return db, nil
}
