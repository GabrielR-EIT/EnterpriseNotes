package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = ""
)

func CreateDB() string {
	var returnMsg string

	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	// Ping the database for connectivity
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := `DROP DATABASE IF EXISTS EnterpriseNotes;`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when creating the 'EnterpriseNotes' database.\n"
		return returnMsg
	}
	sqlQuery = `CREATE DATABASE EnterpriseNotes;`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when creating the 'EnterpriseNotes' database.\n"
		return returnMsg
	}
	//const dbname = "EnterpriseNotes"
	returnMsg += "The 'EnterpriseNotes' database was created successfully.\n"

	return returnMsg
}

func CreateTables() string {
	var returnMsg string

	//Connect to the database
	const dbname = "enterprisenotes"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Ping the database for connectivity
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create the users table
	sqlQuery := `DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		userID INT PRIMARY KEY, 
		userName VARCHAR(100), 
		userReadSetting BOOL, 
		userWriteSetting BOOL
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when creating the 'users' table.\n"
		return returnMsg
	}
	returnMsg += "The 'users' table was created successfully.\n"

	// Create the notes table
	sqlQuery = `DROP TABLE IF EXISTS notes;
	CREATE TABLE notes (
		noteID INT PRIMARY KEY, 
		noteName VARCHAR(100), 
		noteText TEXT, 
		noteCompletionTime timestamp,
		noteStatus VARCHAR(20),
		noteDelegation VARCHAR(20),
		noteSharedUsers VARCHAR(100)
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when creating the 'notes' table.\n"
		return returnMsg
	}
	returnMsg += "The 'notes' table was created successfully.\n"

	// Create the associations table
	sqlQuery = `DROP TABLE IF EXISTS associations;
	CREATE TABLE associations (
		associationID INT PRIMARY KEY, 
		userID INT, 
		noteID INT, 
		associationPerm VARCHAR(20),
		CONSTRAINT fk_user FOREIGN KEY(userID) REFERENCES users(userID) ON DELETE CASCADE,
		CONSTRAINT fk_note FOREIGN KEY(noteID) REFERENCES notes(noteID) ON DELETE CASCADE
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when creating the 'associations' table.\n"
		return returnMsg
	}
	returnMsg += "The 'associations' table was created successfully.\n"

	return returnMsg
}
