package grievance

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.Grievance, error)
	List() ([]*entity.Grievance, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.Grievance) (int, error)
	Update(e *entity.Grievance) (int, error)
	Delete(id int) error
}

//DONE
//Repository interface
type Repository interface {
	Reader
	Writer
}

//DONE
//UseCase interface
type UseCase interface {
	ShowGrievance(id int) (*entity.Grievance, error)
	ListGrievance() ([]*entity.Grievance, error)
	StoreGrievance(name string, description, ReferenceNumber, Comment, State, LocationOccurred string, 
		FillingModeId, GrievanceSubCategoryId, GrievantId,
		 GrievantGroupId int,) (int, error)
	UpdateGrievance(emp *entity.Grievance) (int, error)
	DeleteGrievance(id int) error
}
