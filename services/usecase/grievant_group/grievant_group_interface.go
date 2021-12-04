package grievant_group

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievantGroup, error)
	List() ([]*entity.GrievantGroup, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievantGroup) (int, error)
	Update(e *entity.GrievantGroup) (int, error)
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
	ShowGrievantGroup(id int) (*entity.GrievantGroup, error)
	ListGrievantGroup() ([]*entity.GrievantGroup, error)
	StoreGrievantGroup(name string, description string, GrievantCategoryId int) (int, error)
	UpdateGrievantGroup(emp *entity.GrievantGroup) (int, error)
	DeleteGrievantGroup(id int) error
}
