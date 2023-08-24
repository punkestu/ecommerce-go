package product

import (
	"testing"

	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo/mocks"
	"github.com/punkestu/ecommerce-go/internal/service"
	"github.com/stretchr/testify/assert"
)

var r *mocks.Product
var s *service.Product

var dummyProduct = &entity.Product{
	ID:       1,
	Name:     "The Product",
	Price:    1000,
	SellerID: 1,
	Stock:    100,
}

func TestPerson(t *testing.T) {
	r = mocks.NewProduct(t)
	s = service.NewProduct(r)
	assert.NotNil(t, r)
	assert.NotNil(t, s)
}

func TestCreate(t *testing.T) {
	r.On("Create", entity.Product{
		Name:     dummyProduct.Name,
		Price:    dummyProduct.Price,
		SellerID: dummyProduct.SellerID,
		Stock:    dummyProduct.Stock,
	}).Return(dummyProduct.ID, nil)
	t.Run("Success", func(t *testing.T) {
		id, err := s.Create(request.ProductCreate{
			Name:      dummyProduct.Name,
			Price:     dummyProduct.Price,
			SellerID:  dummyProduct.ID,
			InitStock: dummyProduct.Stock,
		})
		assert.Nil(t, err)
		assert.Equal(t, dummyProduct.ID, id)
	})
}
