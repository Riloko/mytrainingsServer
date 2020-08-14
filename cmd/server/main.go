package main

import (
	"fmt"
	"log"
	routesconfig "mytrainingsserver/internal/server/routes/config"
	"mytrainingsserver/internal/server/routes/router"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()

}

type config struct {
}

func main() {

	// Get router variable
	r := router.Router{}
	routes := r.ListenRoutes()

	// Get Routers structs
	homeRoutes := routesconfig.HomeRoutesConfig{}
	dashboardRoutes := routesconfig.DashboardRoutesConfig{}

	// Get routes from packages
	routes = homeRoutes.ConfigureHomeRoutes(routes)
	routes = dashboardRoutes.ConfigureDashboardRoutes(routes)

	loggedRouter := handlers.LoggingHandler(os.Stdout, routes)

	origins := []string{"http://localhost:8080", "http://localhost:3002"}

	fmt.Printf("Listening on http://localhost%v/ \n", os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), handlers.CORS(handlers.AllowedOrigins(origins))(loggedRouter)))
}
