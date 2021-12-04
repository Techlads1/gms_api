package grievance_state

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceState, error)
	List() ([]*entity.GrievanceState, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceState) (int, error)
	Update(e *entity.GrievanceState) (int, error)
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
	ShowGrievanceState(id int) (*entity.GrievanceState, error)
	ListGrievanceState() ([]*entity.GrievanceState, error)
	StoreGrievanceState(name string, description string, CodeName string) (int, error)
	UpdateGrievanceState(emp *entity.GrievanceState) (int, error)
	DeleteGrievanceState(id int) error
}
