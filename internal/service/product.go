package service

import (
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo"
)

type Product struct {
	product repo.Product
}

func NewProduct(Repo repo.Product) *Product {
	return &Product{product: Repo}
}

func (p *Product) Create(product request.ProductCreate) (int32, error) {
	return p.product.Create(entity.Product{
		Name:     product.Name,
		Price:    product.Price,
		SellerID: product.SellerID,
		Stock:    product.InitStock,
	})
}

func (p *Product) CheckStock(ID int32, stock int32) error {
	product, err := p.product.GetByID(ID)
	if err != nil {
		return err
	}
	if product.Stock < stock {
		return entity.ErrProductOutOfStock
	}
	return nil
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
