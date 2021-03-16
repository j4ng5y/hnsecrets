package server

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	R := mux.NewRouter()

	R.HandleFunc("/api/create_secret", createSecret).Methods("POST")
	R.HandleFunc("/api/remove_secret", removeSecret).Methods("DELETE")
	R.HandleFunc("/api/update_secret", updateSecret).Methods("PUT")
	R.HandleFunc("/api/list_secret", listSecret).Methods("GET")
	R.HandleFunc("/api/list_secrets", listSecrets).Methods("GET")

	return R
}
