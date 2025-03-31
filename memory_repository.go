package ProductManager

import (
	"context"
	"github.com/google/uuid"
	"sync"

	domain "ProductManager/domain"
)

type MemoryRepository struct {
	products map[uuid.UUID]*domain.Product
	lock     sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		products: make(map[uuid.UUID]*domain.Product),
	}
}

func (r *MemoryRepository) AddProduct(ctx context.Context, product *domain.Product) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, exists := r.products[product.Id]; exists {
		return domain.ErrProductAlreadyExists
	}

	r.products[product.Id] = product
	return nil
}

func (r *MemoryRepository) GetProductById(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if _, exists := r.products[id]; !exists {
		return nil, domain.ErrProductNotFound
	}
	return r.products[id], nil
}
