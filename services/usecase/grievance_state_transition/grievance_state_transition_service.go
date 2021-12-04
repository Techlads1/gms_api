package grievance_state_transition

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
	repo := repository.NewGrievanceStateTransitionRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievanceStateTransition(Name string, FromStateId int, ToStateId int, Days int) (int, error) {
	grievance_StateTransition, err := entity.NewGrievanceStateTransition(Name, FromStateId, ToStateId, Days, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_StateTransition)
	if err != nil {
		log.Errorf("error creating GrievanceStateTransition: %v", err)
		return 0, errors.New("error creating GrievanceStateTransition")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceStateTransition(id int) (*entity.GrievanceStateTransition, error) {
	grievance_StateTransition, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceStateTransition: %v", err)
		return grievance_StateTransition, err
	}
	return grievance_StateTransition, nil
}


func (s *Service) ListGrievanceStateTransition() ([]*entity.GrievanceStateTransition, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceStateTransition not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievanceStateTransition(id int) error {
	_, err := s.ShowGrievanceStateTransition(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceStateTransition")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceStateTransition(dep *entity.GrievanceStateTransition) (int, error) {
	err := dep.ValidateUpdateGrievanceStateTransition()
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
