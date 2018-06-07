package users

import (
	"encoding/json"
	"fmt"
	"github.com/callbinuvarghese/MYSQL/mysql"
	"net/http"
)

/*
curl -X POST -H 'Content-Type: application/x-www-form-urlencoded'
  -H "Origin: http://example.com" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: X-Requested-With" \
  -X OPTIONS --verbose \
 -d 'firstname=Ben&lastname=Varghese&city=Cumming&email=ben.varghese@accenture.com&age=42' "http://localhost:8080/users"
{"id":"d966178c-6745-11e8-90ac-6a0001865cd0"}

*/
// Post -- handles POST request to /users/new to create new user
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string

	user, errs := FormToUser(r)

	var created = false
	if len(errs) == 0 {
		fmt.Println("creating a new user")
		//tx, err := db.Begin()
		//checkErr(err)
		stmt, err := mysql.Db.Prepare("INSERT users( age, firstname, lastname, email, city) VALUES (?, ?, ?, ?, ?)")
		checkErr(err)
		defer stmt.Close() // danger!
		_, err = stmt.Exec(user.Age, user.FirstName, user.LastName, user.Email, user.City)
		if err != nil {
			errs = append(errs, err.Error())
			fmt.Println("Error creating a new user")
		} else {
			//tx, err := db.Commmit()
			checkErr(err)
			created = true
			fmt.Println("created a new user")
		}
	}

	if created {
		fmt.Println("FirstName", user.FirstName)
		json.NewEncoder(w).Encode(NewUserResponse{ID: user.FirstName})
	} else {
		fmt.Println("errors", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
