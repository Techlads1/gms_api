package grievance_category

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceCategory, error)
	List() ([]*entity.GrievanceCategory, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceCategory) (int, error)
	Update(e *entity.GrievanceCategory) (int, error)
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
	ShowGrievanceCategory(id int) (*entity.GrievanceCategory, error)
	ListGrievanceCategory() ([]*entity.GrievanceCategory, error)
	StoreGrievanceCategory(name string, description string, CodeName string) (int, error)
	UpdateGrievanceCategory(emp *entity.GrievanceCategory) (int, error)
	DeleteGrievanceCategory(id int) error
}
