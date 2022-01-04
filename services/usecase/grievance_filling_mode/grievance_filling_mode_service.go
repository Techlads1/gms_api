package grievance_filling_mode

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
	repo := repository.NewGrievanceFillingModeRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievanceFillingMode(Name, Description, CodeName string) (int, error) {
	grievance_filling_mode, err := entity.NewGrievanceFillingMode(Name, Description, CodeName, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_filling_mode)
	if err != nil {
		log.Errorf("error creating GrievanceFillingMode: %v", err)
		return 0, errors.New("error creating GrievanceFillingMode")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceFillingMode(id int) (*entity.GrievanceFillingMode, error) {
	grievance_filling_mode, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceFillingMode: %v", err)
		return grievance_filling_mode, err
	}
	return grievance_filling_mode, nil
}


func (s *Service) ListGrievanceFillingMode() ([]*entity.GrievanceFillingMode, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceFillingMode not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievanceFillingMode(id int) error {
	_, err := s.ShowGrievanceFillingMode(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceFillingMode")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceFillingMode(dep *entity.GrievanceFillingMode) (int, error) {
	err := dep.ValidateUpdateGrievanceFillingMode()
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
