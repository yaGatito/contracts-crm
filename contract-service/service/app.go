package service

import (
	"context"
	sqlitedb "contract-service/adapters/db"
	"contract-service/models"
	"contract-service/repo"
)

type ContractService struct {
	personRepo              repo.PersonRepository
	legalPersonRepo         repo.LegalPersonRepository
	contractRepo            repo.ContractRepository
	contractLegalPersonRepo *sqlitedb.ContractLegalPersonRepository
	contractPersonRepo      *sqlitedb.ContractPersonRepository
}

func NewContractService(pr repo.PersonRepository, lpr repo.LegalPersonRepository, cr repo.ContractRepository, clpr *sqlitedb.ContractLegalPersonRepository, cpr *sqlitedb.ContractPersonRepository) *ContractService {
	return &ContractService{
		personRepo:              pr,
		legalPersonRepo:         lpr,
		contractRepo:            cr,
		contractLegalPersonRepo: clpr,
		contractPersonRepo:      cpr,
	}
}

func (cs *ContractService) CreatePerson(ctx context.Context, person *models.Person) error {
	return cs.personRepo.Save(ctx, person)
}

func (cs *ContractService) GetPerson(ctx context.Context, id uint) (*models.Person, error) {
	return cs.personRepo.Get(ctx, id)
}

func (cs *ContractService) DeletePerson(ctx context.Context, id uint) error {
	return cs.personRepo.Delete(ctx, id)
}

func (cs *ContractService) CreateLegalPerson(ctx context.Context, legalPerson *models.LegalPerson) error {
	return cs.legalPersonRepo.Save(ctx, legalPerson)
}

func (cs *ContractService) GetLegalPerson(ctx context.Context, id uint) (*models.LegalPerson, error) {
	return cs.legalPersonRepo.Get(ctx, id)
}

func (cs *ContractService) DeleteLegalPerson(ctx context.Context, id uint) error {
	return cs.legalPersonRepo.Delete(ctx, id)
}

func (cs *ContractService) CreateContract(ctx context.Context, contract *models.Contract) error {
	return cs.contractRepo.Save(ctx, contract)
}

func (cs *ContractService) GetContractDetails(ctx context.Context, id uint) (models.ContractDetails, error) {
	return cs.contractRepo.GetContractDetails(ctx, id)
}

func (cs *ContractService) SearchLegalPersonContracts(
	ctx context.Context,
	filter models.SearchLegalPersonContractsFilter,
) ([]models.Contract, error) {
	return cs.contractRepo.SearchLegalPersonContracts(ctx, filter)
}

func (cs *ContractService) SearchPersonContracts(
	ctx context.Context,
	filter models.SearchPersonContractsFilter,
) ([]models.Contract, error) {
	return cs.contractRepo.SearchPersonContracts(ctx, filter)
}

func (cs *ContractService) UpdateContract(ctx context.Context, contract *models.Contract) error {
	return cs.contractRepo.Update(ctx, contract)
}

func (cs *ContractService) DeleteContract(ctx context.Context, id uint) error {
	return cs.contractRepo.Delete(ctx, id)
}

func (cs *ContractService) AddPersonToContract(ctx context.Context, contractId uint, personId uint) error {
	return cs.contractPersonRepo.Save(ctx, &models.ContractPerson{
		ContractID: contractId,
		PersonID:   personId,
	})
}

func (cs *ContractService) RemovePersonFromContract(ctx context.Context, contractId uint, personId uint) error {
	return cs.contractPersonRepo.Delete(ctx, contractId, personId)
}

func (cs *ContractService) AddLegalPersonToContract(ctx context.Context, contractId uint, legalPersonId uint) error {
	return cs.contractLegalPersonRepo.Save(ctx, &models.ContractLegalPerson{
		ContractID:    contractId,
		LegalPersonID: legalPersonId,
	})
}

func (cs *ContractService) RemoveLegalPersonFromContract(ctx context.Context, contractId uint, legalPersonId uint) error {
	return cs.contractLegalPersonRepo.Delete(ctx, contractId, legalPersonId)
}
