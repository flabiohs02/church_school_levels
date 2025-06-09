package repository

import (
	"church_school_levels/domain"
)

type SchoolLevelRepository interface {
	Save(sl *domain.SchoolLevel) error
	FindByID(id uint) (*domain.SchoolLevel, error)
	FindAll() ([]domain.SchoolLevel, error)
	Update(sl *domain.SchoolLevel) error
	Delete(id uint) error
}
