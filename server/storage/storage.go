package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"errors"
)

var database *sql.DB

func DBConnect() *sql.DB {
	if database == nil {
		pwd, _ := os.Getwd()
		db := pwd + "/db.sqlite"
		log.Print("Connecting to DB ", db)
		datab, err := sql.Open("sqlite3", db)
		if err != nil {
			log.Fatal("Error opening database")
		}
		database = datab
	} else {
		log.Print("Ping the connection")
		database.Ping()
	}
	return database
}

func DBClose() {
	if database != nil {
		database.Close()
	}
}

func Insert(id, ciphertext string) error {
	db := DBConnect()
	_, err := db.Exec("INSERT INTO encrypted VALUES (?, ?)", id, ciphertext)
	if err != nil {
		log.Print(err.Error())
		return errors.New("Error inserting lift into DB")
	}
	return nil
}

func Select(id string) (string, error) {
	db := DBConnect()
	var cyphertext string
	err := db.QueryRow("SELECT ciphertext FROM encrypted WHERE id=?", id).Scan(&cyphertext)
	if err != nil {
		log.Print(err.Error())
		return "", errors.New("Unable to find message in db")
	}
	return cyphertext, nil
}