package entity

import (
	"errors"
	"time"
)

type GrievanceAppealReason struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}

func NewGrievanceAppealReason(Name , Description string, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceAppealReason, error) {

	grievance_AppealReason := &GrievanceAppealReason{
		Name:       					 Name,
		Description: 					 Description,
		UpdatedAt:             UpdateAt,
		CreatedAt:             CreatedAt,
	}
	err := grievance_AppealReason.ValidateNewGrievanceAppealReason()
	if err != nil {
		return nil, err
	}
	return grievance_AppealReason, nil
}


func (dep *GrievanceAppealReason) ValidateUpdateGrievanceAppealReason() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance AppealReason id, field is required")
	}
	return nil
}


func (dep *GrievanceAppealReason) ValidateNewGrievanceAppealReason() error {
	if dep.Name == "" {
		return errors.New("invalid grievance AppealReason name, field is required")
	}


	return nil
}
