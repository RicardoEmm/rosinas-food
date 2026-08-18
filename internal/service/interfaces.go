package service

import (
	"context"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
	"github.com/google/uuid"
)

type CustomerRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.Customer, error)
	FindAll(ctx context.Context) ([]*domain.Customer, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	Save(ctx context.Context, customer *domain.Customer) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type MaterialRepo interface {
	FindById(ctx context.Context, id uint) (*domain.Material, error)
	FindAll(ctx context.Context) ([]*domain.Material, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	Save(ctx context.Context, material *domain.Material) error
	DeleteById(ctx context.Context, id uint) error
}
