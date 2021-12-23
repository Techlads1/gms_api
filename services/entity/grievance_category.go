package entity

import (
	"errors"
	"time"
)

type GrievanceCategory struct {
	Id          int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name        string    `json:"name" form:"name" validate:"required"`
	CodeName    string    `json:"code_name" form:"code_name" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type GrievanceCategoryRepository interface {
	Create(arg *GrievantCategory) (int, error)
	Get(id int) (*GrievantCategory, error)
	Update(arg *GrievantCategory) (int, error)
	All() ([]*GrievantCategory, error)
	Delete(id int) error
}

func NewGrievanceCategory(Name, Description, CodeName string, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceCategory, error) {

	grievance_category := &GrievanceCategory{
		Name:        Name,
		Description: Description,
		CodeName:    CodeName,
		UpdatedAt:   UpdateAt,
		CreatedAt:   CreatedAt,
	}
	err := grievance_category.ValidateNewGrievanceCategory()
	if err != nil {
		return nil, err
	}
	return grievance_category, nil
}

func (dep *GrievanceCategory) ValidateUpdateGrievanceCategory() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance category id, field is required")
	}
	return nil
}

func (dep *GrievanceCategory) ValidateNewGrievanceCategory() error {
	if dep.Name == "" {
		return errors.New("invalid grievance category name, field is required")
	}

	if dep.CodeName == "" {
		return errors.New("invalid grievance category code name, field is required")
	}
	return nil
}
