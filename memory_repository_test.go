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

func TestGetAllProductsShouldReturnAllProducts(t *testing.T) {
	repository := NewMemoryRepository()
	product1, _ := domain.NewProduct("product1", 5.5)
	product2, _ := domain.NewProduct("product2", 10.0)
	err := repository.AddProduct(context.Background(), product1)
	if err != nil {
		t.Error("Product should be added without error", err)
	}
	err = repository.AddProduct(context.Background(), product2)
	if err != nil {
		t.Error("Product should be added without error", err)
	}
	products, err := repository.GetAllProducts(context.Background())
	if err != nil {
		t.Error("Products should be retrieved without error", err)
	}
	if len(products) != 2 {
		t.Error("Should have two products")
	}
}

func TestGetAllProductsShouldReturnAnEmptyArrayIfProductDoesNotExist(t *testing.T) {
	repository := NewMemoryRepository()
	products, err := repository.GetAllProducts(context.Background())
	if err != nil {
		t.Error("Products should be retrieved without error", err)
	}
	if len(products) != 0 {
		t.Error("Should have no products")
	}
}

func TestUpdateProductShouldUpdateAProduct(t *testing.T) {
	repository := NewMemoryRepository()
	product, _ := domain.NewProduct("product", 5.5)
	err := repository.AddProduct(context.Background(), product)
	if err != nil {
		t.Error("Product should be added without error", err)
	}
	product.Name = "updated"
	err = repository.UpdateProduct(context.Background(), product)
	if err != nil {
		t.Error("Product should be updated without error", err)
	}
	productFromRepo, err := repository.GetProductById(context.Background(), product.Id)
	if err != nil {
		t.Error("Product should be retrieved without error", err)
	}
	if productFromRepo.Name != "updated" {
		t.Error("Product name should be updated")
	}
}

func TestUpdateProductShouldReturnAnErrorIfProductDoesNotExist(t *testing.T) {
	repository := NewMemoryRepository()
	product, _ := domain.NewProduct("product", 5.5)
	err := repository.UpdateProduct(context.Background(), product)
	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Error("Product should not be updated", err)
	}
}
