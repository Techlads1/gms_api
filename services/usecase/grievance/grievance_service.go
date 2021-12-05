package grievance

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
	repo := repository.NewGrievanceRepository()

	return &Service{
		repo: repo,
	}
}


func (s *Service) StoreGrievance(Name, Description, ReferenceNumber, Comment, State, LocationOccurred string, 
	FillingModeId, GrievanceSubCategoryId, GrievantId,
	 GrievantGroupId int,) (int, error) {
	grievance_State, err := entity.NewGrievance(Name, Description,ReferenceNumber, Comment, State, LocationOccurred , 
		FillingModeId, GrievanceSubCategoryId, GrievantId,
		 GrievantGroupId,  time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_State)
	if err != nil {
		log.Errorf("error creating Grievance: %v", err)
		return 0, errors.New("error creating Grievance")
	}
	return departmentID, nil
}


func (s *Service) ShowGrievance(id int) (*entity.Grievance, error) {
	grievance_State, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting Grievance: %v", err)
		return grievance_State, err
	}
	return grievance_State, nil
}


func (s *Service) ListGrievance() ([]*entity.Grievance, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("Grievance not found")
	}
	return departments, err
}


func (s *Service) DeleteGrievance(id int) error {
	_, err := s.ShowGrievance(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting Grievance")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievance(dep *entity.Grievance) (int, error) {
	err := dep.ValidateUpdateGrievance()
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
