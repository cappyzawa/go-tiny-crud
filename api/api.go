package api

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/cappyzawa/go-tiny-crud/db"
)

type User struct {
	Id   string	`json: "id"`
	Name string	`json: "name"`
	Ruby string `json: "ruby"`
}

var (
	user User
	users []User
)

func ListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users = GetUsers(users)
	json.NewEncoder(w).Encode(users)
}

func UserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user := GetUser(params.ByName("id"))
	json.NewEncoder(w).Encode(user)
}

func GetUser(queryId string) User{
	mysql := db.ConnectMysql()
	defer mysql.Close()

	notFoundErr := mysql.QueryRow(`SELECT * FROM users WHERE id = ?`, queryId).Scan(&user.Id, &user.Name, &user.Ruby)
	if notFoundErr != nil {
		log.Fatal("not found user by " + queryId)
	}

	return user
}

func GetUsers(users []User) []User{
	mysql := db.ConnectMysql()
	defer mysql.Close()

	rows, err := mysql.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Ruby)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}
