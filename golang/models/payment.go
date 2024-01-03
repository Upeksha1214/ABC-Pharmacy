package models

type Payment struct {
	ID             string           `json:"id" bson:"id"`
	Name           string           `json:"name" bson:"name"`
	MobileNo       string           `json:"mobileNo" bson:"mobileNo"`
	Email          string           `json:"email" bson:"email"`
	Address        string           `json:"address" bson:"address"`
	BillingType    string           `json:"billingType" bson:"billingType"`
	PaymentDetails PaymentDetails   `json:"paymentDetails" bson:"paymentDetails"`
	PaymentDate    string           `json:"paymentDate" bson:"paymentDate"`
	Cart           []CartItemDetail `json:"cart" bson:"cart"`
}

type PaymentDetails struct {
	PaymentMethod string `json:"paymentMethod" bson:"paymentMethod"`
	TransactionID string `json:"transactionID" bson:"transactionID"`
	
}

type CartItemDetail struct {
	ItemID           string  `json:"itemID" bson:"itemID"`
	Name             string  `json:"name" bson:"name"`
	Brand            string  `json:"brand" bson:"brand"`
	Qty              int     `json:"qty" bson:"qty"`
	UnitPrice        float64 `json:"unitPrice" bson:"unitPrice"`
	TotalUnitsPrice  float64 `json:"totalUnitsPrice" bson:"totalUnitsPrice"`
	
}


