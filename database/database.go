package database

import (
	"database/sql"
	"fmt"
	"log"

	"crypto/sha256"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Open(databasePath string) {
	var err error
	db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
}

// test username and password
// LegalizeCarrots - Password123
func CheckPassword(username string, password string) bool {
	rows, _ := db.Query(fmt.Sprintf("SELECT passwordHash FROM passwordHashes WHERE username = '%s'", username))

	if rows.Next() {
		var hash string
		rows.Scan(&hash)
		return fmt.Sprintf("%x", sha256.Sum256([]byte(password))) == hash
	}
	fmt.Println("nothing found")
	return false
}
