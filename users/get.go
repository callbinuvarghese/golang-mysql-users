package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/callbinuvarghese/MYSQL/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//function to check error
func checkErr(err error) {
	fmt.Println("checkerr:1")
	if err != nil {
		panic(err)
	}
}

//curl -L "http://localhost:8080/users/"
// Get -- handles GET request to /users/ to fetch all users
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params (unused here)
func Get(w http.ResponseWriter, r *http.Request) {
	var count int
	var rowcollection []User

	fmt.Println("list-users")

	rows, err := mysql.Db.Query("SELECT uid, age, firstname, lastname, email, city FROM users")
	checkErr(err)
	fmt.Println("list-users:2")

	defer rows.Close()
	for rows.Next() {
		fmt.Println("list-users:3")
		var r User
		err = rows.Scan(&r.ID,
			&r.Age,
			&r.FirstName,
			&r.LastName,
			&r.Email,
			&r.City)
		checkErr(err)
		rowcollection = append(rowcollection, r)
		count++
	}
	fmt.Println("list-users: count ", count)
	//json.NewEncoder(w).Encode(AllUsersResponse{Users: rowcollection})
	json.NewEncoder(w).Encode(rowcollection)
}

//curl -L "http://localhost:8080/users/d966178c-6745-11e8-90ac-6a0001865cd0"
// GetOne -- handles GET request to /users/{user_uuid} to fetch one user
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func GetOne(w http.ResponseWriter, r *http.Request) {
	var user User
	var errs []string
	var err error
	var found = false

	fmt.Println("get-user")
	vars := mux.Vars(r)
	user.ID, err = strconv.Atoi(vars["user_uuid"])
	if err != nil {
		fmt.Println("get-user: Could not get id")
		errs = append(errs, err.Error())
	} else {
		fmt.Println("get-user: userid", user.ID)
		stmt, err := mysql.Db.Prepare("SELECT age,firstname,lastname,city,email FROM users WHERE uid = ?")
		checkErr(err)
		fmt.Println("get-user: prepared SQL")
		err = stmt.QueryRow(user.ID).Scan(&user.Age, &user.FirstName, &user.LastName, &user.City, &user.Email)
		if err != nil {
			fmt.Println("get-user: error in SQL")
			if err == sql.ErrNoRows {
				// there were no rows, but otherwise no error occurred
				fmt.Println("get-user: No user found for userid:", user.ID)
				found = false
			} else {
				fmt.Println("get-user: fatal error in SQL")
				log.Fatal(err)
			}
		} else {
			fmt.Println("get-user: found user")
			found = true
		}
		if !found {
			errs = append(errs, "User not found")
		}
	}

	if found {
		json.NewEncoder(w).Encode(GetUserResponse{User: user})
	} else {
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
