package mysql

/*
CREATE DATABASE IF NOT EXISTS user;
CREATE USER 'gouser'@'localhost' IDENTIFIED BY 'K..y';
GRANT ALL PRIVILEGES ON user.* TO 'gouser'@'localhost';
FLUSH PRIVILEGES;
USE user;
CREATE TABLE `users` (
        `uid` INT(10) NOT NULL AUTO_INCREMENT,
        `age` INT(10) NOT NULL,
        `firstname` VARCHAR(64) NULL DEFAULT NULL,
        `lastname` VARCHAR(64) NULL DEFAULT NULL,
        `city` VARCHAR(64) NULL DEFAULT NULL,
        `email` VARCHAR(64) NULL DEFAULT NULL,
        PRIMARY KEY (`uid`)
    );
*/
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
)

//contains the database connection
var Db *sql.DB // Note the sql package provides the namespace

func init() {
	// constants
	const (
		DBHost     = ""
		DBAuthUser = "gouser"
		DBAuthPass = "Accenture01!"
		DBSchema   = "user"
		DBPort     = "9042"
	)
	var err error
	// variables
	var (
		dbHost     string
		dbAuthUser string
		dbAuthPass string
		dbSchema   string
		dSN        string
	)

	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}

	dbHost = env("DB_HOST", DBHost)
	dbAuthUser = env("DB_USER", DBAuthUser)
	dbAuthPass = env("DB_PASSWORD", DBAuthPass)
	dbSchema = env("DB_SCHEMA", DBSchema)
	//dbPort, err := strconv.Atoi(env("DB_PORT", DBPort))
	if err != nil {
		panic(err)
	}

	dSN = fmt.Sprintf("%s:%s@%s/%s?timeout=30s&charset=utf8", dbAuthUser, dbAuthPass, dbHost, dbSchema)
	fmt.Printf("Connect to %s\n", dSN)

	Db, err = sql.Open("mysql", dSN)
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	yt := reflect.TypeOf(Db).Kind()
	fmt.Printf("%T: %s\n", yt, yt)
	fmt.Println("mysql init done")
}
