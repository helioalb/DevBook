package router

import "github.com/gorilla/mux"

func Build() *mux.Router {
	return mux.NewRouter()
}
