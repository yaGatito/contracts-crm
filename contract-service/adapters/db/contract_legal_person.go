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

func (r *ContractLegalPersonRepository) Save(ctx context.Context, contract *models.ContractLegalPerson) error {
	return r.db.WithContext(ctx).Save(contract).Error
}

func (r *ContractLegalPersonRepository) GetLegalPersonsByContractID(ctx context.Context, contractId uint) ([]models.LegalPerson, error) {
	var relations []models.ContractLegalPerson

	err := r.db.WithContext(ctx).
		Preload("LegalPerson").
		Where(models.ContractLegalPerson{ContractID: contractId}).
		Find(&relations).Error

	if err != nil {
		return nil, err
	}

	legalPersons := make([]models.LegalPerson, 0, len(relations))
	for _, relation := range relations {
		legalPersons = append(legalPersons, relation.LegalPerson)
	}

	return legalPersons, nil
}

func (r *ContractLegalPersonRepository) Delete(
	ctx context.Context,
	contractID uint,
	legalPersonID uint,
) error {
	return r.db.WithContext(ctx).
		Where(models.ContractLegalPerson{
			ContractID:    contractID,
			LegalPersonID: legalPersonID,
		}).
		Delete(&models.ContractLegalPerson{}).Error
}
