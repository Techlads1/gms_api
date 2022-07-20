package grievance_forward

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceForward, error)
	List() ([]*entity.GrievanceForward, error)
	ListNew() ([]*entity.GrievanceForward, error)
	ListApproved() ([]*entity.GrievanceForward, error)
	ListDenied() ([]*entity.GrievanceForward, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceForward) (int, error)
	Update(e *entity.GrievanceForward) (int, error)
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
	ShowGrievanceForward(id int) (*entity.GrievanceForward, error)
	ListGrievanceForward() ([]*entity.GrievanceForward, error)
	ListNewGrievanceForward() ([]*entity.GrievanceForward, error)
	ListApprovedGrievanceForward() ([]*entity.GrievanceForward, error)
	ListDeniedGrievanceForward() ([]*entity.GrievanceForward, error)
	StoreGrievanceForward(ReferenceNumber string, State string, description string) (int, error)
	UpdateGrievanceForward(emp *entity.GrievanceForward) (int, error)
	DeleteGrievanceForward(id int) error
}
