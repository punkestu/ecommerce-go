package request

type ProductCreate struct {
	Name      string `json:"name"`
	Price     int32  `json:"price"`
	SellerID  int32  `json:"seller_id"`
	InitStock int32  `json:"init_stock"`
}
