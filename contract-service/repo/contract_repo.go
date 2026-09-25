package repo

import (
	"context"
	"contract-service/models"
)

type ContractRepository interface {
	Save(ctx context.Context, contract *models.Contract) error
	Get(ctx context.Context, id uint) (*models.Contract, error)
	Update(ctx context.Context, contract *models.Contract) error
	Delete(ctx context.Context, id uint) error
	GetContractDetails(ctx context.Context, id uint) (models.ContractDetails, error)

	SearchLegalPersonContracts(
		ctx context.Context,
		filter models.SearchLegalPersonContractsFilter,
	) ([]models.Contract, error)

	SearchPersonContracts(
		ctx context.Context,
		filter models.SearchPersonContractsFilter,
	) ([]models.Contract, error)
}
