package main

import (
	_ "enterprise-notes/webFunctions"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// --- User CRUD Functions ---//
// Create User Function
func createUser(db *sqlx.DB, userName string, userReadSetting bool, userWriteSetting bool) string {
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`INSERT INTO users(userName, userReadSetting, userWriteSetting) VALUES ('%s', %t, %t);`, userName, userReadSetting, userWriteSetting)
	_, err := db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to create the user.\nGot %s\n", err)
	}

	// Create struct for new user
	newUser := User{
		Name:          userName,
		Read_Setting:  userReadSetting,
		Write_Setting: userWriteSetting,
	}

	// Add new note struct to Notes slice
	Users = append(Users, newUser)
	returnMsg += fmt.Sprintf("A new user has been successfully added.\nDetails:\n%v\nThere are now %v users in the database.\n", newUser, strconv.Itoa(len(Users)))
	return returnMsg
}

// Read User Function
func readUser(db *sqlx.DB, userID int) string {
	returnMsg := ""

	var user User
	sqlQuery := fmt.Sprintf(`SELECT * FROM users WHERE userID = %d`, userID)
	queryRow := db.Get(&user, sqlQuery)
	if queryRow != nil {
		log.Printf("An error occurred when reading user information.\nGot %s\n", queryRow)
	}

	returnMsg += fmt.Sprintf("User details:\n%v\n", user)
	return returnMsg
}

// Update User Function
func updateUser(db *sqlx.DB, userID int, userName string, userReadSetting bool, userWriteSetting bool) string {
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`UPDATE users SET userName = '%s', userReadSetting = %t, userWriteSetting = %t WHERE userID = %d`, userName, userReadSetting, userWriteSetting, userID)
	_, err := db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when updating the user information.\nGot %s\n", err)
	}

	returnMsg += "The user information has been successfully updated."
	return returnMsg
}

// Delete User Function
func deleteUser(db *sqlx.DB, userID int) string {
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`DELETE FROM users WHERE userID = %d;`, userID)
	_, err := db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to delete the user.\nGot %s\n", err)
	}

	//set the values of the user to null to remove it from the slice
	Users[userID-1] = User{}
	returnMsg += fmt.Sprintf("The record for user with ID %d has been successfully deleted.", userID)
	return returnMsg
}

// --- Note CRUD Functions ---//
// Create Note Function
func createNote(db *sqlx.DB, noteName string, noteText string, noteStatus string, noteDelegation int, noteSharedUsers string) string {
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`INSERT INTO notes(noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers) VALUES ('%s', '%s', CURRENT_TIMESTAMP, '%s', %d, ARRAY%s)`, noteName, noteText, noteStatus, noteDelegation, noteSharedUsers)
	_, err := db.Exec(sqlQuery)
	if err != nil {
		log.Println("An error occurred when creating a new note.\nGot\n", err)
	}

	noteCompletionTime := fmt.Sprint(time.Now().Round(60 * time.Second))

	// Create struct for new note
	newNote := Note{
		Name:            noteName,
		Text:            noteText,
		Completion_Time: noteCompletionTime,
		Status:          noteStatus,
		Delegation:      noteDelegation,
		Shared_Users:    noteSharedUsers,
	}

	// Add new note struct to Notes slice
	Notes = append(Notes, newNote)

	returnMsg += fmt.Sprintf("Your new note has been successfully added.\nDetails:\n%v\nThere are now %v notes in the database.", newNote, strconv.Itoa(len(Notes)))
	return returnMsg

	// Create an association between the user and note
}

// Read Note Function
func readNote(db *sqlx.DB, noteID int) string {
	returnMsg := ""

	var note Note
	sqlQuery := fmt.Sprintf(`SELECT * FROM notes WHERE noteID = %d`, noteID)
	queryRow := db.Get(&note, sqlQuery)
	if queryRow != nil {
		log.Printf("An error occurred when reading the note.\nGot %s\n", queryRow)
	}

	returnMsg += fmt.Sprintf("Note details:\n%v\n", note)
	return returnMsg
}

// Update Note Function
func updateNote(db *sqlx.DB, noteID int, noteName string, noteText string, noteStatus string, noteDelegation int, noteSharedUsers string) string {
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`UPDATE notes SET noteName = '%s', noteText = '%s', noteStatus = '%s', noteDelegation = %d, noteSharedUsers = ARRAY%s WHERE noteID = %d`, noteName, noteText, noteStatus, noteDelegation, noteSharedUsers, noteID)
	_, err := db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when updating the note.\nGot %s\n", err)
	}

	returnMsg += "The note has been successfully updated."
	return returnMsg
}

// Delete Note Function
func deleteNote(db *sqlx.DB, noteID int) string {
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`DELETE FROM notes WHERE noteID = %d;`, noteID)
	_, err := db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to delete the note.\nGot %s\n", err)
	}

	//set the values of the note to null to remove it from the slice
	Notes[noteID-1] = Note{}

	returnMsg += fmt.Sprintf("The record for note with ID %d has been successfully deleted.", noteID)
	return returnMsg
}

// --- Extended Note Functions ---//
// Find Note Function
func findNote(db *sqlx.DB, inputPattern string) (bool, string) {
	result := false
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`SELECT * FROM notes WHERE noteText ~ '%s';`, inputPattern)
	queryRows, err := db.Query(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to find note text matching the given pattern.\nGot %s\n", err)
	}

	result = true
	returnMsg += fmt.Sprintf("At least one match was successfully found for that pattern. Result:\n%v\n", queryRows)
	return result, returnMsg
}

// Analyse Note Function
func analyseNote(db *sqlx.DB, inputPattern string, noteID int) string {
	returnMsg := ""

	sqlQuery := fmt.Sprintf(`SELECT count(*) FROM notes CROSS JOIN LATERAL regexp_matches(noteText, '%s', 'g') WHERE noteID = %d;`, inputPattern, noteID)
	queryRows, err := db.Query(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to find text matching the given pattern.\nGot %s\n", err)
	}
	queryCount := 0
	for queryRows.Next() {
		err := queryRows.Scan(&queryCount)
		if err != nil {
			log.Printf("An error occurred when trying to retrieve the count of pattern matches.\nGot %s\n", err)
		}
	}

	returnMsg += fmt.Sprintf("The analysis returned %v instances of \"%s\" in the text.", queryCount, inputPattern)
	return returnMsg
}

// --- Main ---//
func main() {
	// Create the Enterprise Notes database
	fmt.Print(CreateDB())

	// Connect to the database
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

	fmt.Print(CreateTables(db))
	fmt.Print(PopulateTables(db))

	// Start the Gin router
	router := gin.New()
	go fmt.Print(StartServer(router, db))
}
