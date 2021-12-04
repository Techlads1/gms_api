package grievance_state_action

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
	repo := repository.NewGrievanceStateActionRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievanceStateAction(Name string, RolePerformAction string, StateId int) (int, error) {
	grievance_StateAction, err := entity.NewGrievanceStateAction(Name, RolePerformAction, StateId, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_StateAction)
	if err != nil {
		log.Errorf("error creating GrievanceStateAction: %v", err)
		return 0, errors.New("error creating GrievanceStateAction")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceStateAction(id int) (*entity.GrievanceStateAction, error) {
	grievance_StateAction, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceStateAction: %v", err)
		return grievance_StateAction, err
	}
	return grievance_StateAction, nil
}


func (s *Service) ListGrievanceStateAction() ([]*entity.GrievanceStateAction, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceStateAction not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievanceStateAction(id int) error {
	_, err := s.ShowGrievanceStateAction(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceStateAction")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceStateAction(dep *entity.GrievanceStateAction) (int, error) {
	err := dep.ValidateUpdateGrievanceStateAction()
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
