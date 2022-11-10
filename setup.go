package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	// Database Info
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "enterprisenotes"

	// HTTP Info
	page       = "/Enterprise_Notes/"
	serverPort = ":8080"
)

// --- Database Functions --- //
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
func CreateTables(db *sqlx.DB) string {
	var returnMsg string

	// Create the users table
	sqlQuery := `DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		userID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
		userName VARCHAR(100), 
		userReadSetting BOOL DEFAULT false, 
		userWriteSetting BOOL DEFAULT false
	);`
	_, err := db.Exec(sqlQuery)
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
func PopulateTables(db *sqlx.DB) string {
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

	// Add the Test Data to Slices
	Users = append(Users, data.Users...)
	Notes = append(Notes, data.Notes...)
	Associations = append(Associations, data.Associations...)

	// Truncate the users, notes, and associations tables
	sqlQuery := `TRUNCATE users, notes, associations RESTART IDENTITY CASCADE;`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when truncating the 'users', 'notes', and 'associations' tables.\nGot %s\n", err)
	}

	returnMsg += "The 'users', 'notes', and 'associations' tables were truncated successfully.\n"

	// Populate the users table
	for _, user := range Users {
		createUser(db, user.Name, user.Read_Setting, user.Write_Setting)
	}

	returnMsg += "The 'users' table was populated successfully.\n"

	// Populate the notes table
	for _, note := range Notes {
		createNote(db, note.Name, note.Text, note.Status, note.Delegation, note.Shared_Users)
	}

	returnMsg += "The 'notes' table was populated successfully.\n"

	// Create the associations table
	for _, association := range Associations {
		//newID++
		sqlQuery := fmt.Sprintf(`INSERT INTO associations (userID, noteID, associationPerm) VALUES (%d, %d, '%s')`, association.UserID, association.NoteID, association.Permission)
		_, err = db.Exec(sqlQuery)
		if err != nil {
			log.Printf("An error occurred when populating the 'associations' table.\nGot %s\n", err)
		}
	}

	returnMsg += "The 'associations' table was populated successfully.\n"
	return returnMsg
}

// --- API Functions --- //

// --- HTTP Server Functions --- //
// Start New Server Function
func StartServer(router *gin.Engine, db *sqlx.DB) string {
	returnMsg := ""

	tmpl := template.Must(template.ParseFS(tmplEmbed, "templates/*/*.html"))
	router.SetHTMLTemplate(tmpl)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/**/*")

	router.Use(connectDatabase(db))
	router.GET("/users", handlerReadUsers)
	router.GET("/notes", handlerReadNotes)
	router.POST("/users", handlerCreateUser)
	router.POST("/notes", handlerCreateNote)
	router.POST("/users-put", handlerUpdateUser)
	router.POST("/notes-put", handlerUpdateNote)
	// Redirect the user to the Users page
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/users")
	})

	log.Println("Server started")
	log.Fatalln(router.Run(serverPort)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	returnMsg += "The server has been successfully started."
	return returnMsg
}
