package grievance_time_extension

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
	repo := repository.NewGrievanceTimeExtensionRepository()

	return &Service{
		repo: repo,
	}
}



func (s *Service) StoreGrievanceTimeExtension(GrievanceId, GFUId int, State, Description, Comment string) (int, error) {
	grievance_TimeExtension, err := entity.NewGrievanceTimeExtension(GrievanceId, GFUId, State, Description, Comment, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_TimeExtension)
	if err != nil {
		log.Errorf("error creating Grievance Time Extension: %v", err)
		return 0, errors.New("error creating Grievance Time Extension")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievanceTimeExtension(id int) (*entity.GrievanceTimeExtension, error) {
	grievance_TimeExtension, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting Grievance Time Extension: %v", err)
		return grievance_TimeExtension, err
	}
	return grievance_TimeExtension, nil
}

func (s *Service) ListGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error) {
	grievance_TimeExtension, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceTimeExtension not found")
	}
	return grievance_TimeExtension, err
}

func (s *Service) ListNewGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error) {
	grievance_TimeExtension, err := s.repo.ListNew()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceTimeExtension not found")
	}
	return grievance_TimeExtension, err
}


func (s *Service) ListApprovedGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error) {
	grievance_TimeExtension, err := s.repo.ListApproved()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceTimeExtension not found")
	}
	return grievance_TimeExtension, err
}

func (s *Service) ListDeniedGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error) {
	grievance_TimeExtension, err := s.repo.ListDenied()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceTimeExtension not found")
	}
	return grievance_TimeExtension, err
}


func (s *Service) DeleteGrievanceTimeExtension(id int) error {
	_, err := s.ShowGrievanceTimeExtension(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceTimeExtension")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceTimeExtension(dep *entity.GrievanceTimeExtension) (int, error) {
	err := dep.ValidateUpdateGrievanceTimeExtension()
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
