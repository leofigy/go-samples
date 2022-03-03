package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbFile string
	flush  bool
)

const (
	dbType  = "sqlite3"
	dbUsers = `CREATE TABLE users (
		"idt" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"user" TEXT,
		"pass" TEXT
	  );`  // SQL Statement for Create Table

)

type User struct {
	ID   int
	User string
	Pass string
}

func init() {
	flag.StringVar(&dbFile, "db", "secret.db", "db file name")
	flag.BoolVar(&flush, "flush", false, "recreates the db")
}

func main() {
	flag.Parse()

	if strings.TrimSpace(dbFile) == "" {
		log.Fatalln("please provide a valid dbFile path")
	}

	if flush {
		if err := flushDB(dbFile); err != nil {
			log.Fatalln(err)
		}
	}

	db, err := sql.Open(dbType, dbFile)
	if err != nil {
		log.Fatalln(err)
	}

	if flush {
		if err := createTable(db); err != nil {
			log.Fatalln(err)
		}
	}

	defer db.Close()

	// dummy records don't use those pass
	inputUsers := []User{
		User{User: "bot", Pass: "pass"},
		User{User: "droid", Pass: "pass"},
		User{User: "trolls", Pass: "pass"},
	}

	for _, usr := range inputUsers {
		err = insertUser(db, &usr)
		if err != nil {
			log.Println(err)
		}
	}

	reads, err := readTable(db)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(reads)

}

func flushDB(path string) error {
	err := os.Remove(dbFile)
	if err != nil {
		return err
	}

	f, err := os.Create(dbFile)
	if err != nil {
		return err
	}

	return f.Close()
}

func createTable(db *sql.DB) error {

	log.Println("Create users")
	statement, err := db.Prepare(dbUsers) // Prepare SQL Statement
	if err != nil {
		return err
	}
	_, err = statement.Exec() // Execute SQL Statements
	return err
}

func readTable(db *sql.DB) ([]User, error) {
	row, err := db.Query("SELECT * FROM users ORDER by user")
	if err != nil {
		return nil, err
	}

	defer row.Close()

	users := []User{}

	for row.Next() {
		usr := User{}
		err = row.Scan(
			&(usr.ID),
			&(usr.User),
			&(usr.Pass),
		)

		if err != nil {
			log.Println("Error ", err)
			continue
		}

		users = append(users, usr)
	}

	return users, nil
}

func insertUser(db *sql.DB, user *User) error {
	insertStudentSQL := `INSERT INTO users(user, pass) VALUES (?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.User, user.Pass)
	return err
}
