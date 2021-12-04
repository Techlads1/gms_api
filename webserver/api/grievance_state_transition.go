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
	"github.com/tzdit/sample_api/services/usecase/grievance_state_Transition"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceStateTransition(c echo.Context) error {

	d := models.GrievanceStateTransition{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceStateTransition: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceStateTransition")
	}

	service := grievance_state_Transition.NewService()

	_, err := service.StoreGrievanceStateTransition(d.Description, d.FromStateId, d.ToStateId, d.Days)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceStateTransition")

	} else {
		pp.Println("creating GrievanceStateTransition...")
		return c.JSON(http.StatusOK, "GrievanceStateTransition created")
	}

}


func ListGrievanceStateTransitions(c echo.Context) error {
	service := grievance_state_Transition.NewService()
	grievance_state_Transitions, err := service.ListGrievanceStateTransition()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceStateTransition %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceStateTransition")
	}

	var json_grievance_state_Transitions []*models.GrievanceStateTransition
	for _, GrievanceStateTransition := range grievance_state_Transitions {
		u := &models.GrievanceStateTransition{
			Id:                   	GrievanceStateTransition.Id,
			Description:       			GrievanceStateTransition.Description,
			FromStateId: 						GrievanceStateTransition.FromStateId,
			ToStateId: 							GrievanceStateTransition.ToStateId,
			Days: 							    GrievanceStateTransition.Days,
			UpdatedAt:        			GrievanceStateTransition.UpdatedAt,
			CreatedAt:             	GrievanceStateTransition.CreatedAt,
		}

		json_grievance_state_Transitions = append(json_grievance_state_Transitions, u)
	}
	return c.JSON(http.StatusOK, json_grievance_state_Transitions)
}


func ShowGrievanceStateTransition(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceStateTransition id")
	}

	service := grievance_state_Transition.NewService()
	GrievanceStateTransition, err := service.ShowGrievanceStateTransition(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceStateTransition with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceStateTransition{
			Id:                   	GrievanceStateTransition.Id,
			Description:       			GrievanceStateTransition.Description,
			FromStateId: 						GrievanceStateTransition.FromStateId,
			ToStateId: 							GrievanceStateTransition.ToStateId,
			Days: 							    GrievanceStateTransition.Days,
			UpdatedAt:        			GrievanceStateTransition.UpdatedAt,
			CreatedAt:             	GrievanceStateTransition.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceStateTransition(c echo.Context) error {
	d := models.GrievanceStateTransition{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceStateTransition")
	}

	service := grievance_state_Transition.NewService()


	ent, err := service.ShowGrievanceStateTransition(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceStateTransition")
	}
	ent.Id = d.Id
	ent.Description = d.Description
	ent.FromStateId = d.FromStateId
	ent.ToStateId = d.ToStateId
	ent.Days = d.Days
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceStateTransition(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceStateTransition")

	} else {
		pp.Println("updating GrievanceStateTransition...")
		return c.JSON(http.StatusOK, "GrievanceStateTransition updated")
	}
}

//DeleteGrievantState delete a GrievantState
func DeleteGrievanceStateTransition(c echo.Context) error {
	d := models.GrievanceStateTransition{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_state_Transition.NewService()

	err := service.DeleteGrievanceStateTransition(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceStateTransition")

	} else {
		pp.Println("deleting GrievanceStateTransition...")
		return c.JSON(http.StatusOK, "GrievanceStateTransition deleted")
	}
}
