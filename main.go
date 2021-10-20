package main

import (
	"github.com/ArtisanCloud/PowerKeyVault/boostrap"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	logger "github.com/ArtisanCloud/PowerKeyVault/loggerManager"
	tester "github.com/ArtisanCloud/PowerKeyVault/tests"
)
import _ "github.com/ArtisanCloud/PowerKeyVault/config"
import "github.com/ArtisanCloud/PowerKeyVault/routes"

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	// init project
	boostrap.InitProject()

	// Initialize the Logger
	tester.TestFun()

	// Router the router as the default one provided by Gin
	Router = gin.Default()

	// Initialize the routes
	routes.InitializeRoutes(Router)

	// Start serving the application
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080®")
	err := Router.Run(config.AppConfigure.Server.Host + ":" + config.AppConfigure.Server.Port)
	if err != nil {
		logger.Error("router error:", map[string]interface{}{
			"err": err,
		})
	}

}
