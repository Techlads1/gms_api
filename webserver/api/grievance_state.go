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
	"github.com/tzdit/sample_api/services/usecase/grievance_state"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceState(c echo.Context) error {

	d := models.GrievanceState{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceState: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceState")
	}

	service := grievance_state.NewService()

	_, err := service.StoreGrievanceState(d.Name, d.Description, d.CodeName ,d.Days)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceState")

	} else {
		pp.Println("creating GrievanceState...")
		return c.JSON(http.StatusOK, "GrievanceState created")
	}

}


func ListGrievanceStates(c echo.Context) error {
	service := grievance_state.NewService()
	grievance_states, err := service.ListGrievanceState()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceState %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceState")
	}

	var json_grievance_states []*models.GrievanceState
	for _, GrievanceState := range grievance_states {
		u := &models.GrievanceState{
			Id:                   	GrievanceState.Id,
			Name:       						GrievanceState.Name,
			Description: 						GrievanceState.Description,
			CodeName: 							GrievanceState.CodeName,
			Days: 							    GrievanceState.Days,
			UpdatedAt:        			GrievanceState.UpdatedAt,
			CreatedAt:             	GrievanceState.CreatedAt,
		}

		json_grievance_states = append(json_grievance_states, u)
	}
	return c.JSON(http.StatusOK, json_grievance_states)
}


func ShowGrievanceState(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceState id")
	}

	service := grievance_state.NewService()
	GrievanceState, err := service.ShowGrievanceState(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceState with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceState{
			Id:                   	GrievanceState.Id,
			Name:       						GrievanceState.Name,
			Description: 						GrievanceState.Description,
			CodeName: 							GrievanceState.CodeName,
			Days: 							    GrievanceState.Days,
			UpdatedAt:        			GrievanceState.UpdatedAt,
			CreatedAt:             	GrievanceState.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceState(c echo.Context) error {
	d := models.GrievanceState{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceState")
	}

	service := grievance_state.NewService()


	ent, err := service.ShowGrievanceState(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceState")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.CodeName = d.CodeName
	ent.Days = d.Days
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceState(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceState")

	} else {
		pp.Println("updating GrievanceState...")
		return c.JSON(http.StatusOK, "GrievanceState updated")
	}
}

//DeleteGrievantState delete a GrievantState
func DeleteGrievanceState(c echo.Context) error {
	d := models.GrievanceState{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_state.NewService()

	err := service.DeleteGrievanceState(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceState")

	} else {
		pp.Println("deleting GrievanceState...")
		return c.JSON(http.StatusOK, "GrievanceState deleted")
	}
}
