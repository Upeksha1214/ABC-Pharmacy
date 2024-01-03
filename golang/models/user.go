package models

type User struct {
	Name      string `json:"name" bson:"name"`
	MobileNo string `json:"mobile-no" bson:"mobile-no"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Address string `json:"address" bson:"address"`
	
}
