package entity

import (
	"errors"
	"time"
)

type GrievantGroup struct {
	Id                    int       `json:"id,omitempty" params:"id" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required,max=200"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	GrievantCategoryId 		int				`json:"grievant_category_id" form:"grievant_category_id" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}

func NewGrievantGroup(Name , Description string, GrievantCategoryId int, UpdateAt time.Time, CreatedAt time.Time) (*GrievantGroup, error) {

	grievant_Group := &GrievantGroup{
		Name:       					 Name,
		Description: 					 Description,
		UpdatedAt:             UpdateAt,
		CreatedAt:             CreatedAt,
	}
	err := grievant_Group.ValidateNewGrievantGroup()
	if err != nil {
		return nil, err
	}
	return grievant_Group, nil
}


func (dep *GrievantGroup) ValidateUpdateGrievantGroup() error {
	if dep.Id < 1 {
		return errors.New("invalid grievant Group id, field is required")
	}
	return nil
}


func (dep *GrievantGroup) ValidateNewGrievantGroup() error {
	if dep.Name == ""  {
		return errors.New("invalid grievant Group name, field is required")
	}

	if dep.GrievantCategoryId <= 0  {
		return errors.New("invalid grievant category id, field is required")
	}

	return nil
}