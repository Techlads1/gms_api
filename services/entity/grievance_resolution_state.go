package entity

import (
	"errors"
	"time"
)

type GrievanceResolutionState struct {
	Id                    		int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	Description 							string    `json:"description" form:"description" validate:"required"`
	UpdatedAt             		time.Time `json:"updated_at,omitempty"`
	CreatedAt             		time.Time `json:"created_at,omitempty"`
}


func NewGrievanceResolutionState(Name , Description string, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceResolutionState, error) {

	grievance_ResolutionState := &GrievanceResolutionState{
		Name:       					 Name,
		Description: 					 Description,
		UpdatedAt:             UpdateAt,
		CreatedAt:             CreatedAt,
	}
	err := grievance_ResolutionState.ValidateNewGrievanceResolutionState()
	if err != nil {
		return nil, err
	}
	return grievance_ResolutionState, nil
}


func (dep *GrievanceResolutionState) ValidateUpdateGrievanceResolutionState() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance ResolutionState id, field is required")
	}
	return nil
}


func (dep *GrievanceResolutionState) ValidateNewGrievanceResolutionState() error {
	if dep.Name == "" {
		return errors.New("invalid grievance ResolutionState name, field is required")
	}


	return nil
}