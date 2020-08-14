package router

import (
	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
}

// ListenRoutes ...
func (r Router) ListenRoutes() *mux.Router {
	routes := mux.NewRouter()
	return routes
}
