package grievance_state_transition

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceStateTransition, error)
	List() ([]*entity.GrievanceStateTransition, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceStateTransition) (int, error)
	Update(e *entity.GrievanceStateTransition) (int, error)
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
	ShowGrievanceStateTransition(id int) (*entity.GrievanceStateTransition, error)
	ListGrievanceStateTransition() ([]*entity.GrievanceStateTransition, error)
	StoreGrievanceStateTransition(Description string, FromStateId int, ToStateId int) (int, error)
	UpdateGrievanceStateTransition(emp *entity.GrievanceStateTransition) (int, error)
	DeleteGrievanceStateTransition(id int) error
}
