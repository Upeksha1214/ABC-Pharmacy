package services

import "abc-pharmacy.com/Upeksha1214/models"

type CategoryService interface {
	CreateCategory(*models.Category) error
	GetCategory(*string) (*models.Category, error)
	GetAll() ([]*models.Category, error)
	UpdateCategory(*models.Category) error
	DeleteCategory(*string) error

}
