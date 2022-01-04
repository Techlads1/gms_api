package grievant_category

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievantCategory, error)
	List() ([]*entity.GrievantCategory, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievantCategory) (int, error)
	Update(e *entity.GrievantCategory) (int, error)
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
	ShowGrievantCategory(id int) (*entity.GrievantCategory, error)
	ListGrievantCategory() ([]*entity.GrievantCategory, error)
	StoreGrievantCategory(name string, description string) (int, error)
	UpdateGrievantCategory(emp *entity.GrievantCategory) (int, error)
	DeleteGrievantCategory(id int) error
}
