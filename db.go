package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gopkg.in/gorp.v1"
)

// Initialize database connection
func initDb(dbHost string, dbSocket string, dbConnType int64, dbUser string, dbPass string,
	dbName string, production bool) *gorp.DbMap {
	connType := "@tcp(" + dbHost + ")"
	if dbConnType == 2 {
		connType = "@unix(" + dbSocket + ")"
	}

	db, err := sql.Open("mysql",
		dbUser+":"+dbPass+connType+"/"+dbName+"?parseTime=true")
	checkErr(err, "sql.Open failed")
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"MyISAM", "UTF8"}}

	if production {
		dbMap.TraceOff()
	} else {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "app:", log.Lmicroseconds))
	}

	// Set database timezone
	offset := hoursMinutes(timeZoneOffset(loc))
	dbMap.Exec(fmt.Sprintf("SET time_zone = '%s'", offset))
	result, _ := dbMap.SelectStr("SELECT now()")
	log.Println("Database server current time: " + result)

	return dbMap
}

func getExample(id int) ExampleTable {
	var data ExampleTable
	err := dbMap.SelectOne(&data, "SELECT * FROM ExampleTable WHERE id=?", id)
	checkErr(err, "[DB getExample]")
	return data
}

func getAllExample() []ExampleTable {
	var data []ExampleTable
	_, err := dbMap.Select(&data, "SELECT * FROM ExampleTable")
	checkErr(err, "[DB getAllExample]")
	return data
}

func insertExample(name string, description string) bool {
	query := "INSERT INTO ExampleTable (name,description) VALUES (?,?)"
	_, err := dbMap.Exec(query, name, description)
	checkErr(err, "[DB insertExample]")
	if err != nil {
		return false
	}

	return true
}

func updateExample(name string, id int) bool {
	query := `UPDATE ExampleTable SET name=? WHERE id=?`
	_, err := dbMap.Exec(query, name, id)
	checkErr(err, "[DB updateExample]")
	if err != nil {
		return false
	}

	return true
}

func deleteExample(name string) bool {
	query := `DELETE FROM ExampleTable WHERE name = ?`
	_, err := dbMap.Exec(query, name)
	checkErr(err, "[DB deleteExample]")
	if err != nil {
		return false
	}

	return true
}
