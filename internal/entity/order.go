package entity

type Order struct {
	ID        string `json:"id"`
	PersonId  int32  `json:"person_id"`
	ProductId int32  `json:"product_id"`
	Qty       int32  `json:"qty"`
}
