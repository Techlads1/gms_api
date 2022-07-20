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
	"github.com/tzdit/sample_api/services/usecase/grievance_resolution_state"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceResolutionState(c echo.Context) error {

	d := models.GrievanceResolutionState{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceResolutionState: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceResolutionState")
	}

	service := grievance_resolution_state.NewService()

	_, err := service.StoreGrievanceResolutionState(d.Name, d.Description)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceResolutionState")

	} else {
		pp.Println("creating GrievanceResolutionState...")
		return c.JSON(http.StatusOK, "GrievanceResolutionState created")
	}

}


func ListGrievanceResolutionStates(c echo.Context) error {
	service := grievance_resolution_state.NewService()
	grievance_resolution_states, err := service.ListGrievanceResolutionState()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceResolutionState %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceResolutionState")
	}

	var json_grievance_resolution_states []*models.GrievanceResolutionState
	for _, GrievanceResolutionState := range grievance_resolution_states {
		u := &models.GrievanceResolutionState{
			Id:                   	GrievanceResolutionState.Id,
			Name:       						GrievanceResolutionState.Name,
			Description: 						GrievanceResolutionState.Description,
			UpdatedAt:        			GrievanceResolutionState.UpdatedAt,
			CreatedAt:             	GrievanceResolutionState.CreatedAt,
		}

		json_grievance_resolution_states = append(json_grievance_resolution_states, u)
	}
	return c.JSON(http.StatusOK, json_grievance_resolution_states)
}


func ShowGrievanceResolutionState(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceResolutionState id")
	}

	service := grievance_resolution_state.NewService()
	GrievanceResolutionState, err := service.ShowGrievanceResolutionState(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceResolutionState with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceResolutionState{
			Id:                   	GrievanceResolutionState.Id,
			Name:       						GrievanceResolutionState.Name,
			Description: 						GrievanceResolutionState.Description,
			UpdatedAt:        			GrievanceResolutionState.UpdatedAt,
			CreatedAt:             	GrievanceResolutionState.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceResolutionState(c echo.Context) error {
	d := models.GrievanceResolutionState{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceResolutionState")
	}

	service := grievance_resolution_state.NewService()


	ent, err := service.ShowGrievanceResolutionState(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceResolutionState")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceResolutionState(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceResolutionState")

	} else {
		pp.Println("updating GrievanceResolutionState...")
		return c.JSON(http.StatusOK, "GrievanceResolutionState updated")
	}
}

//DeleteGrievantResolutionState delete a GrievantResolutionState
func DeleteGrievanceResolutionState(c echo.Context) error {
	d := models.GrievanceResolutionState{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_resolution_state.NewService()

	err := service.DeleteGrievanceResolutionState(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceResolutionState")

	} else {
		pp.Println("deleting GrievanceResolutionState...")
		return c.JSON(http.StatusOK, "GrievanceResolutionState deleted")
	}
}
