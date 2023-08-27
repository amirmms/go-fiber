package services

import (
	"fiber_test/configs"
	"fiber_test/database"
	"gorm.io/gorm"
)

type UserService struct {
	conf configs.Config
	DB   *gorm.DB
}

func InitUserService(conf configs.Config) UserService {
	return UserService{
		conf: conf,
		DB:   database.GetDBClient(),
	}
}
