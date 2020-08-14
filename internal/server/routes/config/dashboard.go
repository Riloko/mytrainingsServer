package routesconfig

import (
	"mytrainingsserver/internal/server/auth"
	"mytrainingsserver/internal/server/routes"

	"github.com/gorilla/mux"
)

// DashboardRoutesConfig ...
type DashboardRoutesConfig struct {
	auth auth.Auth
	dr   routes.Dashboard
}

// ConfigureDashboardRoutes ...
func (d DashboardRoutesConfig) ConfigureDashboardRoutes(router *mux.Router) *mux.Router {
	dashboardAuth := d.auth.SecuredRoute
	dashboardAuth = d.auth.New(dashboardAuth)

	// Get Routes
	router.Handle("/dashboard", dashboardAuth.Handler(d.dr.GetDashboard())).Methods("GET")

	// Post Routes

	// Put Routes

	// Delete Routes

	return router
}
