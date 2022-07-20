package grievance_time_extension

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceTimeExtension, error)
	List() ([]*entity.GrievanceTimeExtension, error)
	ListNew() ([]*entity.GrievanceTimeExtension, error)
	ListApproved() ([]*entity.GrievanceTimeExtension, error)
	ListDenied() ([]*entity.GrievanceTimeExtension, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceTimeExtension) (int, error)
	Update(e *entity.GrievanceTimeExtension) (int, error)
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
	ShowGrievanceTimeExtension(id int) (*entity.GrievanceTimeExtension, error)
	ListGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error)
	ListNewGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error)
	ListApprovedGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error)
	ListDeniedGrievanceTimeExtension() ([]*entity.GrievanceTimeExtension, error)
	StoreGrievanceTimeExtension(ReferenceNumber string, State string, description string) (int, error)
	UpdateGrievanceTimeExtension(emp *entity.GrievanceTimeExtension) (int, error)
	DeleteGrievanceTimeExtension(id int) error
}
