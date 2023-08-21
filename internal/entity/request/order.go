package request

type OrderCreate struct {
	PersonId  int32 `json:"person_id"`
	ProductId int32 `json:"product_id"`
	Qty       int32 `json:"qty"`
}
