package main

import (
	"embed"
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

//go:embed templates
var tmplEmbed embed.FS

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

// Middleware to connect the database for each request that uses this
// middleware.
func connectDatabase(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
	}
}

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

	// //Connect to the database
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// // Ping the database for connectivity
	// db, err := sqlx.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

	// // Connect to the database
	// const dbname = "enterprisenotes"
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// // Ping the database for connectivity
	// db, err := sqlx.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	log.Println(err)
	// }

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
		createNote(db, note.Name, note.Text, note.Completion_Time, note.Status, note.Delegation, note.Shared_Users)
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

// Return all users as JSON
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Users)
}

// Return all notes as JSON
func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Notes)
}

// --- HTTP Server Functions --- //
// Start New Server Function
func StartServer(router *gin.Engine, db *sqlx.DB) string {
	returnMsg := ""

	tmpl := template.Must(template.ParseFS(tmplEmbed, "templates/*/*.html"))
	router.SetHTMLTemplate(tmpl)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.GET("/users", webFunctions.ReadUsers(db))
	// router.GET("/users/:guid", webFunctions.ReadUser(db))
	// router.POST("/users", webFunctions.CreateUser(db))
	// router.DELETE("/users/:guid", webFunctions.DeleteUser(db))
	// router.PUT("/users/:guid", webFunctions.UpdateUser(db))

	// router.GET("/notes", webFunctions.ReadNotes(db))
	// router.GET("/notes/:guid", webFunctions.ReadNote(db))
	// router.POST("/notes", webFunctions.CreateNote(db))
	// router.DELETE("/notes/:guid", webFunctions.DeleteNote(db))
	// router.PUT("/notes/:guid", webFunctions.UpdateNote(db))

	// router.Static("/css", "./static/css")
	// router.Static("/img", "./static/img")
	// router.Static("/scss", "./static/scss")
	// router.Static("/vendor", "./static/vendor")
	// router.Static("/js", "./static/js")
	// router.StaticFile("/favicon.ico", "./img/favicon.ico")

	// router.LoadHTMLFiles(
	// 	"./templates/views/users.html",
	// 	"./templates/views/notes.html",
	// )
	router.LoadHTMLGlob("templates/**/*")

	router.Use(connectDatabase(db))
	router.GET("/users", func(c *gin.Context) {
		users := []User{}
		c.HTML(http.StatusOK, "views/users.html", gin.H{
			"Users": users,
		})
	})
	router.GET("/notes", func(c *gin.Context) {
		notes := []Note{}
		c.HTML(http.StatusOK, "views/notes.html", gin.H{
			"Notes": notes,
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/users/")
	})

	//controller.Router(router)

	log.Println("Server started")
	log.Fatalln(router.Run(serverPort)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	returnMsg += "The server has been successfully started."
	return returnMsg
}
