package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/package/util"
	"github.com/tzdit/sample_api/services/usecase/grievance"
	"github.com/tzdit/sample_api/webserver/models"
)


func StoreGrievance(c echo.Context) error {

	d := models.Grievance{}
	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}
	pp.Printf("Grievance: %v\n", d)



	rand.Seed(time.Now().UnixNano())
	d.ReferenceNumber = string(fmt.Sprintf("%v",time.Now().Unix())) + "/" + time.Now().Month().String() +
	 "/" + fmt.Sprintf("%v", rand.Intn(999-111 +1) + 111) + "-" + fmt.Sprintf("%v", rand.Intn(999-222)+222)
	d.GrievantId = 1
	d.Comment = "Grievance Started."
	d.State = "Open"

		//Validate Data
		// if err := c.Validate(&d); err != nil {
		// 	pp.Print(err.)
		// 	return c.JSON(http.StatusInternalServerError, "could not validate Grievance")
		// }

	service := grievance.NewService()
	pp.Printf("Grievance: %v\n", d)


	no, err := service.StoreGrievance(d.Name, d.Description, d.ReferenceNumber, d.Comment, d.State, d.LocationOccurred, 
		d.FillingModeId, d.GrievanceSubCategoryId, d.GrievantId,
		 d.GrievantGroupId)

pp.Print(no)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error creating create Grievance")

	} else {
		pp.Println("creating Grievance...")
		return c.JSON(http.StatusOK, "Grievance created")
	}

}


func ListGrievances(c echo.Context) error {
	service := grievance.NewService()
	grievances, err := service.ListGrievance()
	if util.CheckError(err) {
		log.Errorf("error listing create Grievance %v", err)
		return c.JSON(http.StatusInternalServerError, "error listing create Grievance")
	}

	var json_grievances []*models.Grievance
	for _, Grievance := range grievances {
		u := &models.Grievance{
			Id:                   	Grievance.Id,
			Name:       						Grievance.Name,
			ReferenceNumber:        Grievance.ReferenceNumber,
			Comment:        				Grievance.Comment,
			State:        					Grievance.State,
			LocationOccurred:       Grievance.LocationOccurred,
			FillingModeId:         	Grievance.FillingModeId,
			GrievanceSubCategoryId:	Grievance.GrievanceSubCategoryId,
			GrievantId:             Grievance.GrievantId,
			GrievantGroupId:  			Grievance.GrievantGroupId,
			Description: 						Grievance.Description,
			UpdatedAt:        			Grievance.UpdatedAt,
			CreatedAt:             	Grievance.CreatedAt,
		}

		json_grievances = append(json_grievances, u)
	}
	return c.JSON(http.StatusOK, json_grievances)
}


func ShowGrievance(c echo.Context) error {
	cID, errParseInt := strconv.ParseInt(c.Param("id"), 10, 64)
	if util.CheckError(errParseInt) {

		return c.JSON(http.StatusInternalServerError, "invalid Grievance id")
	}

	service := grievance.NewService()
	Grievance, err := service.ShowGrievance(int(cID))
	if util.CheckError(err) {
		errMsg := fmt.Sprintf("could not get Grievance with id %v", cID)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	u := &models.Grievance{
			Id:                   	Grievance.Id,
			Name:       						Grievance.Name,
			ReferenceNumber:        Grievance.ReferenceNumber,
			Comment:        				Grievance.Comment,
			State:        					Grievance.State,
			LocationOccurred:       Grievance.LocationOccurred,
			FillingModeId:         	Grievance.FillingModeId,
			GrievanceSubCategoryId:	Grievance.GrievanceSubCategoryId,
			GrievantId:             Grievance.GrievantId,
			GrievantGroupId:  			Grievance.GrievantGroupId,
			Description: 						Grievance.Description,
			UpdatedAt:        			Grievance.UpdatedAt,
			CreatedAt:             	Grievance.CreatedAt,
		}
	
	return c.JSON(http.StatusOK, u)
}


func UpdateGrievance(c echo.Context) error {
	d := models.Grievance{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	//Validate Data
	if err := c.Validate(&d); err != nil {

		return c.JSON(http.StatusInternalServerError, "could not validate Grievance")
	}

	service := grievance.NewService()


	ent, err := service.ShowGrievance(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving Grievance")
	}
	ent.Id = d.Id
	ent.Name = d.Name
	ent.Description = d.Description
	ent.UpdatedAt = time.Now()

	_, err = service.UpdateGrievance(ent)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error updating Grievance")

	} else {
		pp.Println("updating Grievance...")
		return c.JSON(http.StatusOK, "Grievance updated")
	}
}

//DeleteGrievantState delete a GrievantState
func DeleteGrievance(c echo.Context) error {
	d := models.Grievance{}

	if err := c.Bind(&d); err != nil {
		log.Errorf("%s\n", err)
	}

	service := grievance.NewService()

	err := service.DeleteGrievance(d.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error deleting Grievance")

	} else {
		pp.Println("deleting Grievance...")
		return c.JSON(http.StatusOK, "Grievance deleted")
	}
}
