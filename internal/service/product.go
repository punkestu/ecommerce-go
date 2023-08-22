package service

import (
	"github.com/punkestu/ecommerce-go/internal/domain"
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo"
)

type Product struct {
	product repo.Product
	user    domain.Person
}

func NewProduct(Repo repo.Product, userDomain domain.Person) *Product {
	return &Product{product: Repo, user: userDomain}
}

func (p *Product) Create(product request.ProductCreate) (int32, error) {
	if _, err := p.user.GetByID(product.SellerID); err != nil {
		return int32(0), err
	}
	return p.product.Create(entity.Product{
		Name:     product.Name,
		Price:    product.Price,
		SellerID: product.SellerID,
		Stock:    product.InitStock,
	})
}

func (p *Product) GetAll() ([]*entity.Product, error) {
	return p.product.GetAll()
}

func (p *Product) GetByName(name string) ([]*entity.Product, error) {
	return p.product.GetByName(name)
}

func (p *Product) GetByID(id int32) (*entity.Product, error) {
	return p.product.GetByID(id)
}
