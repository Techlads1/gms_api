package grievance_state_action

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceStateAction, error)
	List() ([]*entity.GrievanceStateAction, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceStateAction) (int, error)
	Update(e *entity.GrievanceStateAction) (int, error)
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
	ShowGrievanceStateAction(id int) (*entity.GrievanceStateAction, error)
	ListGrievanceStateAction() ([]*entity.GrievanceStateAction, error)
	StoreGrievanceStateAction(name string, role_perform_action string, state_id int) (int, error)
	UpdateGrievanceStateAction(emp *entity.GrievanceStateAction) (int, error)
	DeleteGrievanceStateAction(id int) error
}
