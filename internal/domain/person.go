package domain

import "github.com/punkestu/ecommerce-go/internal/entity"

type Person interface {
	GetByID(int32) (*entity.Person, error)
	GetByEmail(string) (*entity.Person, error)
	Create(entity.Person) (int32, error)
}
