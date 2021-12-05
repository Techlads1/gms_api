package grievance_filling_mode

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceFillingMode, error)
	List() ([]*entity.GrievanceFillingMode, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceFillingMode) (int, error)
	Update(e *entity.GrievanceFillingMode) (int, error)
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
	ShowGrievanceFillingMode(id int) (*entity.GrievanceFillingMode, error)
	ListGrievanceFillingMode() ([]*entity.GrievanceFillingMode, error)
	StoreGrievanceFillingMode(name string, description string, CodeName string) (int, error)
	UpdateGrievanceFillingMode(emp *entity.GrievanceFillingMode) (int, error)
	DeleteGrievanceFillingMode(id int) error
}
