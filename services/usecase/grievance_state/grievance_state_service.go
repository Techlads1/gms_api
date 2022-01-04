package grievance_state

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
	repo := repository.NewGrievanceStateRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievanceState(Name, Description, CodeName string, Days int) (int, error) {
	grievance_State, err := entity.NewGrievanceState(Name, Description, CodeName, Days, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_State)
	if err != nil {
		log.Errorf("error creating GrievanceState: %v", err)
		return 0, errors.New("error creating GrievanceState")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceState(id int) (*entity.GrievanceState, error) {
	grievance_State, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceState: %v", err)
		return grievance_State, err
	}
	return grievance_State, nil
}


func (s *Service) ListGrievanceState() ([]*entity.GrievanceState, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceState not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievanceState(id int) error {
	_, err := s.ShowGrievanceState(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceState")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceState(dep *entity.GrievanceState) (int, error) {
	err := dep.ValidateUpdateGrievanceState()
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
