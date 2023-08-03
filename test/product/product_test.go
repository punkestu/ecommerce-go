package product

import (
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo/mocks"
	"github.com/punkestu/ecommerce-go/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

var r *mocks.Product
var s *service.Product

var dummyProduct = &entity.Product{
	ID:       1,
	Name:     "The Product",
	Price:    1000,
	SellerID: 1,
}

func TestPerson(t *testing.T) {
	r = mocks.NewProduct(t)
	s = service.NewProduct(r)
	assert.NotNil(t, r)
	assert.NotNil(t, s)
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		r.On("Create", entity.Product{
			ID:       int32(0),
			Name:     dummyProduct.Name,
			Price:    dummyProduct.Price,
			SellerID: dummyProduct.SellerID,
		}).Return(dummyProduct.ID, nil)
		id, err := s.Create(request.ProductCreate{
			Name:     dummyProduct.Name,
			Price:    dummyProduct.Price,
			SellerID: dummyProduct.ID,
		})
		assert.Nil(t, err)
		assert.Equal(t, dummyProduct.ID, id)
	})
}
