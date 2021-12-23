package grievance_appeal_reason

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
	repo := repository.NewGrievanceAppealReasonRepository()

	return &Service{
		repo: repo,
	}
}

func (s *Service) StoreGrievanceAppealReason(Name, Description string) (int, error) {
	grievance_AppealReason, err := entity.NewGrievanceAppealReason(Name, Description, time.Now(), time.Now())

	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Store(grievance_AppealReason)
	if err != nil {
		log.Errorf("error creating GrievanceAppealReason: %v", err)
		return 0, errors.New("error creating GrievanceAppealReason")
	}
	return departmentID, nil
}

func (s *Service) ShowGrievanceAppealReason(id int) (*entity.GrievanceAppealReason, error) {
	grievance_AppealReason, err := s.repo.Show(id)
	if err != nil {
		log.Errorf("error getting GrievanceAppealReason: %v", err)
		return grievance_AppealReason, err
	}
	return grievance_AppealReason, nil
}

func (s *Service) ListGrievanceAppealReason() ([]*entity.GrievanceAppealReason, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("GrievanceAppealReason not found")
	}
	return departments, err
}

func (s *Service) DeleteGrievanceAppealReason(id int) error {
	_, err := s.ShowGrievanceAppealReason(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting GrievanceAppealReason")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateGrievanceAppealReason(dep *entity.GrievanceAppealReason) (int, error) {
	err := dep.ValidateUpdateGrievanceAppealReason()
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
