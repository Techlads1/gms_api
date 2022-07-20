package grievance_category

import (
	"errors"

	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/entity"
	"github.com/tzdit/sample_api/services/repository"

	"time"
)


type Service struct {
	repo Repository
}


func NewService() *Service {
	repo := repository.NewGrievanceCategoryRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievanceCategory(Name, Description, CodeName string) (int, error) {
	grievance_category, err := entity.NewGrievanceCategory(Name, Description, CodeName, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_category)
	if err != nil {
		log.Errorf("error creating GrievanceCategory: %v", err)
		return 0, errors.New("error creating GrievanceCategory")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceCategory(id int) (*entity.GrievanceCategory, error) {
	grievance_category, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceCategory: %v", err)
		return grievance_category, err
	}
	return grievance_category, nil
}


func (s *Service) ListGrievanceCategory() ([]*entity.GrievanceCategory, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceCategory not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievanceCategory(id int) error {
	_, err := s.ShowGrievanceCategory(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceCategory")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceCategory(dep *entity.GrievanceCategory) (int, error) {
	err := dep.ValidateUpdateGrievanceCategory()
	if err != nil {
		log.Error(err)
		return dep.Id, err
	}
	dep.UpdatedAt = time.Now()
	_, err = s.repo.Update(dep)
	if err != nil {
		log.Error(err)
		return dep.Id, err
	}
	return dep.Id, err
}
