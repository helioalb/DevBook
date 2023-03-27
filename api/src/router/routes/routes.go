package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Handler               func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Handler)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	return r
}
