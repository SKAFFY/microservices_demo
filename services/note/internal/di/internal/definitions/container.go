package definitions

import (
	"github.com/SKAFFY/microservices_demo/pkg/gracefulServer"
	"github.com/SKAFFY/microservices_demo/pkg/routes"
	"github.com/gorilla/mux"
)

// Container is a root dependency injection container. It is required to describe
// your services.
type Container struct {
	// put the list of your services here
	// for example
	//  log *log.Logger

	// also, you can describe your services in a separate container
	// repositories RepositoryContainer
	apiContainer APIContainer
}

// this is a separate container
// type RepositoryContainer {
// 	entityRepository domain.EntityRepository
// }

type APIContainer struct {
	server *gracefulServer.Server

	publicRouter *mux.Router

	routes *[]routes.Routes
}
