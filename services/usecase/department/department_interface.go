package department

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Get(id int) (*entity.Department, error)
	List() ([]*entity.Department, error)
}

//DONE
//Writer interface
type Writer interface {
	Create(e *entity.Department) (int, error)
	Update(e *entity.Department) (int, error)
	Delete(id, deletedBy int) error
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
	GetDepartment(id int) (*entity.Department, error)
	ListDepartment() ([]*entity.Department, error)
	CreateDepartment(name string, code string, createdBy int) (int, error)
	UpdateDepartment(emp *entity.Department) (int, error)
	DeleteDepartment(id, deletebBy int) error
}
