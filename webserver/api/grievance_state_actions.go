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
	"github.com/tzdit/sample_api/services/usecase/grievance_state_action"
	"github.com/tzdit/sample_api/webserver/models"
)

func StoreGrievanceStateAction(c echo.Context) error {

	d := models.GrievanceStateAction{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceStateAction: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceStateAction")
	}

	service := grievance_state_action.NewService()

	_, err := service.StoreGrievanceStateAction(d.Name, d.RolePerformAction, d.StateId)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceStateAction")

	} else {
		pp.Println("creating GrievanceStateAction...")
		return c.JSON(http.StatusOK, "GrievanceStateAction created")
	}

}

func ListGrievanceStateActions(c echo.Context) error {
	service := grievance_state_action.NewService()
	grievance_state_actions, err := service.ListGrievanceStateAction()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceStateAction %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceStateAction")
	}

	var json_grievance_state_actions []*models.GrievanceStateAction
	for _, GrievanceStateAction := range grievance_state_actions {
		u := &models.GrievanceStateAction{
			Id:                GrievanceStateAction.Id,
			Name:              GrievanceStateAction.Name,
			RolePerformAction: GrievanceStateAction.RolePerformAction,
			StateId:           GrievanceStateAction.StateId,
			UpdatedAt:         GrievanceStateAction.UpdatedAt,
			CreatedAt:         GrievanceStateAction.CreatedAt,
		}

		json_grievance_state_actions = append(json_grievance_state_actions, u)
	}
	return c.JSON(http.StatusOK, json_grievance_state_actions)
}

func ShowGrievanceStateAction(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceStateAction id")
	}

	service := grievance_state_action.NewService()
	GrievanceStateAction, err := service.ShowGrievanceStateAction(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceStateAction with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceStateAction{
		Id:                GrievanceStateAction.Id,
		Name:              GrievanceStateAction.Name,
		RolePerformAction: GrievanceStateAction.RolePerformAction,
		StateId:           GrievanceStateAction.StateId,
		UpdatedAt:         GrievanceStateAction.UpdatedAt,
		CreatedAt:         GrievanceStateAction.CreatedAt,
	}

	return c.JSON(http.StatusOK, u)
}

func UpdateGrievanceStateAction(c echo.Context) error {
	d := models.GrievanceStateAction{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceStateAction")
	}

	service := grievance_state_action.NewService()

	ent, err := service.ShowGrievanceStateAction(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceStateAction")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.RolePerformAction = d.RolePerformAction
	ent.StateId = d.StateId
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceStateAction(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceStateAction")

	} else {
		pp.Println("updating GrievanceStateAction...")
		return c.JSON(http.StatusOK, "GrievanceStateAction updated")
	}
}

//DeleteGrievantState delete a GrievantState
func DeleteGrievanceStateAction(c echo.Context) error {
	d := models.GrievanceStateAction{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_state_action.NewService()

	err := service.DeleteGrievanceStateAction(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceStateAction")

	} else {
		pp.Println("deleting GrievanceStateAction...")
		return c.JSON(http.StatusOK, "GrievanceStateAction deleted")
	}
}
