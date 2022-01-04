package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/package/util"
	"github.com/tzdit/sample_api/services/usecase/grievance_category"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceCategory(c echo.Context) error {

	d := models.GrievanceCategory{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceCategory: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceCategory")
	}

	service := grievance_category.NewService()

	_, err := service.StoreGrievanceCategory(d.Name, d.Description, d.CodeName)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceCategory")

	} else {
		pp.Println("creating GrievanceCategory...")
		return c.JSON(http.StatusOK, "GrievanceCategory created")
	}

}


func ListGrievanceCategories(c echo.Context) error {
	service := grievance_category.NewService()
	grievance_categories, err := service.ListGrievanceCategory()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceCategory %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceCategory")
	}

	var json_grievance_categories []*models.GrievanceCategory
	for _, GrievanceCategory := range grievance_categories {
		u := &models.GrievanceCategory{
			Id:                   	GrievanceCategory.Id,
			Name:       						GrievanceCategory.Name,
			Description: 						GrievanceCategory.Description,
			CodeName: 							GrievanceCategory.CodeName,
			UpdatedAt:        			GrievanceCategory.UpdatedAt,
			CreatedAt:             	GrievanceCategory.CreatedAt,
		}

		json_grievance_categories = append(json_grievance_categories, u)
	}
	return c.JSON(http.StatusOK, json_grievance_categories)
}


func ShowGrievanceCategory(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceCategory id")
	}

	service := grievance_category.NewService()
	GrievanceCategory, err := service.ShowGrievanceCategory(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceCategory with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceCategory{
			Id:                   	GrievanceCategory.Id,
			Name:       						GrievanceCategory.Name,
			Description: 						GrievanceCategory.Description,
			CodeName: 							GrievanceCategory.CodeName,
			UpdatedAt:        			GrievanceCategory.UpdatedAt,
			CreatedAt:             	GrievanceCategory.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceCategory(c echo.Context) error {
	d := models.GrievanceCategory{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceCategory")
	}

	service := grievance_category.NewService()


	ent, err := service.ShowGrievanceCategory(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceCategory")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.CodeName = d.CodeName
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceCategory(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceCategory")

	} else {
		pp.Println("updating GrievanceCategory...")
		return c.JSON(http.StatusOK, "GrievanceCategory updated")
	}
}

//DeleteGrievantCategory delete a GrievantCategory
func DeleteGrievanceCategory(c echo.Context) error {
	d := models.GrievanceCategory{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_category.NewService()

	err := service.DeleteGrievanceCategory(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceCategory")

	} else {
		pp.Println("deleting GrievanceCategory...")
		return c.JSON(http.StatusOK, "GrievanceCategory deleted")
	}
}
