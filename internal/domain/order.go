package domain

import "github.com/punkestu/ecommerce-go/internal/entity"

type Order interface {
	GetByID(string) (*entity.Order, error)
	GetByUser(int32) ([]*entity.Order, error)
	GetByProduct(int32) ([]*entity.Order, error)
	Create(*entity.Order) (string, error)
}
