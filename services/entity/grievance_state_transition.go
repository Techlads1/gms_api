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
	SequenceNumber        int       `json:"sequence_number,omitempty" form:"sequence_number" validate:"omitempty,numeric"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}


func NewGrievanceStateTransition( Description string, FromStateId int, ToStateId int,UpdateAt time.Time, CreatedAt time.Time) (*GrievanceStateTransition, error) {


	grievance_StateTransition := &GrievanceStateTransition{
		FromStateId:       		 FromStateId,
		Description: 					 Description,
		ToStateId:						 ToStateId,
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

	if dep.ToStateId == dep.FromStateId {
		return errors.New("invalid grievance State Transition can not transite to it self")
	}

	return nil
}


func (dep *GrievanceStateTransition) ValidateNewGrievanceStateTransition() error {
	if dep.Description == "" {
		return errors.New("invalid grievance StateTransition name, field is required")
	}

	if dep.ToStateId == dep.FromStateId {
		return errors.New("invalid grievance State Transition can not transite to it self")
	}
	
	return nil
}






