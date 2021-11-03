package routes

import (
	"github.com/tzdit/sample_api/webserver/api"

	"github.com/labstack/echo/v4"
)

//APIRouters Init Router
func APIRouters(app *echo.Echo) {

	//Protected api should be defined in this group
	//This api is only accessed by authenticated user
	//api := app.Group("/api/v1", middlewares.Jwt()) //remove the middleware if you want to make public
	aim := app.Group("/github.com/tzdit/sample_api/api/v1") //remove the middleware if you want to make public

	departments := aim.Group("/departments")
	{
		departments.POST("/create", api.CreateDepartment)
		departments.GET("/list", api.ListDepartments)
		departments.GET("/get/:id", api.GetDepartment)
		departments.POST("/update", api.UpdateDepartment)
		departments.POST("/delete", api.DeleteDepartment)
	}

}
