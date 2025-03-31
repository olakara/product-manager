package ProductManager

import (
	domain "ProductManager/domain"
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestNewMemoryRepositoryShouldReturnARepositoryImplementation(t *testing.T) {
	repository := NewMemoryRepository()
	if repository == nil {
		t.Error("Repository should not be nil")
	}
}

func TestAddProductShouldAddProduct(t *testing.T) {
	repository := NewMemoryRepository()
	product, _ := domain.NewProduct("product", 5.5)
	err := repository.AddProduct(context.Background(), product)
	if err != nil {
		t.Error("Product should be added without error", err)
	}
}

func TestAddProductShouldReturnAnErrorIfProductAlreadyExists(t *testing.T) {
	repository := NewMemoryRepository()
	product, _ := domain.NewProduct("product", 5.5)
	err := repository.AddProduct(context.Background(), product)
	if err != nil {
		t.Error("Product should be added without error", err)
	}
	err = repository.AddProduct(context.Background(), product)
	if !errors.Is(err, domain.ErrProductAlreadyExists) {
		t.Error("Product should not be added again", err)
	}
}

func TestGetProductShouldReturnAProduct(t *testing.T) {
	repository := NewMemoryRepository()
	product, _ := domain.NewProduct("product", 5.5)
	err := repository.AddProduct(context.Background(), product)
	if err != nil {
		t.Error("Product should be added without error", err)
	}
	productFromRepo, err := repository.GetProductById(context.Background(), product.Id)
	if err != nil {
		t.Error("Product should be retrieved without error", err)
	}
	if productFromRepo == nil {
		t.Error("Product should not be nil")
	}
	if productFromRepo.Id != product.Id {
		t.Error("Product id should be the same")
	}
}

func TestGetProductShouldReturnAnErrorIfProductDoesNotExist(t *testing.T) {
	repository := NewMemoryRepository()

	productFromRepo, err := repository.GetProductById(context.Background(), uuid.New())
	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Error("Product should not be retrieved", err)
	}
	if productFromRepo != nil {
		t.Error("Product should be nil")
	}
}
