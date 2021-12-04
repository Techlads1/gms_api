package grievance_sub_category

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
	repo := repository.NewGrievanceSubCategoryRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievanceSubCategory(Name, Description, CodeName string) (int, error) {
	grievanceSub_category, err := entity.NewGrievanceSubCategory(Name, Description, CodeName, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievanceSub_category)
	if err != nil {
		log.Errorf("error creating GrievanceSubCategory: %v", err)
		return 0, errors.New("error creating GrievanceSubCategory")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceSubCategory(id int) (*entity.GrievanceSubCategory, error) {
	grievanceSub_category, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceSubCategory: %v", err)
		return grievanceSub_category, err
	}
	return grievanceSub_category, nil
}


func (s *Service) ListGrievanceSubCategory() ([]*entity.GrievanceSubCategory, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceSubCategory not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievanceSubCategory(id int) error {
	_, err := s.ShowGrievanceSubCategory(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceSubCategory")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceSubCategory(dep *entity.GrievanceSubCategory) (int, error) {
	err := dep.ValidateUpdateGrievanceSubCategory()
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
