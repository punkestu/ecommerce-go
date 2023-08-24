package service

import (
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo"
)

type Order struct {
	order repo.Order
}

func NewOrder(Repo repo.Order) *Order {
	return &Order{order: Repo}
}

func (o *Order) Create(r request.OrderCreate) (string, error) {
	return o.order.Create(&entity.Order{
		PersonId:  r.PersonId,
		ProductId: r.ProductId,
		Qty:       r.Qty,
		State:     entity.StateWaitForPayment,
	})
}

func (o *Order) Process(ID string) error {
	order, err := o.order.GetByID(ID)
	if err != nil {
		return err
	}
	if order.State >= entity.StateProcessing {
		return entity.ErrPaid
	}
	return o.order.UpdateState(ID, entity.StateProcessing)
}

func (o *Order) Deliver(ID string) error {
	order, err := o.order.GetByID(ID)
	if err != nil {
		return err
	}
	if order.State < entity.StateProcessing {
		return entity.ErrNotPaid
	}
	if order.State >= entity.StateDelivery {
		return entity.ErrDelivered
	}
	return o.order.UpdateState(ID, entity.StateDelivery)
}

func (o *Order) Finish(ID string) error {
	order, err := o.order.GetByID(ID)
	if err != nil {
		return err
	}
	if order.State < entity.StateDelivery {
		return entity.ErrNotDelivered
	}
	if order.State >= entity.StateFinish {
		return entity.ErrFinished
	}
	return o.order.UpdateState(ID, entity.StateFinish)
}

func (o *Order) Cancel(ID string) error {
	order, err := o.order.GetByID(ID)
	if err != nil {
		return err
	}
	if order.State == entity.StateCanceled {
		return entity.ErrCanceled
	}
	return o.order.UpdateState(ID, entity.StateCanceled)
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
