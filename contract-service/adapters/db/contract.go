package sqlitedb

import (
	"context"
	"contract-service/models"
	"contract-service/repo"

	"gorm.io/gorm"
)

type ContractRepository struct {
	db *gorm.DB
}

var _ repo.ContractRepository = (*ContractRepository)(nil)

func NewContractRepository(db *gorm.DB) *ContractRepository {
	return &ContractRepository{db: db}
}

func (r *ContractRepository) Save(ctx context.Context, contract *models.Contract) error {
	return r.db.WithContext(ctx).Save(contract).Error
}

func (r *ContractRepository) Get(ctx context.Context, id uint) (*models.Contract, error) {
	var contract models.Contract

	if err := r.db.WithContext(ctx).First(&contract, id).Error; err != nil {
		return nil, err
	}

	return &contract, nil
}

func (r *ContractRepository) Update(ctx context.Context, contract *models.Contract) error {
	return r.db.WithContext(ctx).Updates(contract).Error
}

func (r *ContractRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Contract{}, id).Error
}
