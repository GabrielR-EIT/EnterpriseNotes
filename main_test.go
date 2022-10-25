package main

import (
	"testing"
	_ "testing"
)

func TestCreateUser(t *testing.T) {
	//Create Database and Tables for New User Data
	CreateDB()
	CreateTables()

	//Create User Data
	userName := "test user"
	userReadSetting := false
	userWriteSetting := false

	createUser(userName, userReadSetting, userWriteSetting)
}

func TestReadUser(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	readUser(1)
}

func TestUpdateUser(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	userName := "updated user"
	userReadSetting := true
	userWriteSetting := true

	updateUser(1, userName, userReadSetting, userWriteSetting)
}

func TestDeleteUser(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	deleteUser(1)
}

func TestCreateNote(t *testing.T) {
	//Create Database and Tables for New Note Data
	CreateDB()
	CreateTables()

	//Create Note Data
	noteName := "test note"
	noteText := "test text"
	noteCompletionTime := "2022-10-23 00:00:00.000"
	noteStatus := "completed"
	noteDelegation := 1
	noteSharedUsers := []int{6, 1}

	createNote(noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
}

func TestReadNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	readNote(1)
}

func TestUpdateNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	//Create New Note Data
	noteID := 1
	noteName := "updated note"
	noteText := "updated text"
	noteCompletionTime := "2022-10-24 00:00:00.000"
	noteStatus := "in-progress"
	noteDelegation := 2
	noteSharedUsers := []int{6, 2}

	updateNote(noteID, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
}

func TestDeleteNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	deleteNote(1)
}

func TestFindNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	//Note: Test all five patterns
	findString := "agenda"

	findNote(findString)
}

func TestAnalyseNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	findString := "agenda"

	analyseNote(findString, 1)
}

func TestMain(t *testing.T) {
	main()
}
