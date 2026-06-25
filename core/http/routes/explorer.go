package routes

import (
	"github.com/labstack/echo/v4"
	coreExplorer "github.com/siperal/hYZmet-localai/core/explorer"
	"github.com/siperal/hYZmet-localai/core/http/endpoints/explorer"
)

func RegisterExplorerRoutes(app *echo.Echo, db *coreExplorer.Database) {
	app.GET("/", explorer.Dashboard())
	app.POST("/network/add", explorer.AddNetwork(db))
	app.GET("/networks", explorer.ShowNetworks(db))
}
