package repository

import (
	"context"

	"github.com/RicardoEmm/rosinas-food/internal/domain/incomes"
	"github.com/RicardoEmm/rosinas-food/internal/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormIncomeRepository struct {
	db *gorm.DB
}

func NewGormIncomeRepositoty(db *gorm.DB) *GormIncomeRepository {
	return &GormIncomeRepository{db: db}
}

func (r *GormIncomeRepository) ChangeToPaid(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Table("incomes").
		Where("id = ?", id).
		UpdateColumn("status", "paid").
		Error
}

func (r *GormIncomeRepository) Create(ctx context.Context, income *incomes.Income) error {
	return r.db.WithContext(ctx).Create(income).Error
}

func (r *GormIncomeRepository) DeleteAllByCustomerId(ctx context.Context, customerId uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&incomes.Income{}, "customer_id = ?", customerId).Error
}

func (r *GormIncomeRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&incomes.Income{}, "id = ?", id).Error
}

func (r *GormIncomeRepository) FindAll(ctx context.Context) ([]*incomes.Income, error) {
	var incomes []*incomes.Income

	if err := r.db.WithContext(ctx).
		Find(&incomes).
		Order("created_at DESC").
		Error; err != nil {
		return nil, err
	}

	return incomes, nil
}

func (r *GormIncomeRepository) FindAllByCustomerId(ctx context.Context, customerId uuid.UUID) ([]*incomes.Income, error) {
	var incomes []*incomes.Income

	if err := r.db.WithContext(ctx).
		Where("customer_id = ?", customerId).
		Find(&incomes).
		Order("created_at DESC").
		Error; err != nil {
		return nil, err
	}

	return incomes, nil
}

func (r *GormIncomeRepository) FindById(ctx context.Context, id uuid.UUID) (*incomes.Income, error) {
	var income *incomes.Income

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&income).Error; err != nil {
		return nil, err
	}

	return income, nil
}

var _ service.IncomeRepo = (*GormIncomeRepository)(nil)
