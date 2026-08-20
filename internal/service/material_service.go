package service

import (
	"context"
	"errors"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
)

var (
	ErrMaterialNotFound          = errors.New("material not found")
	ErrMaterialNotCreated        = errors.New("material could not create")
	ErrMaterialNameAlreadyExists = errors.New("material name already exists")
	ErrInternal                  = errors.New("internal server error")
)

type MaterialService struct {
	materialRepo MaterialRepo
}

type CreateMaterialInput struct {
	Name string
}

func NewMaterialService(materialRepo MaterialRepo) *MaterialService {
	return &MaterialService{materialRepo: materialRepo}
}

func (s *MaterialService) FindById(ctx context.Context, id uint) (*domain.Material, error) {
	material, err := s.materialRepo.FindById(ctx, id)

	if err != nil {
		return nil, ErrMaterialNotFound
	}

	return material, nil
}

func (s *MaterialService) FindAll(ctx context.Context) ([]*domain.Material, error) {
	materials, err := s.materialRepo.FindAll(ctx)

	if err != nil {
		return nil, ErrInternal
	}

	return materials, nil
}

func (s *MaterialService) Create(ctx context.Context, input CreateMaterialInput) error {
	exists, err := s.materialRepo.ExistsByName(ctx, input.Name)

	if err != nil {
		return ErrInternal
	}

	if exists {
		return ErrMaterialNameAlreadyExists
	}

	if err := s.materialRepo.Save(ctx, &domain.Material{Name: input.Name}); err != nil {
		return ErrMaterialNotCreated
	}

	return nil
}

func (s *MaterialService) DeleteById(ctx context.Context, id uint) error {
	return s.materialRepo.DeleteById(ctx, id)
}
