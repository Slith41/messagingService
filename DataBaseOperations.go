package main

import (
	"database/sql"
	"fmt"
	"time"
)

//Dbinfo struct contains information about the database you are using
type Dbinfo struct {
	dbDriver   string
	dbUser     string
	dbPassword string
	dbName     string
}

func insertEmailIntoTable(db Dbinfo, table string, email string) {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db.dbUser, db.dbPassword, db.dbName)
	database, err := sql.Open(db.dbDriver, dataSourceName)
	checkErr(err)
	defer database.Close()

	err = database.QueryRow("INSERT INTO "+table+"(email) VALUES($1);", email).Scan()
	checkErr(err)
}

func insertEmailsIntoTable(db Dbinfo, table string, emails []string) {
	for _, email := range emails {
		insertEmailIntoTable(db, table, email)
	}
}

func selectAllFromTable(db Dbinfo, table string) map[string]time.Time {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db.dbUser, db.dbPassword, db.dbName)
	database, err := sql.Open(db.dbDriver, dataSourceName)
	checkErr(err)
	defer database.Close()

	rows, err := database.Query("SELECT * FROM " + table + ";")
	checkErr(err)
	defer rows.Close()

	emailsMap := make(map[string]time.Time)
	for rows.Next() {
		var email string
		var createdAt time.Time
		err = rows.Scan(&email, &createdAt)
		checkErr(err)

		emailsMap[email] = createdAt
	}
	return emailsMap
}

func deleteBasedOnEmail(db Dbinfo, table string, email string) {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db.dbUser, db.dbPassword, db.dbName)
	database, err := sql.Open(db.dbDriver, dataSourceName)
	checkErr(err)
	defer database.Close()

	stmt, err := database.Prepare("DELETE FROM emails_array where email=$1")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(email)
	checkErr(err)
}
func deleteBaseOnMultipleEmails(db Dbinfo, table string, emails []string) {
	for _, email := range emails {
		deleteBasedOnEmail(db, table, email)
	}
}
