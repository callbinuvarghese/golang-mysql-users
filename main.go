package main

/*
https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/

*/

import (
	"encoding/json"
	"fmt"
	"github.com/callbinuvarghese/MYSQL/mysql"
	"github.com/callbinuvarghese/MYSQL/users"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

/*
  Mysql
  curl -X POST -d 'firstname=User&lastname=Two&city=London&email=user.two@accenture.com&age=32' "http://localhost:8080/users"
  
*/
const AppPort = ":8080"

func main() {
	db := mysql.Db
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)

	router.HandleFunc("/users", users.Get).Methods("GET")
	router.HandleFunc("/users", users.Post).Methods("POST")
	router.HandleFunc("/users/{user_uuid}", users.GetOne)

	fmt.Println("Server listening" + AppPort)
	//log.Fatal(http.ListenAndServe(AppPort, router))
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	//for getting user agent in the log
	//loggedRouter := CombinedLoggingHandler(os.Stdout, router)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	//log.Fatal(http.ListenAndServe(AppPort,
	//	handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(loggedRouter)))

	log.Fatal(http.ListenAndServe(AppPort,
		handlers.CORS(headersOk, originsOk, methodsOk)(loggedRouter)))

}

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
