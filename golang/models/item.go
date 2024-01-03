package models

type Item struct {
	Name           string `json:"name" bson:"name"`
	Description    string `json:"description" bson:"description"`
	Category       string `json:"categorye" bson:"categorye"`
	Brand          string `json:"brand" bson:"brand"`
	Qty_on_hand    string `json:"qty_on_hand" bson:"qty_on_hand"`
	Discount       string `json:"discount" bson:"discount"`
	Volume         string `json:"volume" bson:"volume"`
	Unit_of_volume string `json:"unit_of_volume" bson:"unit_of_volume"`
	Unit_price     string `json:"unit_price" bson:"unit_price"`
}
