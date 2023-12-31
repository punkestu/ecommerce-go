package order

import (
	"testing"

	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo/mocks"
	"github.com/punkestu/ecommerce-go/internal/service"
	"github.com/stretchr/testify/assert"
)

var r *mocks.Order
var s *service.Order

var dummyOrder = &entity.Order{
	ID:        "abcd",
	PersonId:  1,
	ProductId: 1,
	Qty:       10,
	State:     entity.StateProcessing,
}

func TestOrder(t *testing.T) {
	r = mocks.NewOrder(t)
	s = service.NewOrder(r)
	assert.NotNil(t, r)
	assert.NotNil(t, s)
}

func TestState(t *testing.T) {
	r.On("GetByID", dummyOrder.ID).Return(dummyOrder, nil)
	r.On("UpdateState", dummyOrder.ID, entity.StateDelivery).Return(nil)
	t.Run("Success", func(t *testing.T) {
		err := s.Deliver(dummyOrder.ID)
		assert.Nil(t, err)
	})
}

func TestCreate(t *testing.T) {
	r.On("Create", &entity.Order{
		PersonId:  dummyOrder.PersonId,
		ProductId: dummyOrder.ProductId,
		Qty:       dummyOrder.Qty,
		State:     entity.StateWaitForPayment,
	}).Return(dummyOrder.ID, nil)
	t.Run("Success", func(t *testing.T) {
		id, err := s.Create(request.OrderCreate{
			PersonId:  dummyOrder.PersonId,
			ProductId: dummyOrder.ProductId,
			Qty:       dummyOrder.Qty,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, "", id)
		assert.Equal(t, dummyOrder.ID, id)
	})
}
