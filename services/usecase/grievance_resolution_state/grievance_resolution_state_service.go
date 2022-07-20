package grievance_resolution_state

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
	repo := repository.NewGrievanceResolutionStateRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievanceResolutionState(Name, Description string) (int, error) {
	grievance_ResolutionState, err := entity.NewGrievanceResolutionState(Name, Description, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_ResolutionState)
	if err != nil {
		log.Errorf("error creating GrievanceResolutionState: %v", err)
		return 0, errors.New("error creating GrievanceResolutionState")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceResolutionState(id int) (*entity.GrievanceResolutionState, error) {
	grievance_ResolutionState, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceResolutionState: %v", err)
		return grievance_ResolutionState, err
	}
	return grievance_ResolutionState, nil
}


func (s *Service) ListGrievanceResolutionState() ([]*entity.GrievanceResolutionState, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceResolutionState not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievanceResolutionState(id int) error {
	_, err := s.ShowGrievanceResolutionState(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceResolutionState")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceResolutionState(dep *entity.GrievanceResolutionState) (int, error) {
	err := dep.ValidateUpdateGrievanceResolutionState()
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
