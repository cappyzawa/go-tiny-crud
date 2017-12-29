package api

import (
	"database/sql"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   string	`json: "id"`
	Name string	`json: "name"`
	Ruby string `json: "ruby"`
}

func ConnectMysql() *sql.DB {
	mysqlDb, err := sql.Open("mysql", os.Getenv("MYSQL_USER_NAME") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(127.0.0.1:3306)/tiny_crud")
	if err != nil {
		log.Fatal(err)
	}
	return mysqlDb
}

func GetUser(queryId string) User{
	var user User
	mysql := ConnectMysql()
	defer mysql.Close()

	notFoundErr := mysql.QueryRow(`SELECT * FROM users WHERE id = ?`, queryId).Scan(&user.Id, &user.Name, &user.Ruby)
	if notFoundErr != nil {
		log.Fatal("not found user by " + queryId)
	}

	return user
}

func GetUsers(users []User) []User{
	mysql := ConnectMysql()
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
