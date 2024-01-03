package models

type Brand struct {
	BrandName    string `json:"brand-name" bson:"brand-name"`
	Category string `json:"category" bson:"category"`
	
}
