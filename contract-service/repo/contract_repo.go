package repo

import (
	"context"
	"contract-service/models"
)

type ContractRepository interface {
	Save(ctx context.Context, contract *models.Contract) error
	Get(ctx context.Context, id uint) (*models.Contract, error)
	Delete(ctx context.Context, id uint) error
}
