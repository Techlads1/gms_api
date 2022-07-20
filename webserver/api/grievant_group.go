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
	"github.com/tzdit/sample_api/services/usecase/grievant_group"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievantGroup(c echo.Context) error {

	d := models.GrievantGroup{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievantGroup: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievantGroup")
	}

	service := grievant_group.NewService()

	_, err := service.StoreGrievantGroup(d.Name, d.Description, d.GrievantCategoryId)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievantGroup")

	} else {
		pp.Println("creating GrievantGroup...")
		return c.JSON(http.StatusOK, "GrievantGroup created")
	}

}


func ListGrievantGroups(c echo.Context) error {
	service := grievant_group.NewService()
	grievant_group, err := service.ListGrievantGroup()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievantGroup %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievantGroup")
	}

	var json_grievant_group []*models.GrievantGroup
	for _, GrievantGroup := range grievant_group {
		u := &models.GrievantGroup{
			Id:                   	GrievantGroup.Id,
			Name:       						GrievantGroup.Name,
			Description: 						GrievantGroup.Description,
			GrievantCategoryId: 		GrievantGroup.GrievantCategoryId,
			UpdatedAt:        			GrievantGroup.UpdatedAt,
			CreatedAt:             	GrievantGroup.CreatedAt,
		}

		json_grievant_group = append(json_grievant_group, u)
	}
	return c.JSON(http.StatusOK, json_grievant_group)
}


func ShowGrievantGroup(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64) 
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievantGroup id")
	}

	service := grievant_group.NewService()
	GrievantGroup, err := service.ShowGrievantGroup(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievantGroup with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievantGroup{
			Id:                   	GrievantGroup.Id,
			Name:       						GrievantGroup.Name,
			Description: 						GrievantGroup.Description,
			GrievantCategoryId: 		GrievantGroup.GrievantCategoryId,
			UpdatedAt:        			GrievantGroup.UpdatedAt,
			CreatedAt:             	GrievantGroup.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievantGroup(c echo.Context) error {
	d := models.GrievantGroup{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievantGroup")
	}

	service := grievant_group.NewService()


	ent, err := service.ShowGrievantGroup(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievantGroup")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.GrievantCategoryId = d.GrievantCategoryId
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievantGroup(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievantGroup")

	} else {
		pp.Println("updating GrievantGroup...")
		return c.JSON(http.StatusOK, "GrievantGroup updated")
	}
}


func DeleteGrievantGroup(c echo.Context) error {
	d := models.GrievantGroup{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievant_group.NewService()

	err := service.DeleteGrievantGroup(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievantGroup")

	} else {
		pp.Println("deleting GrievantGroup...")
		return c.JSON(http.StatusOK, "GrievantGroup deleted")
	}
}
