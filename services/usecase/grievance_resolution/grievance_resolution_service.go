package grievance_resolution

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
	repo := repository.NewGrievanceResolutionRepository()

	return &Service{
		repo: repo,
	}
}



func (s *Service) StoreGrievanceResolution(GrievanceId, GFUId int, State, Description, Comment string) (int, error) {
	grievance_Resolution, err := entity.NewGrievanceResolution(GrievanceId, GFUId, State, Description, Comment, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_Resolution)
	if err != nil {
		log.Errorf("error creating GrievanceResolution: %v", err)
		return 0, errors.New("error creating GrievanceResolution")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceResolution(id int) (*entity.GrievanceResolution, error) {
	grievance_Resolution, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceResolution: %v", err)
		return grievance_Resolution, err
	}
	return grievance_Resolution, nil
}

func (s *Service) ListGrievanceResolution() ([]*entity.GrievanceResolution, error) {
	grievance_Resolution, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceResolution not found")
	}
	return grievance_Resolution, err
}

func (s *Service) ListNewGrievanceResolution() ([]*entity.GrievanceResolution, error) {
	grievance_Resolution, err := s.repo.ListNew()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceResolution not found")
	}
	return grievance_Resolution, err
}


func (s *Service) ListApprovedGrievanceResolution() ([]*entity.GrievanceResolution, error) {
	grievance_Resolution, err := s.repo.ListApproved()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceResolution not found")
	}
	return grievance_Resolution, err
}

func (s *Service) ListDeniedGrievanceResolution() ([]*entity.GrievanceResolution, error) {
	grievance_Resolution, err := s.repo.ListDenied()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceResolution not found")
	}
	return grievance_Resolution, err
}


func (s *Service) DeleteGrievanceResolution(id int) error {
	_, err := s.ShowGrievanceResolution(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceResolution")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceResolution(dep *entity.GrievanceResolution) (int, error) {
	err := dep.ValidateUpdateGrievanceResolution()
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
