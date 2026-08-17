package repository

import (
	"context"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
	"github.com/RicardoEmm/rosinas-food/internal/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) *GormCustomerRepository {
	return &GormCustomerRepository{db: db}
}

func (r *GormCustomerRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Customer{}, "id = ?", id).Error
}

func (r *GormCustomerRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).Where("phone = ?", phone).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *GormCustomerRepository) FindAll(ctx context.Context) ([]*domain.Customer, error) {
	var customers []*domain.Customer

	if err := r.db.WithContext(ctx).Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *GormCustomerRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	var customer *domain.Customer

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *GormCustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

var _ service.CustomerRepo = (*GormCustomerRepository)(nil)
