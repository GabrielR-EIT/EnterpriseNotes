package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "enterprisenotes"
)

// Create Struct for Test Data
type Data struct {
	Users        []User        `json:"users"`
	Notes        []Note        `json:"notes"`
	Associations []Association `json:"associations"`
}

// Connect to Database Function
// func ConnectToDB() string {
// 	var returnMsg string

// 	// Connect to the database
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

// 	// Ping the database for connectivity
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//defer db.Close()
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	returnMsg += "A successful PostgreSQL connection was made.\n"
// 	return returnMsg
// }

// Create Database Function
func CreateDB() string {
	var returnMsg string

	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	// Ping the database for connectivity
	db, err := sqlx.Open("postgres", psqlInfo)
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

	returnMsg += "The 'EnterpriseNotes' database was created successfully.\n"
	return returnMsg
}

// Create Tables Function
func CreateTables() string {
	var returnMsg string

	//Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Ping the database for connectivity
	db, err := sqlx.Open("postgres", psqlInfo)
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
		userID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
		userName VARCHAR(100), 
		userReadSetting BOOL DEFAULT false, 
		userWriteSetting BOOL DEFAULT false
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatalf("An error occurred when creating the 'users' table.\nGot %s", err)
	}

	returnMsg += "The 'users' table was created successfully.\n"

	// Create the notes table
	sqlQuery = `DROP TABLE IF EXISTS notes;
	CREATE TABLE notes (
		noteID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
		noteName VARCHAR(100), 
		noteText TEXT, 
		noteCompletionTime timestamp DEFAULT CURRENT_TIMESTAMP,
		noteStatus VARCHAR(20) DEFAULT 'none',
		noteDelegation INT,
		noteSharedUsers INT[]
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatalf("An error occurred when creating the 'notes' table.\nGot %s\n", err)
	}

	returnMsg += "The 'notes' table was created successfully.\n"

	// Create the associations table
	sqlQuery = `DROP TABLE IF EXISTS associations;
	CREATE TABLE associations (
		associationID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
		userID INT, 
		noteID INT, 
		associationPerm VARCHAR(20),
		CONSTRAINT fk_user FOREIGN KEY(userID) REFERENCES users(userID) ON DELETE CASCADE,
		CONSTRAINT fk_note FOREIGN KEY(noteID) REFERENCES notes(noteID) ON DELETE CASCADE
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatalf("An error occurred when creating the 'associations' table.\nGot %s\n", err)
	}

	returnMsg += "The 'associations' table was created successfully.\n"
	return returnMsg
}

// Populate Tables Function
func PopulateTables() string {
	var returnMsg string

	// Open the JSON Test Data
	jsonFile, err := os.Open("test_data.json")
	if err != nil {
		log.Printf("An error occurred when reading the JSON file.\nGot %s\n", err)
	}

	defer jsonFile.Close()
	returnMsg += "The JSON file was opened successfully.\n"

	// Unmarshal the JSON Test Data
	byteValue, _ := ioutil.ReadAll(jsonFile)
	data := Data{}
	json.Unmarshal([]byte(byteValue), &data)

	// Add the Test Data to Arrays
	for i := 0; i < len(data.Users); i++ {
		Users = append(Users, data.Users...)
	}
	for i := 0; i < len(data.Notes); i++ {
		Notes = append(Notes, data.Notes...)
	}
	for i := 0; i < len(data.Associations); i++ {
		Associations = append(Associations, data.Associations...)
	}

	// Connect to the database
	const dbname = "enterprisenotes"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Ping the database for connectivity
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	// Truncate the users, notes, and associations tables
	sqlQuery := `TRUNCATE users, notes, associations RESTART IDENTITY CASCADE;`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when truncating the 'users', 'notes', and 'associations' tables.\nGot %s\n", err)
	}

	returnMsg += "The 'users', 'notes', and 'associations' tables were truncated successfully.\n"

	// Populate the users table
	for _, user := range Users {
		createUser(user.Name, user.Read_Setting, user.Write_Setting)
	}

	returnMsg += "The 'users' table was populated successfully.\n"

	// Populate the notes table
	for _, note := range Notes {
		createNote(note.Name, note.Text, note.Completion_Time, note.Status, note.Delegation, note.Shared_Users)
	}

	returnMsg += "The 'notes' table was populated successfully.\n"

	// Create the associations table
	for _, association := range Associations {
		//newID++
		sqlQuery := fmt.Sprintf(`INSERT INTO associations (userID, noteID, associationPerm) VALUES (%d, %d, '%s')`, association.UserID, association.NoteID, association.Permission)
		fmt.Println(sqlQuery)
		_, err = db.Exec(sqlQuery)
		if err != nil {
			log.Printf("An error occurred when populating the 'associations' table.\nGot %s\n", err)
		}
	}

	returnMsg += "The 'associations' table was populated successfully.\n"
	return returnMsg
}
