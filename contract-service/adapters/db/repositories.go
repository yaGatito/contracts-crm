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

func (r *ContractRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Contract{}, id).Error
}

type LegalPersonRepository struct {
	db *gorm.DB
}

var _ repo.LegalPersonRepository = (*LegalPersonRepository)(nil)

func NewLegalPersonRepository(db *gorm.DB) *LegalPersonRepository {
	return &LegalPersonRepository{db: db}
}

func (r *LegalPersonRepository) Save(ctx context.Context, legalPerson *models.LegalPerson) error {
	return r.db.WithContext(ctx).Save(legalPerson).Error
}

func (r *LegalPersonRepository) Get(ctx context.Context, id uint) (*models.LegalPerson, error) {
	var legalPerson models.LegalPerson

	if err := r.db.WithContext(ctx).First(&legalPerson, id).Error; err != nil {
		return nil, err
	}

	return &legalPerson, nil
}

func (r *LegalPersonRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.LegalPerson{}, id).Error
}

type PersonRepository struct {
	db *gorm.DB
}

var _ repo.PersonRepository = (*PersonRepository)(nil)

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Save(ctx context.Context, person *models.Person) error {
	return r.db.WithContext(ctx).Save(person).Error
}

func (r *PersonRepository) Get(ctx context.Context, id uint) (*models.Person, error) {
	var person models.Person

	if err := r.db.WithContext(ctx).First(&person, id).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Person{}, id).Error
}
