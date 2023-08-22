package entity

import "fmt"

type Product struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	SellerID int32  `json:"seller_id"`
	Stock    int32  `json:"stock"`
}

var ErrProductNotFound = fmt.Errorf("product %w", ErrNotFound)
