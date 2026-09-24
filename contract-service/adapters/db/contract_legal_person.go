package sqlitedb

import (
	"context"
	"contract-service/models"

	"gorm.io/gorm"
)

type ContractLegalPersonRepository struct {
	db *gorm.DB
}

// var _ repo.ContractLegalPersonRepository = (*ContractLegalPersonRepository)(nil)

func NewContractLegalPersonRepository(db *gorm.DB) *ContractLegalPersonRepository {
	return &ContractLegalPersonRepository{db: db}
}

func (r *ContractLegalPersonRepository) Save(ctx context.Context, contract *models.Contract) error {
	return r.db.WithContext(ctx).Save(contract).Error
}

func (r *ContractLegalPersonRepository) Get(ctx context.Context, id uint) (*models.Contract, error) {
	var contract models.Contract

	if err := r.db.WithContext(ctx).First(&contract, id).Error; err != nil {
		return nil, err
	}

	return &contract, nil
}

func (r *ContractLegalPersonRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Contract{}, id).Error
}
