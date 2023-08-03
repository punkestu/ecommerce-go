package service

import (
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo"
)

//var ErrWrongPassword = errors.New("wrong password")

type Product struct {
	productRepo repo.Product
}

func NewProduct(Repo repo.Product) *Product {
	return &Product{productRepo: Repo}
}

func (p *Product) Create(product request.ProductCreate) (int32, error) {
	return p.productRepo.Create(entity.Product{
		Name:     product.Name,
		Price:    product.Price,
		SellerID: product.SellerID,
	})
}

func (p *Product) GetAll() ([]*entity.Product, error) {
	return p.productRepo.GetAll()
}

func (p *Product) GetByName(name string) ([]*entity.Product, error) {
	return p.productRepo.GetByName(name)
}

func (p *Product) GetByID(id int32) (*entity.Product, error) {
	return p.productRepo.GetByID(id)
}
