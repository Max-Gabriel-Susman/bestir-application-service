package handler

import (
	"net/http"

	"github.com/Max-Gabriel-Susman/bestir-application-service/internal/application"
	"github.com/Max-Gabriel-Susman/bestir-application-service/internal/foundation/database"
	"github.com/Max-Gabriel-Susman/bestir-application-service/internal/foundation/web"
)

var _ http.Handler = (*web.App)(nil)

// maybe we'll add gitsha and other params later
func API(d Deps) *web.App {
	app := web.NewApp()
	dbrConn := database.NewDBR(d.DB)
	applicationAPI := application.NewAPI(application.NewMySQLStore(dbrConn))
	ApplicationEndpoints(app, applicationAPI)
	return app
}
