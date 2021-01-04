package main

import (
	_ "github.com/lib/pq"
)

func main() {

	var db Dbinfo
	db.dbDriver = "postgres"
	db.dbUser = "postgres"
	db.dbPassword = "liac1912"
	db.dbName = "emails"

	insertEmailIntoTable(db, "emails", "nameisconfidentialinformation@gmail.com")
	selectAllFromTable(db, "emails")
	setupRouts()

}
