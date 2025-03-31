package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Product struct {
	Id    uuid.UUID
	Name  string
	Price float64
}

var (
	ErrPriceNegative = errors.New("price cannot be negative")
)

func NewProduct(name string, price float64) (*Product, error) {
	id := uuid.New()

	if price < 0 {
		return nil, ErrPriceNegative
	}

	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
	}, nil
}
