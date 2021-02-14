package server

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	R := mux.NewRouter()

	return R
}
