package grievant

import (
	"errors"

	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/entity"
	"github.com/tzdit/sample_api/services/repository"

	"time"
)

//Service GrievantCategory usecase
type Service struct {
	repo Repository
}

//NewService Create New Service
func NewService() *Service {
	repo := repository.NewGrievantCategoryRepository()

	return &Service{
		repo: repo,
	}
}

//CreateGrievantCategory creates new GrievantCategory
func (s *Service) StoreGrievantCategory(Name, Description string) (int, error) {
	grievant_category, err := entity.NewGrievantCategory(Name, Description, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievant_category)
	if err != nil {
		log.Errorf("error creating GrievantCategory: %v", err)
		return 0, errors.New("error creating GrievantCategory")
	}
	return departmentID, nil
}

//GetGrievantCategory retrieve single GrievantCategory
func (s *Service) ShowGrievantCategory(id int) (*entity.GrievantCategory, error) {
	grievant_category, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievantCategory: %v", err)
		return grievant_category, err
	}
	return grievant_category, nil
}

//ListGrievantCategories lists GrievantCategories
func (s *Service) ListGrievantCategory() ([]*entity.GrievantCategory, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievantCategory not found")
	}
	return departments, err
}

//DeleteGrievantCategory delete GrievantCategory
func (s *Service) DeleteGrievantCategory(id int) error {
	_, err := s.ShowGrievantCategory(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievantCategory")
	}
	return s.repo.Delete(id)
}

//UpdateDepartment update department
func (s *Service) UpdateGrievantCategory(dep *entity.GrievantCategory) (int, error) {
	err := dep.ValidateUpdateGrievantCategory()
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
