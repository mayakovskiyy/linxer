package linxer_db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase(dbPath string) {
	var err error

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxOpenConns(1)

	createSQLTable := `CREATE TABLE IF NOT EXISTS links (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    link VARCHAR(300),
    shortlink VARCHAR(50),
    date DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createSQLTable)
	if err != nil {
		fmt.Println(err)
	}

}

func SaveData(link string) {
	_, err := db.Exec("INSERT INTO links (link, shortlink) VALUES (?)", link)
	if err != nil {
		fmt.Println(err)
	}
}
