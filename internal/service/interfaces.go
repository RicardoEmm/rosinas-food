package service

import (
	"context"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
	"github.com/RicardoEmm/rosinas-food/internal/domain/incomes"
	"github.com/google/uuid"
)

type CustomerRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.Customer, error)
	FindAll(ctx context.Context) ([]*domain.Customer, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	Save(ctx context.Context, customer *domain.Customer) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type IncomeRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*incomes.Income, error)
	FindAll(ctx context.Context) ([]*incomes.Income, error)
	FindAllByCustomerId(ctx context.Context, customerId uuid.UUID) ([]*incomes.Income, error)
	Create(ctx context.Context, income *incomes.Income) error
	ChangeToPaid(ctx context.Context, id uuid.UUID) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	DeleteAllByCustomerId(ctx context.Context, customerId uuid.UUID) error
}

type MaterialRepo interface {
	FindById(ctx context.Context, id uint) (*domain.Material, error)
	FindAll(ctx context.Context) ([]*domain.Material, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	Save(ctx context.Context, material *domain.Material) error
	DeleteById(ctx context.Context, id uint) error
}
