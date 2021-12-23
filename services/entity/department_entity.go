package entity

import (
	"errors"
	"time"
)

//Department DataStructure
type Department struct {
	Id                    int
	DepartmentTitle       string
	DepartmentDescription string
	DepartmentSize        int
	CampusId              int
	CreatedBy             int
	CreatedAt             time.Time
	UpdatedBy             int
	UpdatedAt             time.Time
	DeletedBy             int
	DeletedAt             time.Time
}

//NewDepartment for validating if department entity is ready for insertion
func NewDepartment(departmentTitle, departmentDescription string,
	departmentSize, campusId, createdBy int, createdAt time.Time) (*Department, error) {

	department := &Department{
		DepartmentTitle:       departmentTitle,
		DepartmentDescription: departmentDescription,
		DepartmentSize:        departmentSize,
		CampusId:              campusId,
		CreatedBy:             createdBy,
		CreatedAt:             createdAt,
	}
	err := department.ValidateNewDepartment()
	if err != nil {
		return nil, err
	}
	return department, nil
}

//ValidateUpdateDepartment validates on updating department
func (dep *Department) ValidateUpdateDepartment() error {
	if dep.Id < 1 {
		return errors.New("invalid department id, field is required")
	}
	return nil
}

//ValidateNewDepartment validates on inserting new department
func (dep *Department) ValidateNewDepartment() error {
	if dep.DepartmentTitle == "" {
		return errors.New("invalid department title, field is required")
	}
	return nil
}
