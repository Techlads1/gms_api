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
	"github.com/tzdit/sample_api/services/usecase/grievance_filling_mode"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievanceFillingMode(c echo.Context) error {

	d := models.GrievanceFillingMode{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("GrievanceFillingMode: %v\n", d)

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceFillingMode")
	}

	service := grievance_filling_mode.NewService()

	_, err := service.StoreGrievanceFillingMode(d.Name, d.Description, d.CodeName)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create GrievanceFillingMode")

	} else {
		pp.Println("creating GrievanceFillingMode...")
		return c.JSON(http.StatusOK, "GrievanceFillingMode created")
	}

}


func ListGrievanceFillingModes(c echo.Context) error {
	service := grievance_filling_mode.NewService()
	grievance_filling_modes, err := service.ListGrievanceFillingMode()
	if util.CheckError(err) {
		log.Errorf("error listing create GrievanceFillingMode %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create GrievanceFillingMode")
	}

	var json_grievance_filling_modes []*models.GrievanceFillingMode
	for _, GrievanceFillingMode := range grievance_filling_modes {
		u := &models.GrievanceFillingMode{
			Id:                   	GrievanceFillingMode.Id,
			Name:       						GrievanceFillingMode.Name,
			Description: 						GrievanceFillingMode.Description,
			CodeName: 							GrievanceFillingMode.CodeName,
			UpdatedAt:        			GrievanceFillingMode.UpdatedAt,
			CreatedAt:             	GrievanceFillingMode.CreatedAt,
		}

		json_grievance_filling_modes = append(json_grievance_filling_modes, u)
	}
	return c.JSON(http.StatusOK, json_grievance_filling_modes)
}


func ShowGrievanceFillingMode(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid GrievanceFillingMode id")
	}

	service := grievance_filling_mode.NewService()
	GrievanceFillingMode, err := service.ShowGrievanceFillingMode(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get GrievanceFillingMode with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.GrievanceFillingMode{
			Id:                   	GrievanceFillingMode.Id,
			Name:       						GrievanceFillingMode.Name,
			Description: 						GrievanceFillingMode.Description,
			CodeName: 							GrievanceFillingMode.CodeName,
			UpdatedAt:        			GrievanceFillingMode.UpdatedAt,
			CreatedAt:             	GrievanceFillingMode.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievanceFillingMode(c echo.Context) error {
	d := models.GrievanceFillingMode{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate GrievanceFillingMode")
	}

	service := grievance_filling_mode.NewService()


	ent, err := service.ShowGrievanceFillingMode(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving GrievanceFillingMode")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.CodeName = d.CodeName
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievanceFillingMode(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating GrievanceFillingMode")

	} else {
		pp.Println("updating GrievanceFillingMode...")
		return c.JSON(http.StatusOK, "GrievanceFillingMode updated")
	}
}

//DeleteGrievantFillingMode delete a GrievantFillingMode
func DeleteGrievanceFillingMode(c echo.Context) error {
	d := models.GrievanceFillingMode{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance_filling_mode.NewService()

	err := service.DeleteGrievanceFillingMode(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting GrievanceFillingMode")

	} else {
		pp.Println("deleting GrievanceFillingMode...")
		return c.JSON(http.StatusOK, "GrievanceFillingMode deleted")
	}
}
