package repository

import (
	"church_school_levels/domain"

	"gorm.io/gorm"
)

type GormSchoolLevelRepository struct {
	db *gorm.DB
}

func NewGormSchoolLevelRepository(db *gorm.DB) *GormSchoolLevelRepository {
	return &GormSchoolLevelRepository{db: db}
}

func (r *GormSchoolLevelRepository) Save(sl *domain.SchoolLevel) error {
	return r.db.Create(sl).Error
}

func (r *GormSchoolLevelRepository) FindByID(id uint) (*domain.SchoolLevel, error) {
	var sl domain.SchoolLevel
	err := r.db.First(&sl, id).Error
	return &sl, err
}

func (r *GormSchoolLevelRepository) FindAll() ([]domain.SchoolLevel, error) {
	var sls []domain.SchoolLevel
	err := r.db.Find(&sls).Error
	return sls, err
}

func (r *GormSchoolLevelRepository) Update(sl *domain.SchoolLevel) error {
	return r.db.Save(sl).Error
}

func (r *GormSchoolLevelRepository) Delete(id uint) error {
	return r.db.Delete(&domain.SchoolLevel{}, id).Error
}
