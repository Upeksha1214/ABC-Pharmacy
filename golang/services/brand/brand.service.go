package services

import "abc-pharmacy.com/Upeksha1214/models"

type BrandService interface {
	CreateBrand(*models.Brand) error
	GetBrand(*string) (*models.Brand, error)
	GetAll() ([]*models.Brand, error)
	UpdateBrand(*models.Brand) error
	DeleteBrand(*string) error
}