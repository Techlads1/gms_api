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
	"github.com/tzdit/sample_api/services/usecase/grievance_appeal_reason"
	"github.com/tzdit/sample_api/webserver/models"
)

func StoreGrievanceAppealReason(c echo.Context) error {

	d := models.GrievanceAppealReason{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceAppealReason: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceAppealReason")
	}

	service := grievance_appeal_reason.NewService()

	_, err := service.StoreGrievanceAppealReason(d.Name, d.Description)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceAppealReason")

	} else {
		pp.Println("creating GrievanceAppealReason...")
		return c.JSON(http.StatusOK, "GrievanceAppealReason created")
	}

}

func ListGrievanceAppealReasons(c echo.Context) error {
	service := grievance_appeal_reason.NewService()
	grievance_appeal_reasons, err := service.ListGrievanceAppealReason()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceAppealReason %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceAppealReason")
	}

	var json_grievance_appeal_reasons []*models.GrievanceAppealReason
	for _, GrievanceAppealReason := range grievance_appeal_reasons {
		u := &models.GrievanceAppealReason{
			Id:          GrievanceAppealReason.Id,
			Name:        GrievanceAppealReason.Name,
			Description: GrievanceAppealReason.Description,
			UpdatedAt:   GrievanceAppealReason.UpdatedAt,
			CreatedAt:   GrievanceAppealReason.CreatedAt,
		}

		json_grievance_appeal_reasons = append(json_grievance_appeal_reasons, u)
	}
	return c.JSON(http.StatusOK, json_grievance_appeal_reasons)
}

func ShowGrievanceAppealReason(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceAppealReason id")
	}

	service := grievance_appeal_reason.NewService()
	GrievanceAppealReason, err := service.ShowGrievanceAppealReason(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceAppealReason with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceAppealReason{
		Id:          GrievanceAppealReason.Id,
		Name:        GrievanceAppealReason.Name,
		Description: GrievanceAppealReason.Description,
		UpdatedAt:   GrievanceAppealReason.UpdatedAt,
		CreatedAt:   GrievanceAppealReason.CreatedAt,
	}

	return c.JSON(http.StatusOK, u)
}

func UpdateGrievanceAppealReason(c echo.Context) error {
	d := models.GrievanceAppealReason{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceAppealReason")
	}

	service := grievance_appeal_reason.NewService()

	ent, err := service.ShowGrievanceAppealReason(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceAppealReason")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceAppealReason(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceAppealReason")

	} else {
		pp.Println("updating GrievanceAppealReason...")
		return c.JSON(http.StatusOK, "GrievanceAppealReason updated")
	}
}

//DeleteGrievantAppealReason delete a GrievantAppealReason
func DeleteGrievanceAppealReason(c echo.Context) error {
	d := models.GrievanceAppealReason{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_appeal_reason.NewService()

	err := service.DeleteGrievanceAppealReason(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceAppealReason")

	} else {
		pp.Println("deleting GrievanceAppealReason...")
		return c.JSON(http.StatusOK, "GrievanceAppealReason deleted")
	}
}
