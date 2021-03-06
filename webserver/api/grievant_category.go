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
	"github.com/tzdit/sample_api/services/usecase/grievant_category"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievantCategory(c echo.Context) error {

	d := models.GrievantCategory{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievantCategory: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievantCategory")
	}

	service := grievant_category.NewService()

	_, err := service.StoreGrievantCategory(d.Name, d.Description)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievantCategory")

	} else {
		pp.Println("creating GrievantCategory...")
		return c.JSON(http.StatusOK, "GrievantCategory created")
	}

}


func ListGrievantCategories(c echo.Context) error {
	service := grievant_category.NewService()
	grievant_categories, err := service.ListGrievantCategory()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievantCategory %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievantCategory")
	}

	var json_grievant_categories []*models.GrievantCategory
	for _, GrievantCategory := range grievant_categories {
		u := &models.GrievantCategory{
			Id:                   	GrievantCategory.Id,
			Name:       						GrievantCategory.Name,
			Description: 						GrievantCategory.Description,
			UpdatedAt:        			GrievantCategory.UpdatedAt,
			CreatedAt:             	GrievantCategory.CreatedAt,
		}

		json_grievant_categories = append(json_grievant_categories, u)
	}
	return c.JSON(http.StatusOK, json_grievant_categories)
}


func ShowGrievantCategory(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64) 
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievantCategory id")
	}

	service := grievant_category.NewService()
	GrievantCategory, err := service.ShowGrievantCategory(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievantCategory with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievantCategory{
			Id:                   	GrievantCategory.Id,
			Name:       						GrievantCategory.Name,
			Description: 						GrievantCategory.Description,
			UpdatedAt:        			GrievantCategory.UpdatedAt,
			CreatedAt:             	GrievantCategory.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievantCategory(c echo.Context) error {
	d := models.GrievantCategory{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievantCategory")
	}

	service := grievant_category.NewService()


	ent, err := service.ShowGrievantCategory(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievantCategory")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievantCategory(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievantCategory")

	} else {
		pp.Println("updating GrievantCategory...")
		return c.JSON(http.StatusOK, "GrievantCategory updated")
	}
}


func DeleteGrievantCategory(c echo.Context) error {
	d := models.GrievantCategory{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievant_category.NewService()

	err := service.DeleteGrievantCategory(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievantCategory")

	} else {
		pp.Println("deleting GrievantCategory...")
		return c.JSON(http.StatusOK, "GrievantCategory deleted")
	}
}
