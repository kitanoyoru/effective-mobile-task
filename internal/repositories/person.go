package repositories

import (
	"github.com/kitanoyoru/effective-mobile-task/internal/dtos"
	"github.com/kitanoyoru/effective-mobile-task/internal/models"
	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{
		db,
	}
}

func (r *PersonRepository) Save(person *models.Person) error {
	return r.db.Save(person).Error
}

func (r *PersonRepository) FindByID(id int) (*models.Person, error) {
	var person models.Person

	if err := r.db.Where(&models.Person{ID: id}).Take(&person).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) FindAll() ([]*models.Person, error) {
	var persons []*models.Person

	if err := r.db.Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonRepository) DeleteByID(id int) error {
	return r.db.Delete(&models.Person{ID: id}).Error
}

func (r *PersonRepository) DeleteManyByID(ids []string) error {
	return r.db.Where("ID in (?)", ids).Delete(&models.Person{}).Error
}

func (r *PersonRepository) PatchByID(id int, d *dtos.PersonPatchDTO) error {
	p, err := r.FindByID(id)
	if err != nil {
		return err
	}

	p.MergeWithPatchDTO(d)

	return r.Save(p)
}
