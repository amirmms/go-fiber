package services

import (
	"fiber_test/configs"
	"fiber_test/database"
	"fiber_test/entities/dtos"
	"fiber_test/entities/models"
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

func (u UserService) FetchUsers() (error, []models.User) {
	users := []models.User{}

	err := u.DB.Find(&users).Error

	if err != nil {
		return err, []models.User{}
	}

	return nil, users
}

func (u UserService) GetUser(req dtos.GetUserDto) (error, models.User) {
	user := models.User{ID: req.Id}

	err := u.DB.First(&user).Error

	if err != nil {
		return err, models.User{}
	}

	return nil, user
}

func (u UserService) CreateUser(req dtos.CreateUserDto) (error, models.User) {
	user := models.User{
		Name:   req.Name,
		Family: req.Family,
		Mobile: req.Mobile,
	}

	err := u.DB.Create(&user).Error

	if err != nil {
		return err, models.User{}
	}

	return nil, user
}

func (u UserService) DeleteUser(req dtos.GetUserDto) (error, models.User) {
	user := models.User{ID: req.Id}

	err := u.DB.Delete(&user).Error

	if err != nil {
		return err, models.User{}
	}

	return nil, user
}

func (u UserService) UpdateUser(req dtos.UpdateUserDto) (error, models.User) {
	user := models.User{
		ID:     req.Id,
		Name:   req.Name,
		Family: req.Family,
		Mobile: req.Mobile,
	}

	err := u.DB.Save(&user).Error

	if err != nil {
		return err, models.User{}
	}

	return nil, user
}
