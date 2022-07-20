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
	"github.com/tzdit/sample_api/services/usecase/grievance_time_extension"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceTimeExtension(c echo.Context) error {

	d := models.GrievanceTimeExtension{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceTimeExtension: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceTimeExtension")
	}

	service := grievance_time_extension.NewService()

	_, err := service.StoreGrievanceTimeExtension(d.GrievanceId, d.GFUId, d.State, d.Description, d.Comment)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceTimeExtension")

	} else {
		pp.Println("creating GrievanceTimeExtension...")
		return c.JSON(http.StatusOK, "GrievanceTimeExtension created")
	}

}


func ListGrievanceTimeExtensions(c echo.Context) error {
	service := grievance_time_extension.NewService()
	grievance_time_extension, err := service.ListGrievanceTimeExtension()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceTimeExtension %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceTimeExtension")
	}

	var json_grievance_time_extension []*models.GrievanceTimeExtension
	for _, GrievanceTimeExtension := range grievance_time_extension {
		u := &models.GrievanceTimeExtension{
			Id:                GrievanceTimeExtension.Id,
			GrievanceId:       GrievanceTimeExtension.GrievanceId,
			GFUId:             GrievanceTimeExtension.GFUId,
			State:       	   GrievanceTimeExtension.State,
			Description: 	   GrievanceTimeExtension.Description,
			Comment: 		   GrievanceTimeExtension.Comment,
			UpdatedAt:         GrievanceTimeExtension.UpdatedAt,
			CreatedAt:         GrievanceTimeExtension.CreatedAt,
		}


		json_grievance_time_extension = append(json_grievance_time_extension, u)
	}
	return c.JSON(http.StatusOK, json_grievance_time_extension)
}


func ListNewGrievanceTimeExtensions(c echo.Context) error {
	service := grievance_time_extension.NewService()
	grievance_time_extension, err := service.ListNewGrievanceTimeExtension()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceTimeExtension %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceTimeExtension")
	}

	var json_grievance_time_extension []*models.GrievanceTimeExtension
	for _, GrievanceTimeExtension := range grievance_time_extension {
		u := &models.GrievanceTimeExtension{
			Id:                GrievanceTimeExtension.Id,
			GrievanceId:       GrievanceTimeExtension.GrievanceId,
			GFUId:             GrievanceTimeExtension.GFUId,
			State:       	   GrievanceTimeExtension.State,
			Description: 	   GrievanceTimeExtension.Description,
			Comment: 		   GrievanceTimeExtension.Comment,
			UpdatedAt:         GrievanceTimeExtension.UpdatedAt,
			CreatedAt:         GrievanceTimeExtension.CreatedAt,
		}


		json_grievance_time_extension = append(json_grievance_time_extension, u)
	}
	return c.JSON(http.StatusOK, json_grievance_time_extension)
}


func ListApprovedGrievanceTimeExtensions(c echo.Context) error {
	service := grievance_time_extension.NewService()
	grievance_time_extension, err := service.ListApprovedGrievanceTimeExtension()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceTimeExtension %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceTimeExtension")
	}

	var json_grievance_time_extension []*models.GrievanceTimeExtension
	for _, GrievanceTimeExtension := range grievance_time_extension {
		u := &models.GrievanceTimeExtension{
			Id:                GrievanceTimeExtension.Id,
			GrievanceId:       GrievanceTimeExtension.GrievanceId,
			GFUId:             GrievanceTimeExtension.GFUId,
			State:       	   GrievanceTimeExtension.State,
			Description: 	   GrievanceTimeExtension.Description,
			Comment: 		   GrievanceTimeExtension.Comment,
			UpdatedAt:         GrievanceTimeExtension.UpdatedAt,
			CreatedAt:         GrievanceTimeExtension.CreatedAt,
		}


		json_grievance_time_extension = append(json_grievance_time_extension, u)
	}
	return c.JSON(http.StatusOK, json_grievance_time_extension)
}


func ListDeniedGrievanceTimeExtensions(c echo.Context) error {
	service := grievance_time_extension.NewService()
	grievance_time_extension, err := service.ListDeniedGrievanceTimeExtension()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceTimeExtension %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceTimeExtension")
	}

	var json_grievance_time_extension []*models.GrievanceTimeExtension
	for _, GrievanceTimeExtension := range grievance_time_extension {
		u := &models.GrievanceTimeExtension{
			Id:                GrievanceTimeExtension.Id,
			GrievanceId:       GrievanceTimeExtension.GrievanceId,
			GFUId:             GrievanceTimeExtension.GFUId,
			State:       	   GrievanceTimeExtension.State,
			Description: 	   GrievanceTimeExtension.Description,
			Comment: 		   GrievanceTimeExtension.Comment,
			UpdatedAt:         GrievanceTimeExtension.UpdatedAt,
			CreatedAt:         GrievanceTimeExtension.CreatedAt,
		}


		json_grievance_time_extension = append(json_grievance_time_extension, u)
	}
	return c.JSON(http.StatusOK, json_grievance_time_extension)
}



func ShowGrievanceTimeExtension(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceTimeExtension id")
	}

	service := grievance_time_extension.NewService()
	GrievanceTimeExtension, err := service.ShowGrievanceTimeExtension(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceTimeExtension with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceTimeExtension{
			Id:                GrievanceTimeExtension.Id,
			GrievanceId:       GrievanceTimeExtension.GrievanceId,
			GFUId:             GrievanceTimeExtension.GFUId,
			State:       	   GrievanceTimeExtension.State,
			Description: 	   GrievanceTimeExtension.Description,
			Comment: 		   GrievanceTimeExtension.Comment,
			UpdatedAt:         GrievanceTimeExtension.UpdatedAt,
			CreatedAt:         GrievanceTimeExtension.CreatedAt,
		}

	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceTimeExtension(c echo.Context) error {
	d := models.GrievanceTimeExtension{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceTimeExtension")
	}

	service := grievance_time_extension.NewService()


	ent, err := service.ShowGrievanceTimeExtension(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceTimeExtension")
	}
	ent.Id = d.Id
	ent.GrievanceId = d.GrievanceId
	ent.GFUId = d.GFUId
	ent.State = d.State
	ent.Description = d.Description
	ent.Comment = d.Comment
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceTimeExtension(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceTimeExtension")

	} else {
		pp.Println("updating GrievanceTimeExtension...")
		return c.JSON(http.StatusOK, "GrievanceTimeExtension updated")
	}
}

//DeleteGrievantTimeExtension delete a GrievantTimeExtension
func DeleteGrievanceTimeExtension(c echo.Context) error {
	d := models.GrievanceTimeExtension{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_time_extension.NewService()

	err := service.DeleteGrievanceTimeExtension(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceTimeExtension")

	} else {
		pp.Println("deleting GrievanceTimeExtension...")
		return c.JSON(http.StatusOK, "GrievanceTimeExtension deleted")
	}
}
