package department

import (
	"errors"

	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/entity"
	"github.com/tzdit/sample_api/services/repository"

	"time"
)

//Service department usecase
type Service struct {
	repo Repository
}

//NewService Create New Service
func NewService() *Service {
	repo := repository.NewDepartment()

	return &Service{
		repo: repo,
	}
}

//CreateDepartment creates new department
func (s *Service) CreateDepartment(departmentTitle, departmentDescription string,
	departmentSize, campusId, createdBy int) (int, error) {
	department, err := entity.NewDepartment(departmentTitle, departmentDescription, departmentSize, campusId,
		createdBy, time.Now())
	if err != nil {
		log.Error(err)
		return 0, err
	}

	departmentID, err := s.repo.Create(department)
	if err != nil {
		log.Errorf("error creating department: %v", err)
		return 0, errors.New("error creating department")
	}
	return departmentID, nil
}

//GetDepartment retrieve single department
func (s *Service) GetDepartment(id int) (*entity.Department, error) {
	department, err := s.repo.Get(id)
	if err != nil {
		log.Errorf("error getting department: %v", err)
		return department, err
	}
	return department, nil
}

//ListDepartment lists departments
func (s *Service) ListDepartment() ([]*entity.Department, error) {
	departments, err := s.repo.List()
	if err != nil {
		log.Error(err)
		return nil, errors.New("department not found")
	}
	return departments, err
}

//DeleteDepartment delete department
func (s *Service) DeleteDepartment(id, deletedBy int) error {
	_, err := s.GetDepartment(id)
	if err != nil {
		log.Error(err)
		return errors.New("error deleting department")
	}
	return s.repo.Delete(id, deletedBy)
}

//UpdateDepartment update department
func (s *Service) UpdateDepartment(dep *entity.Department) (int, error) {
	err := dep.ValidateUpdateDepartment()
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
