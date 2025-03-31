package domain

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

type Repository interface {
	AddProduct(context.Context, *Product) error
	GetProductById(context.Context, uuid.UUID) (*Product, error)
	GetAllProducts(context.Context) ([]*Product, error)
	UpdateProduct(context.Context, *Product) error
	RemoveProduct(context.Context, uuid.UUID) error
}

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)
