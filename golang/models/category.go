package models

type Category struct {
	Category string `json:"category" bson:"category"`
	Type      string `json:"type" bson:"type"`
}
