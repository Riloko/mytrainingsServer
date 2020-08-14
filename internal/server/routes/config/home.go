package routesconfig

import (
	"mytrainingsserver/internal/server/routes"

	"github.com/gorilla/mux"
)

// HomeRoutesConfig ...
type HomeRoutesConfig struct {
	r routes.Home
}

// ConfigureHomeRoutes ...
func (h HomeRoutesConfig) ConfigureHomeRoutes(router *mux.Router) *mux.Router {
	// Get routes
	router.HandleFunc("/", h.r.GetHome()).Methods("GET")
	router.HandleFunc("/login", h.r.GetLoginHome()).Methods("GET")
	router.HandleFunc("/registration", h.r.GetRegistrationHome()).Methods("GET")
	router.HandleFunc("/login/restore", h.r.GetRestorePassHome()).Methods("GET")

	// Post routes
	router.HandleFunc("/login", h.r.PostLoginHome()).Methods("POST")
	router.HandleFunc("/registration", h.r.PostRegistrationHome()).Methods("POST")
	router.HandleFunc("/login/restore", h.r.PostRestorePassHome()).Methods("POST")

	return router
}
