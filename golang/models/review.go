package models

type Review struct {
	Item_Id          string `json:"item_Id" bson:"item_Id"`
	Description      string `json:"description" bson:"description"`
	Ratings          string `json:"ratings" bson:"ratings"`
	Review_user_name string `json:"review_user_name" bson:"review_user_name"`
	Review_date      string `json:"review_date" bson:"review_date"`
}
