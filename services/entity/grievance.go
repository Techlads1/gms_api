package entity

import (
	"errors"
	"time"
)



type Grievance struct {
	Id                    	int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       							string    `json:"name" form:"name" validate:"required"`
	ReferenceNumber       	string    `json:"reference_number" form:"reference_number"`
	Comment       					string    `json:"comment" form:"comment" validate:"required"`
	State       						string    `json:"state" form:"state" validate:"required"`
	Description 						string    `json:"description" form:"description" validate:"required"`
	LocationOccurred      	string    `json:"location_occurred" form:"location_occurred" validate:"required"`
	FillingModeId         	int       `json:"filling_mode_id,omitempty" form:"filling_mode_id" validate:"omitempty,numeric"`
	GrievanceSubCategoryId	int       `json:"grievance_sub_category_id,omitempty" form:"grievance_sub_category_id" validate:"omitempty,numeric"`
	GrievantId              int       `json:"grievant_id,omitempty" form:"grievant_id" validate:"omitempty,numeric"`
	GrievantGroupId         int       `json:"grievant_group_id,omitempty" form:"grievant_group_id" validate:"omitempty,numeric"`
	UpdatedAt            		time.Time `json:"updated_at,omitempty"`
	CreatedAt             	time.Time `json:"created_at,omitempty"`
}

func NewGrievance(Name , Description, ReferenceNumber, Comment, State, LocationOccurred string, 
	FillingModeId, GrievanceSubCategoryId, GrievantId,
	 GrievantGroupId int,  UpdateAt time.Time, CreatedAt time.Time) (*Grievance, error) {

	grievance_state := &Grievance{
		Name:       					 Name,
		Description: 					 Description,
		ReferenceNumber: 			 ReferenceNumber,
		Comment: 							 Comment,	
		State: 								 State,
		LocationOccurred:      LocationOccurred,
		FillingModeId:         FillingModeId,	
		GrievanceSubCategoryId: GrievanceSubCategoryId,	
		GrievantId:            GrievantId,
		GrievantGroupId:       GrievantGroupId,
		UpdatedAt:             UpdateAt,
		CreatedAt:             CreatedAt,
	}
	err := grievance_state.ValidateNewGrievance()
	if err != nil {
		return nil, err
	}
	return grievance_state, nil
}


func (dep *Grievance) ValidateUpdateGrievance() error {
	if dep.Id < 1 {
		return errors.New("invalid grievance State id, field is required")
	}
	return nil
}


func (dep *Grievance) ValidateNewGrievance() error {
	if dep.Name == "" {
		return errors.New("invalid grievance State name, field is required")
	}

	return nil
}
