package grievant_group

import (
	"errors"

	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/entity"
	"github.com/tzdit/sample_api/services/repository"

	"time"
)

//Service GrievantGroup usecase
type Service struct {
	repo Repository
}

//NewService Create New Service
func NewService() *Service {
	repo := repository.NewGrievantGroupRepository()

	return &Service{
		repo: repo,
	}
}

//CreateGrievantGroup creates new GrievantGroup
func (s *Service) StoreGrievantGroup(Name, Description string, GrievantCategoryId int) (int, error) {
	grievant_Group, err := entity.NewGrievantGroup(Name, Description, GrievantCategoryId, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievant_Group)
	if err != nil {
		log.Errorf("error creating GrievantGroup: %v", err)
		return 0, errors.New("error creating GrievantGroup")
	}
	return departmentID, nil
}

//GetGrievantGroup retrieve single GrievantGroup
func (s *Service) ShowGrievantGroup(id int) (*entity.GrievantGroup, error) {
	grievant_Group, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievantGroup: %v", err)
		return grievant_Group, err
	}
	return grievant_Group, nil
}

//ListGrievantCategories lists GrievantCategories
func (s *Service) ListGrievantGroup() ([]*entity.GrievantGroup, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievantGroup not found")
	}
	return departments, err
}

//DeleteGrievantGroup delete GrievantGroup
func (s *Service) DeleteGrievantGroup(id int) error {
	_, err := s.ShowGrievantGroup(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievantGroup")
	}
	return s.repo.Delete(id)
}

//UpdateDepartment update department
func (s *Service) UpdateGrievantGroup(dep *entity.GrievantGroup) (int, error) {
	err := dep.ValidateUpdateGrievantGroup()
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
