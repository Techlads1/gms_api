package entity

import (
	"errors"
	"time"
)

type GrievanceForward struct {
	Id              int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	GrievanceId int    `json:"grievance_id" form:"grievance_id"`
	State           string    `json:"state" form:"state" validate:"required"`
	FromGFUId		int	  	  `json:"fromgfu_id" form:"fromgfu_id" validate:"required"`
	ToGFUId			int	  	  `json:"togfu_id" form:"togfu_id" validate:"required"`
	Description     string    `json:"description" form:"description" validate:"required"`
	Comment       	string    `json:"comment" form:"comment" validate:"required"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

func NewGrievanceForward(GrievanceId int, State string, FromGFUId int, ToGFUId int, Description string, Comment string,UpdateAt time.Time, CreatedAt time.Time) (*GrievanceForward, error) {

	grievance_Forward := &GrievanceForward{
		GrievanceId: GrievanceId,
		State:           State,
		FromGFUId:     	 FromGFUId,
		ToGFUId:     	 ToGFUId,
		Description:     Description,
		Comment:     	 Comment,
		UpdatedAt:       UpdateAt,
		CreatedAt:       CreatedAt,
	}
	err := grievance_Forward.ValidateNewGrievanceForward()
	if err != nil {
		return nil, err
	}
	return grievance_Forward, nil
}

func (dep *GrievanceForward) ValidateUpdateGrievanceForward() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance Forward id, field is required")
	}
	return nil
}

func (dep *GrievanceForward) ValidateNewGrievanceForward() error {
	if dep.Description == "" {
		return errors.New("invalid grievance description, field is required")
	}
	return nil
}