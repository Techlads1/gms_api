package grievance_forward

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
	repo := repository.NewGrievanceForwardRepository()

	return &Service{
		repo: repo,
	}
}



func (s *Service) StoreGrievanceForward(GrievanceId int, State string,  FromGFUId int, ToGFUId int, Description, Comment string) (int, error) {
	grievance_Forward, err := entity.NewGrievanceForward(GrievanceId, State, FromGFUId, ToGFUId, Description, Comment, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_Forward)
	if err != nil {
		log.Errorf("error creating GrievanceForward: %v", err)
		return 0, errors.New("error creating GrievanceForward")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceForward(id int) (*entity.GrievanceForward, error) {
	grievance_Forward, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceForward: %v", err)
		return grievance_Forward, err
	}
	return grievance_Forward, nil
}

func (s *Service) ListGrievanceForward() ([]*entity.GrievanceForward, error) {
	grievance_Forward, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceForward not found")
	}
	return grievance_Forward, err
}

func (s *Service) ListNewGrievanceForward() ([]*entity.GrievanceForward, error) {
	grievance_Forward, err := s.repo.ListNew()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceForward not found")
	}
	return grievance_Forward, err
}


func (s *Service) ListApprovedGrievanceForward() ([]*entity.GrievanceForward, error) {
	grievance_Forward, err := s.repo.ListApproved()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceForward not found")
	}
	return grievance_Forward, err
}

func (s *Service) ListDeniedGrievanceForward() ([]*entity.GrievanceForward, error) {
	grievance_Forward, err := s.repo.ListDenied()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceForward not found")
	}
	return grievance_Forward, err
}


func (s *Service) DeleteGrievanceForward(id int) error {
	_, err := s.ShowGrievanceForward(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceForward")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceForward(dep *entity.GrievanceForward) (int, error) {
	err := dep.ValidateUpdateGrievanceForward()
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
