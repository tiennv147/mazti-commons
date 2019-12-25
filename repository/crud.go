package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	commonerror "github.com/tiennv147/mazti-commons/errors"
)

type CrudRepo interface {
	DB() *gorm.DB

	Create(resource interface{}) error
	Get(resource interface{}, ID uint) error
	Delete(resource interface{}) error
	Update(resource interface{}) error
	List(resources interface{}, offset int, limit int) error
	Count(resource interface{}) (int64, error)
}

type crudRepo struct {
	db *gorm.DB
}

func NewCrudRepository(db *gorm.DB) CrudRepo {
	return &crudRepo{
		db: db,
	}
}

func (repo *crudRepo) DB() *gorm.DB {
	return repo.db
}

func (repo *crudRepo) Create(resource interface{}) error {
	err := repo.db.Create(resource).Error
	if err != nil {
		return errors.Wrap(err, "crud.Create")
	}
	return nil
}

func (repo *crudRepo) Get(resource interface{}, ID uint) error {
	r := repo.db.First(resource, ID)
	if r.RecordNotFound() {
		return commonerror.ErrNotFound
	}
	if err := r.Error; err != nil {
		return errors.Wrap(err, "crud.Get")
	}
	return nil
}

func (repo *crudRepo) Delete(ID interface{}) error {
	r := repo.db.Delete(ID)
	if r.RowsAffected == 0 {
		return commonerror.ErrNotFound
	}
	if err := r.Error; err != nil {
		return errors.Wrap(err, "crud.Delete")
	}
	return nil
}

func (repo *crudRepo) Update(resource interface{}) error {
	r := repo.db.Save(resource)
	if err := r.Error; err != nil {
		return errors.Wrap(err, "crud.Update")
	}
	return nil
}

func (repo *crudRepo) List(resources interface{}, offset int, limit int) error {
	q := repo.db.Offset(offset)
	if limit > 0 {
		q = q.Limit(limit)
	}
	r := q.Find(resources)
	if err := r.Error; err != nil {
		return errors.Wrap(err, "crud.List")
	}
	return nil
}

func (repo *crudRepo) Count(resource interface{}) (int64, error) {
	var count int64
	r := repo.db.Model(resource).Count(&count)
	if err := r.Error; err != nil {
		return 0, errors.Wrap(err, "crud.Count")
	}
	return count, nil
}
