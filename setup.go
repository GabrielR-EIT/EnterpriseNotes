package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = ""
)

// Create Struct for Test Data
type Data struct {
	Users        []User        `json:"users"`
	Notes        []Note        `json:"notes"`
	Associations []Association `json:"associations"`
}

// Create Database Function
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

// Create Tables Function
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
		userID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
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
		noteID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
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
		associationID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
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

// Populate Tables Function
func PopulateTables() string {
	var returnMsg string

	// Open the JSON Test Data
	jsonFile, err := os.Open("test_data.json")
	if err != nil {
		returnMsg += "An error occurred when reading the JSON file.\n"
		return returnMsg
	}
	returnMsg += "The JSON file was opened successfully.\n"
	defer jsonFile.Close()

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
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Truncate the users, notes, and associations tables
	sqlQuery := `TRUNCATE users, notes, associations RESTART IDENTITY CASCADE;`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when truncating the 'users', 'notes', and 'associations' tables.\n"
		return returnMsg
	}
	returnMsg += "The 'users', 'notes', and 'associations' tables were truncated successfully.\n"

	// Populate the users table
	// for _, user := range Users {
	// 	sqlQuery = fmt.Sprintf(`INSERT INTO users VALUES (%d, '%s', %t, %t)`, NextID(), user.Name, user.Read_Setting, user.Write_Setting)
	// 	fmt.Println(sqlQuery)
	// 	_, err = db.Exec(sqlQuery)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		returnMsg += "An error occurred when populating the 'users' table.\n"
	// 		return returnMsg
	// 	}
	// }
	// returnMsg += "The 'users' table was populated successfully.\n"

	// var newID int
	// newID = 0

	// Populate the users table (ALT)
	for _, user := range Users {
		//newID++
		createUser(user.Name, user.Read_Setting, user.Write_Setting)
		// sqlQuery = fmt.Sprintf(`INSERT INTO users (userName, userReadSetting, userWriteSetting) VALUES ('%s', %t, %t)`, user.Name, user.Read_Setting, user.Write_Setting)
		// fmt.Println(sqlQuery)
		// _, err = db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	returnMsg += "An error occurred when populating the 'users' table.\n"
		// 	return returnMsg
		// }
	}
	returnMsg += "The 'users' table was populated successfully.\n"

	// Populate the notes table
	// for _, note := range Notes {
	// 	sqlQuery := fmt.Sprintf(`INSERT INTO notes VALUES (%d, '%s', '%s', '%s', '%s', '%s', %v)`, NextID(), note.Name, note.Text, note.Completion_Time, note.Status, note.Delegation, note.Shared_Users)
	// 	fmt.Println(sqlQuery)
	// 	_, err = db.Exec(sqlQuery)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		returnMsg += "An error occurred when populating the 'notes' table.\n"
	// 		return returnMsg
	// 	}
	// }
	// returnMsg += "The 'notes' table was populated successfully.\n"

	//newID = 0

	// Populate the notes table (ALT)
	for _, note := range Notes {
		//newID++
		createNote(note.Name, note.Text, note.Completion_Time, note.Status, note.Delegation, note.Shared_Users)
		// sqlQuery := fmt.Sprintf(`INSERT INTO notes (noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers) VALUES ('%s', '%s', '%s', '%s', '%s', '%v')`, note.Name, note.Text, note.Completion_Time, note.Status, note.Delegation, note.Shared_Users)
		// fmt.Println(sqlQuery)
		// _, err = db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	returnMsg += "An error occurred when populating the 'notes' table.\n"
		// 	return returnMsg
		// }
	}
	returnMsg += "The 'notes' table was populated successfully.\n"

	// Create the associations table
	// for _, association := range Associations {
	// 	newID++
	// 	sqlQuery := fmt.Sprintf(`INSERT INTO associations VALUES (%d, %d, %d, '%s')`, NextID(), association.userID, association.noteID, association.permission)
	// 	fmt.Println(sqlQuery)
	// 	_, err = db.Exec(sqlQuery)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		returnMsg += "An error occurred when populating the 'associations' table.\n"
	// 		return returnMsg
	// 	}
	// }
	// returnMsg += "The 'associations' table was populated successfully.\n"

	// return returnMsg

	//newID = 0

	// Create the associations table (ALT)
	for _, association := range Associations {
		//newID++
		sqlQuery := fmt.Sprintf(`INSERT INTO associations (userID, noteID, associationPerm) VALUES (%d, %d, '%s')`, association.UserID, association.NoteID, association.Permission)
		fmt.Println(sqlQuery)
		_, err = db.Exec(sqlQuery)
		if err != nil {
			log.Fatal(err)
			returnMsg += "An error occurred when populating the 'associations' table.\n"
			return returnMsg
		}
	}
	returnMsg += "The 'associations' table was populated successfully.\n"

	return returnMsg
}
