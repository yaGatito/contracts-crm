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

func (r *ContractPersonRepository) Save(ctx context.Context, contract *models.ContractPerson) error {
	return r.db.WithContext(ctx).Save(contract).Error
}

func (r *ContractPersonRepository) GetPersonsByContractID(
	ctx context.Context,
	contractID uint,
) ([]models.Person, error) {
	var relations []models.ContractPerson

	err := r.db.WithContext(ctx).
		Preload("Person").
		Where(models.ContractPerson{ContractID: contractID}).
		Find(&relations).Error

	if err != nil {
		return nil, err
	}

	persons := make([]models.Person, 0, len(relations))
	for _, relation := range relations {
		persons = append(persons, relation.Person)
	}

	return persons, nil
}

func (r *ContractPersonRepository) Delete(
	ctx context.Context,
	contractID uint,
	personID uint,
) error {
	return r.db.WithContext(ctx).
		Where(models.ContractPerson{
			ContractID: contractID,
			PersonID:   personID,
		}).
		Delete(&models.ContractPerson{}).Error
}
