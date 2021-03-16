package server

import (
	"log"
	"net/http"
)

// Router Handlers / Endpoints

R.HandleFunc("/api/secrets", getSecrets).Methods("GET")
R.HandleFunc("/api/secrets/{id}", getSecret).Methods("GET")
R.HandleFunc("/api/secrets", createSecret).Methods("POST")
R.HandleFunc("/api/secrets/{id}", updateSecret).Methods("PUT")
R.HandleFunc("/api/secrets/{id}", deleteSecret).Methods("DELETE")
log.Fatal(http.ListenAndServe(":8080", R))