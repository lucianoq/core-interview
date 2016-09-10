package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"errors"
)

var database *sql.DB
var dbfile string

func Check() error {
	pwd, _ := os.Getwd()
	dbfile = pwd + "/db.sqlite"
	_, err := os.Stat(dbfile)
	if os.IsNotExist(err) {
		return errors.New("DB file does not exist")
	}
	return nil
}

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
		return err
	}
	return nil
}

func Select(id string) (string, error) {
	db := DBConnect()
	var cyphertext string
	err := db.QueryRow("SELECT ciphertext FROM encrypted WHERE id=?", id).Scan(&cyphertext)
	if err != nil {
		return "", err
	}
	return cyphertext, nil
}