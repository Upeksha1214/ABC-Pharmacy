package services

import "abc-pharmacy.com/Upeksha1214/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
	Login(username, password string) (*models.User, error)
}
