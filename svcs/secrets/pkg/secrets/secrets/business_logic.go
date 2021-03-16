package secrets

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var secrets []Secret

func createSecret() {
	var secret Secret
	secret.ID = strconv.Itoa(rand.Intn(10000000))
	secrets = append(secrets, secret)
}

func deleteSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range secrets {
		if item.ID == params["id"] {
			secrets = append(secrets[:index], secrets[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(secrets)
}
