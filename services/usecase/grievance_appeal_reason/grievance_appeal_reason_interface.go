package grievance_appeal_reason

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceAppealReason, error)
	List() ([]*entity.GrievanceAppealReason, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceAppealReason) (int, error)
	Update(e *entity.GrievanceAppealReason) (int, error)
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
	ShowGrievanceAppealReason(id int) (*entity.GrievanceAppealReason, error)
	ListGrievanceAppealReason() ([]*entity.GrievanceAppealReason, error)
	StoreGrievanceAppealReason(name string, description string) (int, error)
	UpdateGrievanceAppealReason(emp *entity.GrievanceAppealReason) (int, error)
	DeleteGrievanceAppealReason(id int) error
}
