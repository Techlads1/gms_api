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
	"github.com/tzdit/sample_api/services/usecase/grievance_sub_category"
	"github.com/tzdit/sample_api/webserver/models"
)

func StoreGrievanceSubCategory(c echo.Context) error {

	d := models.GrievanceSubCategory{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceSubCategory: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceSubCategory")
	}

	service := grievance_sub_category.NewService()

	_, err := service.StoreGrievanceSubCategory(d.Name, d.Description, d.CodeName, d.GrievanceCategoryId)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceSubCategory")

	} else {
		pp.Println("creating GrievanceSubCategory...")
		return c.JSON(http.StatusOK, "GrievanceSubCategory created")
	}

}

func ListGrievanceSubCategories(c echo.Context) error {
	service := grievance_sub_category.NewService()
	grievance_sub_categories, err := service.ListGrievanceSubCategory()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceSubCategory %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceSubCategory")
	}

	var json_grievance_sub_categories []*models.GrievanceSubCategory
	for _, GrievanceSubCategory := range grievance_sub_categories {
		u := &models.GrievanceSubCategory{
			Id:                  GrievanceSubCategory.Id,
			Name:                GrievanceSubCategory.Name,
			Description:         GrievanceSubCategory.Description,
			CodeName:            GrievanceSubCategory.CodeName,
			GrievanceCategoryId: GrievanceSubCategory.GrievanceCategoryId,
			UpdatedAt:           GrievanceSubCategory.UpdatedAt,
			CreatedAt:           GrievanceSubCategory.CreatedAt,
		}

		json_grievance_sub_categories = append(json_grievance_sub_categories, u)
	}
	return c.JSON(http.StatusOK, json_grievance_sub_categories)
}

func ShowGrievanceSubCategory(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceSubCategory id")
	}

	service := grievance_sub_category.NewService()
	GrievanceSubCategory, err := service.ShowGrievanceSubCategory(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceSubCategory with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceSubCategory{
		Id:                  GrievanceSubCategory.Id,
		Name:                GrievanceSubCategory.Name,
		Description:         GrievanceSubCategory.Description,
		CodeName:            GrievanceSubCategory.CodeName,
		GrievanceCategoryId: GrievanceSubCategory.GrievanceCategoryId,
		UpdatedAt:           GrievanceSubCategory.UpdatedAt,
		CreatedAt:           GrievanceSubCategory.CreatedAt,
	}

	return c.JSON(http.StatusOK, u)
}

func UpdateGrievanceSubCategory(c echo.Context) error {
	d := models.GrievanceSubCategory{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceSubCategory")
	}

	service := grievance_sub_category.NewService()

	ent, err := service.ShowGrievanceSubCategory(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceSubCategory")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.CodeName = d.CodeName
	ent.GrievanceCategoryId = d.GrievanceCategoryId
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceSubCategory(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceSubCategory")

	} else {
		pp.Println("updating GrievanceSubCategory...")
		return c.JSON(http.StatusOK, "GrievanceSubCategory updated")
	}
}

//DeleteGrievantCategory delete a GrievantCategory
func DeleteGrievanceSubCategory(c echo.Context) error {
	d := models.GrievanceSubCategory{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_sub_category.NewService()

	err := service.DeleteGrievanceSubCategory(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceSubCategory")

	} else {
		pp.Println("deleting GrievanceSubCategory...")
		return c.JSON(http.StatusOK, "GrievanceSubCategory deleted")
	}
}
