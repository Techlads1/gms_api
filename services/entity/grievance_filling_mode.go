package entity

import (
	"errors"
	"time"
)

type GrievanceFillingMode struct {
	Id          int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name        string    `json:"name" form:"name" validate:"required"`
	CodeName    string    `json:"code_name" form:"code_name" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func NewGrievanceFillingMode(Name, Description, CodeName string, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceFillingMode, error) {

	grievance_FillingMode := &GrievanceFillingMode{
		Name:        Name,
		Description: Description,
		CodeName:    CodeName,
		UpdatedAt:   UpdateAt,
		CreatedAt:   CreatedAt,
	}
	err := grievance_FillingMode.ValidateNewGrievanceFillingMode()
	if err != nil {
		return nil, err
	}
	return grievance_FillingMode, nil
}

func (dep *GrievanceFillingMode) ValidateUpdateGrievanceFillingMode() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance FillingMode id, field is required")
	}
	return nil
}

func (dep *GrievanceFillingMode) ValidateNewGrievanceFillingMode() error {
	if dep.Name == "" {
		return errors.New("invalid grievance FillingMode name, field is required")
	}

	if dep.CodeName == "" {
		return errors.New("invalid grievance FillingMode code name, field is required")
	}
	return nil
}
