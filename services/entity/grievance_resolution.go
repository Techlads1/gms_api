package entity

import (
	"errors"
	"time"
)

type GrievanceResolution struct {
	Id              int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	GrievanceId 	int    	  `json:"grievance_id" form:"grievance_id"`
	GFUId 			int       `json:"gfu_id" form:"gfu_id"`
	State           string    `json:"state" form:"state" validate:"required"`
	Description     string    `json:"description" form:"description" validate:"required"`
	Comment       	string    `json:"comment" form:"comment" validate:"required"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

func NewGrievanceResolution(GrievanceId, GFUId int, State string, Description string, Comment string,UpdateAt time.Time, CreatedAt time.Time) (*GrievanceResolution, error) {

	grievance_Resolution := &GrievanceResolution{
		GrievanceId: 	 GrievanceId,
		GFUId: 	 		 GFUId,
		State:           State,
		Description:     Description,
		Comment:     	 Comment,
		UpdatedAt:       UpdateAt,
		CreatedAt:       CreatedAt,
	}
	err := grievance_Resolution.ValidateNewGrievanceResolution()
	if err != nil {
		return nil, err
	}
	return grievance_Resolution, nil
}

func (dep *GrievanceResolution) ValidateUpdateGrievanceResolution() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance resolution id, field is required")
	}
	if dep.GrievanceId < 1 {
		return errors.New("invalid grievance resolution grievanceId, field is required")
	}
	return nil
}

func (dep *GrievanceResolution) ValidateNewGrievanceResolution() error {
	if dep.Description == "" {
		return errors.New("invalid grievance Description name, field is required")
	}


	return nil
}