package entity

import (
	"errors"
	"time"
)

type GrievanceStateTransition struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	FromStateId      			int    		`json:"from_state_id" form:"from_state_id" validate:"required"`
	ToStateId      				int    		`json:"to_state_id" form:"to_state_id" validate:"required"`
	Description 					string    `json:"description" form:"description" validate:"required"`
	Days                  int       `json:"days,omitempty" form:"days" validate:"omitempty,numeric"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}


func NewGrievanceStateTransition( Description string, FromStateId int, ToStateId int, Days int, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceStateTransition, error) {


	grievance_StateTransition := &GrievanceStateTransition{
		FromStateId:       		 FromStateId,
		Description: 					 Description,
		ToStateId:						 ToStateId,
		Days:                  Days,
		UpdatedAt:             UpdateAt,
		CreatedAt:             CreatedAt,
	}
	err := grievance_StateTransition.ValidateNewGrievanceStateTransition()
	if err != nil {
		return nil, err
	}
	return grievance_StateTransition, nil
}


func (dep *GrievanceStateTransition) ValidateUpdateGrievanceStateTransition() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance StateTransition id, field is required")
	}
	return nil
}


func (dep *GrievanceStateTransition) ValidateNewGrievanceStateTransition() error {
	if dep.Description == "" {
		return errors.New("invalid grievance StateTransition name, field is required")
	}

	
	return nil
}






