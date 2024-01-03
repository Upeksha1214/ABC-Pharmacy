package services

import "abc-pharmacy.com/Upeksha1214/models"

type ItemService interface {
	CreateItem(*models.Item) error
	GetItem(*string) (*models.Item, error)
	GetAll() ([]*models.Item, error)
	UpdateItem(*models.Item) error
	DeleteItem(*string) error
}