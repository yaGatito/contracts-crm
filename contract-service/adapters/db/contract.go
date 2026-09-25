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

func (r *ContractRepository) GetContractDetails(ctx context.Context, id uint) (models.ContractDetails, error) {
	var contract models.Contract

	if err := r.db.WithContext(ctx).First(&contract, id).Error; err != nil {
		return models.ContractDetails{}, err
	}

	var personRelations []models.ContractPerson

	err := r.db.WithContext(ctx).
		Preload("Person").
		Where(models.ContractPerson{ContractID: id}).
		Find(&personRelations).Error

	if err != nil {
		return models.ContractDetails{}, err
	}

	persons := make([]models.Person, 0, len(personRelations))
	for _, relation := range personRelations {
		persons = append(persons, relation.Person)
	}

	var legalRelations []models.ContractLegalPerson

	err = r.db.WithContext(ctx).
		Preload("LegalPerson").
		Where(models.ContractLegalPerson{ContractID: id}).
		Find(&legalRelations).Error

	if err != nil {
		return models.ContractDetails{}, err
	}

	legalPersons := make([]models.LegalPerson, 0, len(legalRelations))
	for _, relation := range legalRelations {
		legalPersons = append(legalPersons, relation.LegalPerson)
	}

	return models.ContractDetails{
		Contract:     contract,
		Persons:      persons,
		LegalPersons: legalPersons,
	}, nil
}

func (r *ContractRepository) Update(ctx context.Context, contract *models.Contract) error {
	return r.db.WithContext(ctx).Updates(contract).Error
}

func (r *ContractRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Contract{}, id).Error
}

func (r *ContractRepository) SearchLegalPersonContracts(
	ctx context.Context,
	filter models.SearchLegalPersonContractsFilter,
) ([]models.Contract, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Contract{}).
		Joins(`
			JOIN contract_legal_people clp
				ON clp.contract_id = contracts.id
		`).
		Joins(`
			JOIN legal_people lp
				ON lp.id = clp.legal_person_id
		`)

	if filter.ContractType != "" {
		query = query.Where(
			"contracts.type = ?",
			filter.ContractType,
		)
	}

	if !filter.StartDateFrom.IsZero() {
		query = query.Where(
			"contracts.start_date >= ?",
			filter.StartDateFrom,
		)
	}

	if !filter.StartDateTo.IsZero() {
		query = query.Where(
			"contracts.start_date <= ?",
			filter.StartDateTo,
		)
	}

	if filter.LegalPersonShortName != "" {
		query = query.Where(
			"lp.short_name = ?",
			filter.LegalPersonShortName,
		)
	}

	if filter.LegalPersonName != "" {
		query = query.Where(
			"lp.name = ?",
			filter.LegalPersonName,
		)
	}

	if filter.LocalEDRPOU != 0 {
		query = query.Where(
			"lp.local_edrpou = ?",
			filter.LocalEDRPOU,
		)
	}

	var contracts []models.Contract

	if err := query.Find(&contracts).Error; err != nil {
		return nil, err
	}

	return contracts, nil
}

func (r *ContractRepository) SearchPersonContracts(
	ctx context.Context,
	filter models.SearchPersonContractsFilter,
) ([]models.Contract, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Contract{}).
		Joins(`
			JOIN contract_persons cp
				ON cp.contract_id = contracts.id
		`).
		Joins(`
			JOIN people p
				ON p.id = cp.person_id
		`)

	if filter.ContractType != "" {
		query = query.Where(
			"contracts.type = ?",
			filter.ContractType,
		)
	}

	if !filter.StartDateFrom.IsZero() {
		query = query.Where(
			"contracts.start_date >= ?",
			filter.StartDateFrom,
		)
	}

	if !filter.StartDateTo.IsZero() {
		query = query.Where(
			"contracts.start_date <= ?",
			filter.StartDateTo,
		)
	}

	if filter.PersonName != "" {
		query = query.Where(
			"p.name = ?",
			filter.PersonName,
		)
	}

	if filter.PersonLastname != "" {
		query = query.Where(
			"p.lastname = ?",
			filter.PersonLastname,
		)
	}

	if filter.PersonPatronym != "" {
		query = query.Where(
			"p.patronym = ?",
			filter.PersonPatronym,
		)
	}

	if !filter.PersonDateOfBirth.IsZero() {
		query = query.Where(
			"p.date_of_birth = ?",
			filter.PersonDateOfBirth,
		)
	}

	var contracts []models.Contract

	if err := query.Find(&contracts).Error; err != nil {
		return nil, err
	}

	return contracts, nil
}
