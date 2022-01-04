package grievance_sub_category

import (
	"github.com/tzdit/sample_api/services/entity"
)

//Reader interface
type Reader interface {
	Show(id int) (*entity.GrievanceSubCategory, error)
	List() ([]*entity.GrievanceSubCategory, error)
}

//DONE
//Writer interface
type Writer interface {
	Store(e *entity.GrievanceSubCategory) (int, error)
	Update(e *entity.GrievanceSubCategory) (int, error)
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
	ShowGrievanceSubCategory(id int) (*entity.GrievanceSubCategory, error)
	ListGrievanceSubCategory() ([]*entity.GrievanceSubCategory, error)
	StoreGrievanceSubCategory(name string, description string, CodeName string, GrievanceCategoryId int) (int, error)
	UpdateGrievanceSubCategory(emp *entity.GrievanceSubCategory) (int, error)
	DeleteGrievanceSubCategory(id int) error
}
