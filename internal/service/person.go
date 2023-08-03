package service

import (
	"github.com/punkestu/ecommerce-go/internal/entity"
	"github.com/punkestu/ecommerce-go/internal/entity/request"
	"github.com/punkestu/ecommerce-go/internal/repo"
)

type Person struct {
	userRepo repo.Person
}

func NewPerson(Repo repo.Person) *Person {
	return &Person{userRepo: Repo}
}

func (p *Person) Register(person request.PersonRegister) (int32, error) {
	return p.userRepo.Create(entity.Person{
		Name:     person.Name,
		Email:    person.Email,
		Password: person.Password,
	})
}

func (p *Person) Login(credential request.PersonLogin) (int32, error) {
	person, err := p.userRepo.GetByEmail(credential.Email)
	if err != nil {
		return 0, err
	}
	if person.Password != credential.Password {
		return 0, entity.ErrWrongPassword
	}
	return person.ID, nil
}

func (p *Person) GetProfile(id int32) (*entity.Person, error) {
	return p.userRepo.GetByID(id)
}
