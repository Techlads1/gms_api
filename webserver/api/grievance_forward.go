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
	"github.com/tzdit/sample_api/services/usecase/grievance_forward"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceForward(c echo.Context) error {

	d := models.GrievanceForward{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceForward: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceForward")
	}

	service := grievance_forward.NewService()

	_, err := service.StoreGrievanceForward(d.GrievanceId, d.State, d.FromGFUId, d.ToGFUId, d.Description, d.Comment)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceForward")

	} else {
		pp.Println("creating GrievanceForward...")
		return c.JSON(http.StatusOK, "GrievanceForward created")
	}

}


func ListGrievanceForwards(c echo.Context) error {
	service := grievance_forward.NewService()
	grievance_forward, err := service.ListGrievanceForward()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceForward %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceForward")
	}

	var json_grievance_forward []*models.GrievanceForward
	for _, GrievanceForward := range grievance_forward {
		u := &models.GrievanceForward{
			Id:                   	GrievanceForward.Id,
			GrievanceId:        GrievanceForward.GrievanceId,
			State:       			GrievanceForward.State,
			FromGFUId:       			GrievanceForward.FromGFUId,
			ToGFUId:       			GrievanceForward.ToGFUId,
			Description: 			GrievanceForward.Description,
			Comment: 				GrievanceForward.Comment,
			UpdatedAt:        		GrievanceForward.UpdatedAt,
			CreatedAt:             	GrievanceForward.CreatedAt,
		}

		json_grievance_forward = append(json_grievance_forward, u)
	}
	return c.JSON(http.StatusOK, json_grievance_forward)
}


func ListNewGrievanceForwards(c echo.Context) error {
	service := grievance_forward.NewService()
	grievance_forward, err := service.ListNewGrievanceForward()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceForward %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceForward")
	}

	var json_grievance_forward []*models.GrievanceForward
	for _, GrievanceForward := range grievance_forward {
		u := &models.GrievanceForward{
			Id:                   	GrievanceForward.Id,
			GrievanceId:        GrievanceForward.GrievanceId,
			State:       			GrievanceForward.State,
			FromGFUId:       			GrievanceForward.FromGFUId,
			ToGFUId:       			GrievanceForward.ToGFUId,
			Description: 			GrievanceForward.Description,
			Comment: 				GrievanceForward.Comment,
			UpdatedAt:        		GrievanceForward.UpdatedAt,
			CreatedAt:             	GrievanceForward.CreatedAt,
		}

		json_grievance_forward = append(json_grievance_forward, u)
	}
	return c.JSON(http.StatusOK, json_grievance_forward)
}


func ListApprovedGrievanceForwards(c echo.Context) error {
	service := grievance_forward.NewService()
	grievance_forward, err := service.ListApprovedGrievanceForward()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceForward %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceForward")
	}

	var json_grievance_forward []*models.GrievanceForward
	for _, GrievanceForward := range grievance_forward {
		u := &models.GrievanceForward{
			Id:                   	GrievanceForward.Id,
			GrievanceId:        GrievanceForward.GrievanceId,
			State:       			GrievanceForward.State,
			FromGFUId:       			GrievanceForward.FromGFUId,
			ToGFUId:       			GrievanceForward.ToGFUId,
			Description: 			GrievanceForward.Description,
			Comment: 				GrievanceForward.Comment,
			UpdatedAt:        		GrievanceForward.UpdatedAt,
			CreatedAt:             	GrievanceForward.CreatedAt,
		}

		json_grievance_forward = append(json_grievance_forward, u)
	}
	return c.JSON(http.StatusOK, json_grievance_forward)
}


func ListDeniedGrievanceForwards(c echo.Context) error {
	service := grievance_forward.NewService()
	grievance_forward, err := service.ListDeniedGrievanceForward()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceForward %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceForward")
	}

	var json_grievance_forward []*models.GrievanceForward
	for _, GrievanceForward := range grievance_forward {
		u := &models.GrievanceForward{
			Id:                   	GrievanceForward.Id,
			GrievanceId:        GrievanceForward.GrievanceId,
			State:       			GrievanceForward.State,
			FromGFUId:       			GrievanceForward.FromGFUId,
			ToGFUId:       			GrievanceForward.ToGFUId,
			Description: 			GrievanceForward.Description,
			Comment: 				GrievanceForward.Comment,
			UpdatedAt:        		GrievanceForward.UpdatedAt,
			CreatedAt:             	GrievanceForward.CreatedAt,
		}

		json_grievance_forward = append(json_grievance_forward, u)
	}
	return c.JSON(http.StatusOK, json_grievance_forward)
}



func ShowGrievanceForward(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceForward id")
	}

	service := grievance_forward.NewService()
	GrievanceForward, err := service.ShowGrievanceForward(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceForward with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceForward{
			Id:                   	GrievanceForward.Id,
			GrievanceId:        GrievanceForward.GrievanceId,
			State:       			GrievanceForward.State,
			FromGFUId:       			GrievanceForward.FromGFUId,
			ToGFUId:       			GrievanceForward.ToGFUId,
			Description: 			GrievanceForward.Description,
			Comment: 				GrievanceForward.Comment,
			UpdatedAt:        		GrievanceForward.UpdatedAt,
			CreatedAt:             	GrievanceForward.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceForward(c echo.Context) error {
	d := models.GrievanceForward{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceForward")
	}

	service := grievance_forward.NewService()


	ent, err := service.ShowGrievanceForward(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceForward")
	}
	ent.Id = d.Id
	ent.GrievanceId = d.GrievanceId
	ent.State = d.State
	ent.FromGFUId = d.FromGFUId
	ent.ToGFUId = d.ToGFUId
	ent.Description = d.Description
	ent.Comment = d.Comment
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceForward(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceForward")

	} else {
		pp.Println("updating GrievanceForward...")
		return c.JSON(http.StatusOK, "GrievanceForward updated")
	}
}

//DeleteGrievantForward delete a GrievantForward
func DeleteGrievanceForward(c echo.Context) error {
	d := models.GrievanceForward{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_forward.NewService()

	err := service.DeleteGrievanceForward(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceForward")

	} else {
		pp.Println("deleting GrievanceForward...")
		return c.JSON(http.StatusOK, "GrievanceForward deleted")
	}
}
