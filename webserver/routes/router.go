package routes

import (
	"github.com/tzdit/sample_api/package/validator"
	"github.com/tzdit/sample_api/webserver/middlewares"

	"github.com/labstack/echo/v4"
)

// Routers function
func Routers(app *echo.Echo) {
	//Common middleware for all type of routers
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger(true))
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	//app.Use(middlewares.CSRF())
	app.Use(middlewares.Session())

	app.Validator = validator.GetValidator() //initialize custom validator

	//api routers
	APIRouters(app)

	//websocket api

	//web routes

}
