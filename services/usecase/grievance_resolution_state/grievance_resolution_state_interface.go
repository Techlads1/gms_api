package grievance_resolution_state

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceResolutionState, error)
	List() ([]*entity.GrievanceResolutionState, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceResolutionState) (int, error)
	Update(e *entity.GrievanceResolutionState) (int, error)
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
	ShowGrievanceResolutionState(id int) (*entity.GrievanceResolutionState, error)
	ListGrievanceResolutionState() ([]*entity.GrievanceResolutionState, error)
	StoreGrievanceResolutionState(name string, description string) (int, error)
	UpdateGrievanceResolutionState(emp *entity.GrievanceResolutionState) (int, error)
	DeleteGrievanceResolutionState(id int) error
}
