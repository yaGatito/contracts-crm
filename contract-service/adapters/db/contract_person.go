package sqlitedb

import (
	"context"
	"contract-service/models"

	"gorm.io/gorm"
)

type ContractPersonRepository struct {
	db *gorm.DB
}

// var _ repo.ContractPersonRepository = (*ContractPersonRepository)(nil)

func NewContractPersonRepository(db *gorm.DB) *ContractPersonRepository {
	return &ContractPersonRepository{db: db}
}

func (r *ContractPersonRepository) Save(ctx context.Context, contract *models.Contract) error {
	return r.db.WithContext(ctx).Save(contract).Error
}

func (r *ContractPersonRepository) Get(ctx context.Context, id uint) (*models.Contract, error) {
	var contract models.Contract

	if err := r.db.WithContext(ctx).First(&contract, id).Error; err != nil {
		return nil, err
	}

	return &contract, nil
}

func (r *ContractPersonRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Contract{}, id).Error
}
