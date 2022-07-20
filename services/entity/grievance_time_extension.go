package entity

import (
	"errors"
	"time"
)

type GrievanceTimeExtension struct {
	Id              int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	GrievanceId 	int    	  `json:"grievance_id" form:"grievance_id"`
	GFUId 			int       `json:"gfu_id" form:"gfu_id"`
	State           string    `json:"state" form:"state" validate:"required"`
	Description     string    `json:"description" form:"description" validate:"required"`
	Comment       	string    `json:"comment" form:"comment" validate:"required"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

func NewGrievanceTimeExtension(GrievanceId, GFUId int, State string, Description string, Comment string,UpdateAt time.Time, CreatedAt time.Time) (*GrievanceTimeExtension, error) {

	grievance_TimeExtension := &GrievanceTimeExtension{
		GrievanceId: 	 GrievanceId,
		GFUId: 	 		 GFUId,
		State:           State,
		Description:     Description,
		Comment:     	 Comment,
		UpdatedAt:       UpdateAt,
		CreatedAt:       CreatedAt,
	}
	err := grievance_TimeExtension.ValidateNewGrievanceTimeExtension()
	if err != nil {
		return nil, err
	}
	return grievance_TimeExtension, nil
}

func (dep *GrievanceTimeExtension) ValidateUpdateGrievanceTimeExtension() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance Time Extension id, field is required")
	}
	if dep.GrievanceId < 1 {
		return errors.New("invalid grievance Time Extension grievanceId, field is required")
	}
	return nil
}

func (dep *GrievanceTimeExtension) ValidateNewGrievanceTimeExtension() error {
	if dep.Description == "" {
		return errors.New("invalid grievance Description name, field is required")
	}


	return nil
}