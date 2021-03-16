package secrets

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Secret struct {
	ID     string `json: "id"`
	Secret string `json: "secret"`
}

var secrets []Secret

func createSecret() {
	var secret Secret
	_ = json.NewDecoder(r.Body).Decode(&secret)
	secret.ID = strconv.Itoa(rand.Intn(10000000))
	secrets = append(secrets, secret)
	json.NewEncoder(w).Encode(secret)
}

func deleteSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range secrets {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
