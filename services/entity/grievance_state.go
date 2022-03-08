package entity

import (
	"errors"
	"time"
)

type GrievanceState struct {
	Id                    int       `json:"id,omitempty" params:"id" form:"id" validate:"omitempty,numeric"`
	Name       						string    `json:"name" form:"name" validate:"required"`
	CodeName       				string    `json:"code_name" form:"code_name" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	Days                  int       `json:"days,omitempty" form:"days" validate:"omitempty,numeric"`
	SequenceNumber        int       `json:"sequence_number,omitempty" form:"sequence_number" validate:"omitempty,numeric"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}

func NewGrievanceState(Name , Description, CodeName string, Days int, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceState, error) {

	grievance_state := &GrievanceState{
		Name:       					 Name,
		Description: 					 Description,
		CodeName:							 CodeName,
		Days:							 		 Days,
		UpdatedAt:             UpdateAt,
		CreatedAt:             CreatedAt,
	}
	err := grievance_state.ValidateNewGrievanceState()
	if err != nil {
		return nil, err
	}
	return grievance_state, nil
}


func (dep *GrievanceState) ValidateUpdateGrievanceState() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance State id, field is required")
	}
	return nil
}


func (dep *GrievanceState) ValidateNewGrievanceState() error {
	if dep.Name == "" {
		return errors.New("invalid grievance State name, field is required")
	}

	if dep.CodeName == "" {
		return errors.New("invalid grievance State code name, field is required")
	}
	return nil
}
