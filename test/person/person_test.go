package person

import (
	"errors"
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo/mocks"
	"github.com/punkestu/ecommerce-go/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

var r *mocks.Person
var s *service.Person

var dummyPerson = &entity.Person{
	ID:       1,
	Name:     "minerva",
	Email:    "test@mail.com",
	Password: "test1234",
}

func TestPerson(t *testing.T) {
	r = mocks.NewPerson(t)
	s = service.NewPerson(r)
	assert.NotNil(t, r)
	assert.NotNil(t, s)
}

func TestLogin(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		r.On("GetByEmail", "test@mail.com").Return(dummyPerson, nil)
		id, err := s.Login(request.PersonLogin{
			Email:    dummyPerson.Email,
			Password: dummyPerson.Password,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, 0, id)
		assert.Equal(t, dummyPerson.ID, id)
	})
	t.Run("Wrong Password", func(t *testing.T) {
		r.On("GetByEmail", "test@mail.com").Return(dummyPerson, nil)
		id, err := s.Login(request.PersonLogin{
			Email:    dummyPerson.Email,
			Password: "test123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrWrongPassword, err)
		assert.Equal(t, int32(0), id)
		assert.NotEqual(t, dummyPerson.ID, id)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		r.On("Create", entity.Person{
			ID:       0,
			Name:     dummyPerson.Name,
			Email:    dummyPerson.Email,
			Password: dummyPerson.Password,
		}).Return(dummyPerson.ID, nil)
		id, err := s.Register(request.PersonRegister{
			Name:     dummyPerson.Name,
			Email:    dummyPerson.Email,
			Password: dummyPerson.Password,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, int32(0), id)
		assert.Equal(t, dummyPerson.ID, id)
	})
	t.Run("Failed", func(t *testing.T) {
		r.On("Create", entity.Person{
			ID:       0,
			Name:     "testFailed",
			Email:    dummyPerson.Email,
			Password: dummyPerson.Password,
		}).Return(int32(0), errors.New("failed"))
		id, err := s.Register(request.PersonRegister{
			Name:     "testFailed",
			Email:    dummyPerson.Email,
			Password: dummyPerson.Password,
		})
		assert.NotNil(t, err)
		assert.Equal(t, int32(0), id)
		assert.NotNil(t, dummyPerson.ID, id)
	})
}

func TestGetProfile(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		r.On("GetByID", dummyPerson.ID).Return(dummyPerson, nil)
		person, err := s.GetProfile(dummyPerson.ID)
		assert.Nil(t, err)
		assert.Equal(t, *dummyPerson, *person)
	})
	t.Run("Not Found", func(t *testing.T) {
		r.On("GetByID", int32(0)).Return(nil, entity.ErrNotFound)
		person, err := s.GetProfile(int32(0))
		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, person)
	})
}
