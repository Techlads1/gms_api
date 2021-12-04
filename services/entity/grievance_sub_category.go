package entity

import (
	"errors"
	"time"
)

type GrievanceSubCategory struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	CodeName       				string    `json:"code_name" form:"code_name" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}

func NewGrievanceSubCategory(Name , Description string, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceSubCategory, error) {

	grievance_sub_category := &GrievanceSubCategory{
		Name:       					 Name,
		Description: 					 Description,
		UpdatedAt:             UpdateAt,
		CreatedAt:             CreatedAt,
	}
	err := grievance_sub_category.ValidateNewGrievanceSubCategory()
	if err != nil {
		return nil, err
	}
	return grievance_sub_category, nil
}


func (dep *GrievanceSubCategory) ValidateUpdateGrievanceSubCategory() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance sub category id, field is required")
	}
	return nil
}


func (dep *GrievanceSubCategory) ValidateNewGrievanceSubCategory() error {
	if dep.Name == "" {
		return errors.New("invalid grievance sub category name, field is required")
	}
	return nil
}
