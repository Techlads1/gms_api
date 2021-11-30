package routes

import (
	"github.com/tzdit/sample_api/webserver/api"

	"github.com/labstack/echo/v4"
)

//APIRouters Init Router
func APIRouters(app *echo.Echo) {

	//Protected api should be defined in this group
	//This api is only accessed by authenticated user
	//api := app.Group("/api/v1", middlewares.Jwt()) //remove the middleware if you want to make public /github.com/tzdit/sample_api/api/v1
	aim := app.Group("/api/v1/grm") //remove the middleware if you want to make public

	departments := aim.Group("/departments")
	{
		departments.POST("/create", api.CreateDepartment)
		departments.GET("/list", api.ListDepartments)
		departments.GET("/get/:id", api.GetDepartment)
		departments.POST("/update", api.UpdateDepartment)
		departments.POST("/delete", api.DeleteDepartment)
	}

	grievant_category := aim.Group("/grievant_categories")
	{
		grievant_category.POST("/store", api.StoreGrievantCategory)
		grievant_category.GET("/list", api.ListGrievantCategories)
		grievant_category.GET("/show/:id", api.ShowGrievantCategory)
		grievant_category.POST("/update", api.UpdateGrievantCategory)
		grievant_category.POST("/delete", api.DeleteGrievantCategory)
	}


}
