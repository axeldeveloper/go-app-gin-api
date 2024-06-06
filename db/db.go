package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db := db.Connection()
func Connection() *sql.DB {

	/*DATABASE_HOST=ep-blue-wave-a2q6c5nw.eu-central-1.pg.koyeb.app
	DATABASE_USER=koyeb-adm
	DATABASE_PASSWORD=K3w9UAoghCxd
	DATABASE_NAME=koyebdb*/

	const (
		server_dev   = "ep-blue-wave-a2q6c5nw.eu-central-1.pg.koyeb.app"
		user_dev     = "koyeb-adm"
		password_dev = "K3w9UAoghCxd"
		database_dev = "koyebdb"
	)

	// connStr := "user='koyeb-adm' password=K3w9UAoghCxd host=ep-blue-wave-a2q6c5nw.eu-central-1.pg.koyeb.app dbname='koyebdb'"
	conexao := fmt.Sprintf("user=%s password=%s host=%s dbname=%s",
		user_dev,
		password_dev,
		server_dev,
		database_dev,
	)

	db, err := sql.Open("postgres", conexao)
	// db, err := sql.Open("postgres", "user=postgres dbname=impulso password=postgres sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Printf("VM Connected!\n")
	return db

}
