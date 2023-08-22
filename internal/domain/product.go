package domain

import "github.com/punkestu/ecommerce-go/internal/entity"

type Product interface {
	GetByID(int32) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	GetByName(string) ([]*entity.Product, error)
	Create(entity.Product) (int32, error)
}
