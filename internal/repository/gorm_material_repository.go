package repository

import (
	"context"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
	"github.com/RicardoEmm/rosinas-food/internal/service"
	"gorm.io/gorm"
)

type GormMaterialRepository struct {
	db *gorm.DB
}

func NewGormMaterialRepository(db *gorm.DB) *GormMaterialRepository {
	return &GormMaterialRepository{db: db}
}

func (r *GormMaterialRepository) FindById(ctx context.Context, id uint) (*domain.Material, error) {
	var material *domain.Material

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&material).Error; err != nil {
		return nil, err
	}

	return material, nil
}

func (r *GormMaterialRepository) FindAll(ctx context.Context) ([]*domain.Material, error) {
	var materials []*domain.Material

	if err := r.db.WithContext(ctx).
		Find(&materials).
		Order("name ASC").
		Error; err != nil {
		return nil, err
	}

	return materials, nil
}

func (r *GormMaterialRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).
		Table("materials").
		Where("name = ?", name).
		Count(&count).
		Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *GormMaterialRepository) Save(ctx context.Context, material *domain.Material) error {
	return r.db.WithContext(ctx).Create(material).Error
}

func (r *GormMaterialRepository) DeleteById(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Material{}, "id = ?", id).Error
}

var _ service.MaterialRepo = (*GormMaterialRepository)(nil)
