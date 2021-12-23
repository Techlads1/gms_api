package entity

import (
	"errors"
	"time"
)

type GrievanceStateAction struct {
	Id                int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name              string    `json:"name" form:"name" validate:"required"`
	RolePerformAction string    `json:"role_perform_action" form:"role_perform_action" validate:"required"`
	StateId           int       `json:"state_id,omitempty" form:"state_id" validate:"omitempty,numeric"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}

func NewGrievanceStateAction(Name, RolePerformAction string, StateId int, UpdateAt time.Time, CreatedAt time.Time) (*GrievanceStateAction, error) {

	grievance_StateAction := &GrievanceStateAction{
		Name:              Name,
		RolePerformAction: RolePerformAction,
		StateId:           StateId,
		UpdatedAt:         UpdateAt,
		CreatedAt:         CreatedAt,
	}
	err := grievance_StateAction.ValidateNewGrievanceStateAction()
	if err != nil {
		return nil, err
	}
	return grievance_StateAction, nil
}

func (dep *GrievanceStateAction) ValidateUpdateGrievanceStateAction() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance StateAction id, field is required")
	}
	return nil
}

func (dep *GrievanceStateAction) ValidateNewGrievanceStateAction() error {
	if dep.Name == "" {
		return errors.New("invalid grievance StateAction name, field is required")
	}

	if dep.StateId < 1 {
		return errors.New("invalid grievance StateAction state id, field is required")
	}

	if dep.RolePerformAction == "" {
		return errors.New("invalid grievance StateAction role perform action, field is required")
	}

	return nil
}
