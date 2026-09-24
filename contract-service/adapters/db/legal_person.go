package sqlitedb

import (
	"context"
	"contract-service/models"
	"contract-service/repo"

	"gorm.io/gorm"
)

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
