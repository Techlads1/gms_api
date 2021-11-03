package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/package/util"
	"github.com/tzdit/sample_api/services/usecase/department"
	"github.com/tzdit/sample_api/webserver/models"
)

//CreateDepartment create a new department
func CreateDepartment(c echo.Context) error {

	d := models.Department{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("department: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate department")
	}

	service := department.NewService()

	//TODO make sure you capture these before saving procedure
	//at.CreatedBy=logged user
	createdBy := 17

	_, err := service.CreateDepartment(d.DepartmentTitle, d.DepartmentDescription, d.DepartmentSize, d.CampusId, createdBy)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create department")

	} else {
		pp.Println("creating department...")
		return c.JSON(http.StatusOK, "department created")
	}

}

//ListDepartments list all department
func ListDepartments(c echo.Context) error {
	service := department.NewService()
	departments, err := service.ListDepartment()
	if util.CheckError(err) {
		log.Errorf("error listing create department %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create department")
	}

	var jsonDepartments []*models.Department
	for _, department := range departments {
		u := &models.Department{
			Id:                    department.Id,
			DepartmentTitle:       department.DepartmentTitle,
			DepartmentDescription: department.DepartmentDescription,
			DepartmentSize:        department.DepartmentSize,
			CampusId:              department.CampusId,
			CreatedBy:             department.CreatedBy,
			CreatedAt:             department.CreatedAt,
		}

		jsonDepartments = append(jsonDepartments, u)
	}
	return c.JSON(http.StatusOK, jsonDepartments)
}

//GetDepartment get a single department
func GetDepartment(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64) //Converting String to unint64
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid department id")
	}

	service := department.NewService()
	department, err := service.GetDepartment(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get department with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.Department{
		Id:                    department.Id,
		DepartmentTitle:       department.DepartmentTitle,
		DepartmentDescription: department.DepartmentDescription,
		DepartmentSize:        department.DepartmentSize,
		CampusId:              department.CampusId,
		CreatedBy:             department.CreatedBy,
		CreatedAt:             department.CreatedAt,
	}
	return c.JSON(http.StatusOK, u)
}

//UpdateDepartment update a department
func UpdateDepartment(c echo.Context) error {
	d := models.Department{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "ould not validate department")
	}

	service := department.NewService()

	//TODO make sure you capture these before saving procedure
	//ac.updateBy=logged user

	ent, err := service.GetDepartment(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving department")
	}
	ent.Id = d.Id
	ent.DepartmentTitle = d.DepartmentTitle
	ent.DepartmentDescription = d.DepartmentDescription
	ent.DepartmentSize = d.DepartmentSize
	ent.CampusId = d.CampusId
	ent.UpdatedBy = 17

	_, err = service.UpdateDepartment(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating department")

	} else {
		pp.Println("updating department...")
		return c.JSON(http.StatusOK, "department updated")
	}
}

//DeleteDepartment delete a department
func DeleteDepartment(c echo.Context) error {
	d := models.Department{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := department.NewService()

	//TODO make sure you capture these before saving procedure
	//at.DeletedBy=logged user
	deletedBy := 17

	err := service.DeleteDepartment(d.Id, deletedBy)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting department")

	} else {
		pp.Println("deleting department...")
		return c.JSON(http.StatusOK, "department deleted")
	}
}
