package main

import (
	"fmt"
	"log"
	_ "log"
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// Create Structs
type User struct {
	ID            int    `json:"user_id"`
	Name          string `json:"name"`
	Read_Setting  bool   `json:"read_setting"`
	Write_Setting bool   `json:"write_setting"`
}

type Note struct {
	ID              int    `json:"note_id"`
	Name            string `json:"name"`
	Text            string `json:"text"`
	Completion_Time string `json:"completion_date"`
	Status          string `json:"status"`
	//Delegation field should be a User struct
	Delegation int `json:"delegation"`
	// Shared_Users should be a slice of Users
	Shared_Users string `json:"shared_users"`
}

type Association struct {
	ID         int    `json:"association_id"`
	UserID     int    `json:"user_id"`
	NoteID     int    `json:"note_id"`
	Permission string `json:"permission"`
}

// Create Slices
var Users []User
var Notes []Note
var Associations []Association

// Create User Function
func createUser(userName string, userReadSetting bool, userWriteSetting bool) string {
	returnMsg := ""

	// Connect to the database
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

	sqlQuery := fmt.Sprintf(`INSERT INTO users(userName, userReadSetting, userWriteSetting) VALUES ('%s', %t, %t);`, userName, userReadSetting, userWriteSetting)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to create the user.\nGot %s\n", err)
	}

	// Create struct for new user
	new_user := User{
		Name:          userName,
		Read_Setting:  userReadSetting,
		Write_Setting: userWriteSetting,
	}

	// Add new note struct to Notes slice
	Users = append(Users, new_user)
	returnMsg += fmt.Sprintf("A new user has been successfully added.\nDetails:\n%v\nThere are now %v users in the database.\n", new_user, strconv.Itoa(len(Users)))
	return returnMsg
}

// Read User Function
func readUser(userID int) string {
	returnMsg := ""

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

	var user User
	sqlQuery := fmt.Sprintf(`SELECT * FROM users WHERE userID = %d`, userID)
	queryRow := db.QueryRowx(sqlQuery).StructScan(&user)
	if queryRow != nil {
		log.Printf("An error occurred when reading user information.\nGot %s\n", queryRow)
	}

	returnMsg += fmt.Sprintf("User details:\n%v\n", user)
	return returnMsg
}

// Update User Function
func updateUser(userID int, userName string, userReadSetting bool, userWriteSetting bool) string {
	returnMsg := ""

	// Connect to the database
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

	sqlQuery := fmt.Sprintf(`UPDATE users SET userName = '%s', userReadSetting = %t, userWriteSetting = %t WHERE userID = %d`, userName, userReadSetting, userWriteSetting, userID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when updating the user information.\nGot %s\n", err)
	}

	returnMsg += "The user information has been successfully updated."
	return returnMsg
}

// Delete User Function
func deleteUser(userID int) string {
	returnMsg := ""

	// Connect to the database
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

	sqlQuery := fmt.Sprintf(`DELETE FROM users WHERE userID = %d;`, userID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to delete the user.\nGot %s\n", err)
	}

	//set the values of the user to null to remove it from the slice
	Users[userID-1] = User{}
	returnMsg += fmt.Sprintf("The record for user with ID %d has been successfully deleted.", userID)
	return returnMsg
}

// Create Note Function
func createNote(noteName string, noteText string, noteCompletionTime string, noteStatus string, noteDelegation int, noteSharedUsers string) string {
	returnMsg := ""

	// Connect to the database
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

	sqlQuery := fmt.Sprintf(`INSERT INTO notes(noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers) VALUES ('%s', '%s', '%s', '%s', %d, ARRAY%s)`, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Println("An error occurred when creating a new note.\nGot\n", err)
	}

	// Create struct for new note
	new_note := Note{
		Name:            noteName,
		Text:            noteText,
		Completion_Time: noteCompletionTime,
		Status:          noteStatus,
		Delegation:      noteDelegation,
		Shared_Users:    noteSharedUsers,
	}

	// Add new note struct to Notes slice
	Notes = append(Notes, new_note)

	returnMsg += fmt.Sprintf("Your new note has been successfully added.\nDetails:\n%v\nThere are now %v notes in the database.", new_note, strconv.Itoa(len(Notes)))
	return returnMsg

	// Create an association between the user and note
}

// Read Note Function
func readNote(noteID int) string {
	returnMsg := ""

	// Connect to the database
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

	var note Note
	sqlQuery := fmt.Sprintf(`SELECT * FROM notes WHERE noteID = %d`, noteID)
	queryRow := db.QueryRowx(sqlQuery).StructScan(&note)
	if queryRow != nil {
		log.Printf("An error occurred when reading the note.\nGot %s\n", queryRow)
	}

	returnMsg += fmt.Sprintf("Note details:\n%v\n", note)
	return returnMsg
}

// Update Note Function
func updateNote(noteID int, noteName string, noteText string, noteCompletionTime string, noteStatus string, noteDelegation int, noteSharedUsers string) string {
	returnMsg := ""

	// Connect to the database
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

	sqlQuery := fmt.Sprintf(`UPDATE notes SET noteName = '%s', noteText = '%s', noteCompletionTime = '%s', noteStatus = '%s', noteDelegation = %d, noteSharedUsers = ARRAY%s WHERE noteID = %d`, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers, noteID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when updating the note.\nGot %s\n", err)
	}

	returnMsg += "The note has been successfully updated."
	return returnMsg
}

// Delete Note Function
func deleteNote(noteID int) string {
	returnMsg := ""

	// Connect to the database
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

	sqlQuery := fmt.Sprintf(`DELETE FROM notes WHERE noteID = %d;`, noteID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to delete the note.\nGot %s\n", err)
	}

	//set the values of the note to null to remove it from the slice
	Notes[noteID-1] = Note{}

	returnMsg += fmt.Sprintf("The record for note with ID %d has been successfully deleted.", noteID)
	return returnMsg
}

// Find Note Function
func findNote(inputPattern string) (bool, string) {
	result := false
	returnMsg := ""

	// Connect to the database
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
func analyseNote(inputPattern string, noteID int) string {
	returnMsg := ""

	// Connect to the database
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
	// go StartServer()
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())
}
