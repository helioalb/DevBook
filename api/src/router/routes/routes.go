package routes

import (
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
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
