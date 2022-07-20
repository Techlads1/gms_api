package grievance_resolution

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceResolution, error)
	List() ([]*entity.GrievanceResolution, error)
	ListNew() ([]*entity.GrievanceResolution, error)
	ListApproved() ([]*entity.GrievanceResolution, error)
	ListDenied() ([]*entity.GrievanceResolution, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceResolution) (int, error)
	Update(e *entity.GrievanceResolution) (int, error)
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
	ShowGrievanceResolution(id int) (*entity.GrievanceResolution, error)
	ListGrievanceResolution() ([]*entity.GrievanceResolution, error)
	ListNewGrievanceResolution() ([]*entity.GrievanceResolution, error)
	ListApprovedGrievanceResolution() ([]*entity.GrievanceResolution, error)
	ListDeniedGrievanceResolution() ([]*entity.GrievanceResolution, error)
	StoreGrievanceResolution(ReferenceNumber string, State string, description string) (int, error)
	UpdateGrievanceResolution(emp *entity.GrievanceResolution) (int, error)
	DeleteGrievanceResolution(id int) error
}
