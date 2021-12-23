package entity

import (
	"errors"
	"time"
)

type GrievantCategory struct {
	Id          int       `json:"id,omitempty" param:"id" form:"id" validate:"omitempty,numeric"`
	Name        string    `json:"name" form:"name" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func NewGrievantCategory(Name, Description string, UpdateAt time.Time, CreatedAt time.Time) (*GrievantCategory, error) {

	grievant_category := &GrievantCategory{
		Name:        Name,
		Description: Description,
		UpdatedAt:   UpdateAt,
		CreatedAt:   CreatedAt,
	}
	err := grievant_category.ValidateNewGrievantCategory()
	if err != nil {
		return nil, err
	}
	return grievant_category, nil
}

func (dep *GrievantCategory) ValidateUpdateGrievantCategory() error {
	if dep.Id < 1 {
		return errors.New("invalid grievant category id, field is required")
	}
	return nil
}

func (dep *GrievantCategory) ValidateNewGrievantCategory() error {
	if dep.Name == "" {
		return errors.New("invalid grievant category name, field is required")
	}
	return nil
}
