package api

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

var (
	user  User
	users []User
)

func ListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users = GetUsers(users)
	json.NewEncoder(w).Encode(users)
}

func UserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user := GetUser(user, params.ByName("id"))
	json.NewEncoder(w).Encode(user)
}
