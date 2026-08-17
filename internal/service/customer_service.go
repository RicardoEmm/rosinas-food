package service

import (
	"context"
	"errors"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
	"github.com/RicardoEmm/rosinas-food/internal/dto"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound   = errors.New("customer not found")
	ErrInternalServer     = errors.New("internal error server")
	ErrPhoneAlreadyExists = errors.New("phone already exists")
	ErrCustomerNotCreated = errors.New("customer could not create")
)

type CustomerService struct {
	customerRepo CustomerRepo
}

func NewCustomerService(customerRepo CustomerRepo) *CustomerService {
	return &CustomerService{customerRepo: customerRepo}
}

func (s *CustomerService) FindById(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	customer, err := s.customerRepo.FindById(ctx, id)

	if err != nil {
		return nil, ErrCustomerNotFound
	}

	return customer, nil
}

func (s *CustomerService) FindAll(ctx context.Context) ([]*domain.Customer, error) {
	customers, err := s.customerRepo.FindAll(ctx)

	if err != nil {
		return nil, ErrInternalServer
	}

	return customers, nil
}

func (s *CustomerService) Create(ctx context.Context, input dto.CreateCustomerInput) error {
	ok, err := s.customerRepo.ExistsByPhone(ctx, input.Phone)

	if err != nil {
		return ErrInternalServer
	}

	if !ok {
		return ErrPhoneAlreadyExists
	}

	if err := s.customerRepo.Save(ctx, &domain.Customer{
		FullName: input.FullName,
		Phone:    input.Phone,
		Price:    input.Price,
	}); err != nil {
		return ErrCustomerNotCreated
	}
	return nil
}

func (s *CustomerService) DeleteById(ctx context.Context, id uuid.UUID) error {
	return s.DeleteById(ctx, id)
}
