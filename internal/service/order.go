package service

import (
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo"
)

type Order struct {
	orderRepo repo.Order
}

func NewOrder(Repo repo.Order) *Order {
	return &Order{orderRepo: Repo}
}

func (o *Order) Create(r request.OrderCreate) (string, error) {
	return o.orderRepo.Create(&entity.Order{
		PersonId:  r.PersonId,
		ProductId: r.ProductId,
		Qty:       r.Qty,
	})
}

func (o *Order) GetByUser(userId int32) ([]*entity.Order, error) {
	return o.orderRepo.GetByUser(userId)
}

func (o *Order) GetByID(orderId string) (*entity.Order, error) {
	return o.orderRepo.GetByID(orderId)
}

func (o *Order) GetByProduct(productId int32) ([]*entity.Order, error) {
	return o.orderRepo.GetByProduct(productId)
}
