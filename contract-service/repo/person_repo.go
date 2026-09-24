package repo

import (
	"context"
	"contract-service/models"
)

type LegalPersonRepository interface {
	Save(ctx context.Context, person *models.LegalPerson) error
	Get(ctx context.Context, id uint) (*models.LegalPerson, error)
	Delete(ctx context.Context, id uint) error
}

type PersonRepository interface {
	Save(ctx context.Context, person *models.Person) error
	Get(ctx context.Context, id uint) (*models.Person, error)
	Delete(ctx context.Context, id uint) error
}
