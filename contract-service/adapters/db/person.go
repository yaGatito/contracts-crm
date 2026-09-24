package sqlitedb

import (
	"context"
	"contract-service/models"
	"contract-service/repo"

	"gorm.io/gorm"
)

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
