package services

import "abc-pharmacy.com/Upeksha1214/models"

type CartService interface {
	CreateCart(*models.Cart) error
	GetCart(*string) (*models.Cart, error)
	GetAll() ([]*models.Cart, error)
	UpdateCart(*models.Cart) error
	DeleteCart(*string) error
	
}