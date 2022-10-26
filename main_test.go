package main

import (
	"fmt"
	"regexp"
	"testing"
	_ "testing"
)

type noteTest struct {
	input string
	id    int
	want  bool
	count int
}

var noteTests = []noteTest{
	noteTest{"agenda", 1, true, 1},
	noteTest{"06-358-8588", 2, true, 1},
	noteTest{"Ocean's Eleven", 3, false, 0},
	noteTest{"APOLOGIES", 4, true, 1},
	noteTest{"robing7@student.eit.ac.nz", 5, false, 0},
}

func TestCreateUser(t *testing.T) {
	//Create Database and Tables for New User Data
	CreateDB()
	CreateTables()

	//Create User Data
	userName := "test user"
	userReadSetting := false
	userWriteSetting := false

	//Create string variable for regexp function
	userData := regexp.MustCompile(fmt.Sprintf("%s|%s|%s|", userName, userReadSetting, userWriteSetting))

	got := createUser(userName, userReadSetting, userWriteSetting)
	want := userData.MatchString(got)

	// if got != want {
	// 	t.Errorf("Got %v, wanted %v", got, want)
	// }
}

func TestReadUser(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	got := readUser(1)
	//want := strings.Contains(got, fmt.Sprintf("%s", Users[1]))

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestUpdateUser(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	userName := "updated user"
	userReadSetting := true
	userWriteSetting := true

	got := updateUser(1, userName, userReadSetting, userWriteSetting)
	//want :=

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestDeleteUser(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	got := deleteUser(1)
	//want :=

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
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

	got := createNote(noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
	//want :=

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestReadNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	got := readNote(1)
	//want :=

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestUpdateNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	//Create New Note Data
	noteID := 1
	noteName := "updated note"
	noteText := "updated text"
	noteCompletionTime := ""
	noteStatus := "in-progress"
	noteDelegation := 2
	noteSharedUsers := []int{6, 2}

	got := updateNote(noteID, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
	//want :=

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestDeleteNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	got := deleteNote(1)
	//want :=

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestFindNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	//Note: Test all five patterns
	//findString := "agenda"

	for _, note := range noteTests {
		got, _ := findNote(note.input)
		want := note.want
		if !got {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	}
}

func TestAnalyseNote(t *testing.T) {
	CreateDB()
	CreateTables()
	PopulateTables()

	for _, note := range noteTests {
		got, _ := analyseNote(note.input, note.id)
		want := note.count
		if got != note.count {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
