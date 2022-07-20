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
	"github.com/tzdit/sample_api/services/usecase/grievance_resolution"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceResolution(c echo.Context) error {

	d := models.GrievanceResolution{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceResolution: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceResolution")
	}

	service := grievance_resolution.NewService()

	_, err := service.StoreGrievanceResolution(d.GrievanceId, d.GFUId, d.State, d.Description, d.Comment)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceResolution")

	} else {
		pp.Println("creating GrievanceResolution...")
		return c.JSON(http.StatusOK, "GrievanceResolution created")
	}

}


func ListGrievanceResolutions(c echo.Context) error {
	service := grievance_resolution.NewService()
	grievance_resolution, err := service.ListGrievanceResolution()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceResolution %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceResolution")
	}

	var json_grievance_resolution []*models.GrievanceResolution
	for _, GrievanceResolution := range grievance_resolution {
		u := &models.GrievanceResolution{
			Id:                GrievanceResolution.Id,
			GrievanceId:       GrievanceResolution.GrievanceId,
			GFUId:             GrievanceResolution.GFUId,
			State:       	   GrievanceResolution.State,
			Description: 	   GrievanceResolution.Description,
			Comment: 		   GrievanceResolution.Comment,
			UpdatedAt:         GrievanceResolution.UpdatedAt,
			CreatedAt:         GrievanceResolution.CreatedAt,
		}

		json_grievance_resolution = append(json_grievance_resolution, u)
	}
	return c.JSON(http.StatusOK, json_grievance_resolution)
}


func ListNewGrievanceResolutions(c echo.Context) error {
	service := grievance_resolution.NewService()
	grievance_resolution, err := service.ListNewGrievanceResolution()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceResolution %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceResolution")
	}

	var json_grievance_resolution []*models.GrievanceResolution
	for _, GrievanceResolution := range grievance_resolution {
		u := &models.GrievanceResolution{
			Id:                GrievanceResolution.Id,
			GrievanceId:       GrievanceResolution.GrievanceId,
			GFUId:             GrievanceResolution.GFUId,
			State:       	   GrievanceResolution.State,
			Description: 	   GrievanceResolution.Description,
			Comment: 		   GrievanceResolution.Comment,
			UpdatedAt:         GrievanceResolution.UpdatedAt,
			CreatedAt:         GrievanceResolution.CreatedAt,
		}

		json_grievance_resolution = append(json_grievance_resolution, u)
	}
	return c.JSON(http.StatusOK, json_grievance_resolution)
}


func ListApprovedGrievanceResolutions(c echo.Context) error {
	service := grievance_resolution.NewService()
	grievance_resolution, err := service.ListApprovedGrievanceResolution()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceResolution %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceResolution")
	}

	var json_grievance_resolution []*models.GrievanceResolution
	for _, GrievanceResolution := range grievance_resolution {
		u := &models.GrievanceResolution{
			Id:                GrievanceResolution.Id,
			GrievanceId:       GrievanceResolution.GrievanceId,
			GFUId:             GrievanceResolution.GFUId,
			State:       	   GrievanceResolution.State,
			Description: 	   GrievanceResolution.Description,
			Comment: 		   GrievanceResolution.Comment,
			UpdatedAt:         GrievanceResolution.UpdatedAt,
			CreatedAt:         GrievanceResolution.CreatedAt,
		}

		json_grievance_resolution = append(json_grievance_resolution, u)
	}
	return c.JSON(http.StatusOK, json_grievance_resolution)
}


func ListDeniedGrievanceResolutions(c echo.Context) error {
	service := grievance_resolution.NewService()
	grievance_resolution, err := service.ListDeniedGrievanceResolution()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceResolution %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceResolution")
	}

	var json_grievance_resolution []*models.GrievanceResolution
	for _, GrievanceResolution := range grievance_resolution {
		u := &models.GrievanceResolution{
			Id:                GrievanceResolution.Id,
			GrievanceId:       GrievanceResolution.GrievanceId,
			GFUId:             GrievanceResolution.GFUId,
			State:       	   GrievanceResolution.State,
			Description: 	   GrievanceResolution.Description,
			Comment: 		   GrievanceResolution.Comment,
			UpdatedAt:         GrievanceResolution.UpdatedAt,
			CreatedAt:         GrievanceResolution.CreatedAt,
		}

		json_grievance_resolution = append(json_grievance_resolution, u)
	}
	return c.JSON(http.StatusOK, json_grievance_resolution)
}



func ShowGrievanceResolution(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceResolution id")
	}

	service := grievance_resolution.NewService()
	GrievanceResolution, err := service.ShowGrievanceResolution(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceResolution with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceResolution{
			Id:                GrievanceResolution.Id,
			GrievanceId:       GrievanceResolution.GrievanceId,
			GFUId:             GrievanceResolution.GFUId,
			State:       	   GrievanceResolution.State,
			Description: 	   GrievanceResolution.Description,
			Comment: 		   GrievanceResolution.Comment,
			UpdatedAt:         GrievanceResolution.UpdatedAt,
			CreatedAt:         GrievanceResolution.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceResolution(c echo.Context) error {
	d := models.GrievanceResolution{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceResolution")
	}

	service := grievance_resolution.NewService()


	ent, err := service.ShowGrievanceResolution(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceResolution")
	}
	ent.Id = d.Id
	ent.GrievanceId = d.GrievanceId
	ent.GFUId = d.GFUId
	ent.State = d.State
	ent.Description = d.Description
	ent.Comment = d.Comment
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceResolution(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceResolution")

	} else {
		pp.Println("updating GrievanceResolution...")
		return c.JSON(http.StatusOK, "GrievanceResolution updated")
	}
}

//DeleteGrievantResolution delete a GrievantResolution
func DeleteGrievanceResolution(c echo.Context) error {
	d := models.GrievanceResolution{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_resolution.NewService()

	err := service.DeleteGrievanceResolution(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceResolution")

	} else {
		pp.Println("deleting GrievanceResolution...")
		return c.JSON(http.StatusOK, "GrievanceResolution deleted")
	}
}
