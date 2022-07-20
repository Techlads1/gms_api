package webserver

import (
	"github.com/tzdit/sample_api/webserver/routes"

	"github.com/labstack/echo/v4"
)

//StartWebserver starts a webserver
func StartWebserver() {
	// Echo instance
	e := echo.New()
	//Define renderer
	//e.Renderer = views.Renderer()

	//Disable echo banner
	e.HideBanner = true

	// Routes
	routes.Routers(e)

	//init cache
	//services.Init() //check if this solves the problem
	// Start server
	e.Logger.Fatal(e.Start(":4321")) //TODO: Replace this port with the one in the config file
}
