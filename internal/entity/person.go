package entity

import "fmt"

type Person struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ErrPersonNotFound = fmt.Errorf("person %w", ErrNotFound)
