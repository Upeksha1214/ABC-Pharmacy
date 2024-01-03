package models

type Cart struct {
	UserId     string `json:"user_Id" bson:"user_Id"`
	Item_Id    string `json:"item_Id" bson:"item_Id"`
	Name       string `json:"name" bson:"name"`
	Brand      string `json:"brand" bson:"brand"`
	Qty        string `json:"qty" bson:"qty"`
	Unit_Price string `json:"unit_price" bson:"unit_price"`
	Total      string `json:"total_units_price" bson:"total_units_price"`
}
