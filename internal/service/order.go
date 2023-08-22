package service

import (
	"github.com/punkestu/ecommerce-go/internal/domain"
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo"
)

type Order struct {
	order   repo.Order
	user    domain.Person
	product domain.Product
}

func NewOrder(Repo repo.Order, userDomain domain.Person, productDomain domain.Product) *Order {
	return &Order{order: Repo, user: userDomain, product: productDomain}
}

func (o *Order) Create(r request.OrderCreate) (string, error) {
	if _, err := o.user.GetByID(r.PersonId); err != nil {
		return "", err
	}
	if p, err := o.product.GetByID(r.ProductId); err != nil {
		return "", err
	} else if p.Stock < r.Qty {
		return "", entity.ErrProductOutOfStock
	}
	return o.order.Create(&entity.Order{
		PersonId:  r.PersonId,
		ProductId: r.ProductId,
		Qty:       r.Qty,
	})
}

func (o *Order) GetByUser(userId int32) ([]*entity.Order, error) {
	return o.order.GetByUser(userId)
}

func (o *Order) GetByID(orderId string) (*entity.Order, error) {
	return o.order.GetByID(orderId)
}

func (o *Order) GetByProduct(productId int32) ([]*entity.Order, error) {
	return o.order.GetByProduct(productId)
}
