package cockpit

import (
	"github.com/gorilla/mux"
)

// AppHandler is a common struct that will have methods of all routes
type AppHandler struct{}

// Cockpit is a central operations model for the REST service
type Cockpit struct {
	AppRouter  *mux.Router
	AppRoutes  AppRoutes
	AppHandler AppHandler
}

func InitCockpit() *Cockpit {
	var cockpit Cockpit
	cockpit.AppRouter = mux.NewRouter()

	return &cockpit
}

// TODO: Cockpit Methods to be included
// GetResourcesByEntity
// GetEntity
// GetServiceName
// GetHandlerNameByResourceName
// so on..
