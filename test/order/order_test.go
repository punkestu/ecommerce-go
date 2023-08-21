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
}

func TestOrder(t *testing.T) {
	r = mocks.NewOrder(t)
	s = service.NewOrder(r)
	assert.NotNil(t, r)
	assert.NotNil(t, s)
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		r.On("Create", &entity.Order{
			PersonId:  dummyOrder.PersonId,
			ProductId: dummyOrder.ProductId,
			Qty:       dummyOrder.Qty,
		}).Return(dummyOrder.ID, nil)
		id, err := s.Create(request.OrderCreate{
			PersonId:  dummyOrder.PersonId,
			ProductId: dummyOrder.ProductId,
			Qty:       dummyOrder.Qty,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, "", id)
		assert.Equal(t, dummyOrder.ID, id)
	})
	t.Run("Product not found", func(t *testing.T) {
		r.On("Create", &entity.Order{
			PersonId:  dummyOrder.PersonId,
			ProductId: dummyOrder.ProductId + 1,
			Qty:       dummyOrder.Qty,
		}).Return("", entity.ErrProductNotFound)
		id, err := s.Create(request.OrderCreate{
			PersonId:  dummyOrder.PersonId,
			ProductId: dummyOrder.ProductId + 1,
			Qty:       dummyOrder.Qty,
		})
		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrProductNotFound, err)
		assert.Equal(t, "", id)
		assert.NotEqual(t, dummyOrder.ID, id)
	})
	t.Run("User not found", func(t *testing.T) {
		r.On("Create", &entity.Order{
			PersonId:  dummyOrder.PersonId + 1,
			ProductId: dummyOrder.ProductId,
			Qty:       dummyOrder.Qty,
		}).Return("", entity.ErrPersonNotFound)
		id, err := s.Create(request.OrderCreate{
			PersonId:  dummyOrder.PersonId + 1,
			ProductId: dummyOrder.ProductId,
			Qty:       dummyOrder.Qty,
		})
		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrPersonNotFound, err)
		assert.Equal(t, "", id)
		assert.NotEqual(t, dummyOrder.ID, id)
	})
}
