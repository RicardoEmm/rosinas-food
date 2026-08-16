package service

import (
	"context"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
	"github.com/google/uuid"
)

type CustomerRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.Customer, error)
	FindAll(ctx context.Context) ([]*domain.Customer, error)
	Save(ctx context.Context, customer *domain.Customer) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
