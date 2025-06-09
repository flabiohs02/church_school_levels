package usecase

import (
	"church_school_levels/domain"
	"church_school_levels/repository"
)

type SchoolLevelService struct {
	repo repository.SchoolLevelRepository
}

func NewSchoolLevelService(repo repository.SchoolLevelRepository) *SchoolLevelService {
	return &SchoolLevelService{repo: repo}
}

func (s *SchoolLevelService) CreateSchoolLevel(sl *domain.SchoolLevel) error {
	return s.repo.Save(sl)
}

func (s *SchoolLevelService) GetSchoolLevelByID(id uint) (*domain.SchoolLevel, error) {
	return s.repo.FindByID(id)
}

func (s *SchoolLevelService) GetAllSchoolLevels() ([]domain.SchoolLevel, error) {
	return s.repo.FindAll()
}

func (s *SchoolLevelService) UpdateSchoolLevel(sl *domain.SchoolLevel) error {
	return s.repo.Update(sl)
}

func (s *SchoolLevelService) DeleteSchoolLevel(id uint) error {
	return s.repo.Delete(id)
}
