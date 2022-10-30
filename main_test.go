package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
	{"agenda", 1, true, 1},
	{"06-358-8588", 2, true, 1},
	{"Oceans Eleven", 3, false, 0},
	{"APOLOGIES", 4, true, 1},
	{"robing7@student.eit.ac.nz", 5, false, 0},
}

func TestCreateUser(t *testing.T) {
	//Create Database and Tables for New User Data
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())

	// Connect to the database
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

	//Create User Data
	test_user := User{
		Name:          "test user",
		Read_Setting:  false,
		Write_Setting: false,
	}

	//Create string variable for regexp function
	// userData := regexp.MustCompile(fmt.Sprintf("%s|%s|%s|", userName, userReadSetting, userWriteSetting))

	got := fmt.Sprint(createUser(test_user.Name, test_user.Read_Setting, test_user.Write_Setting))
	want := fmt.Sprintf("A new user has been successfully added.\nDetails:\n%v\nThere are now %v users in the database.\n", test_user, strconv.Itoa(len(Users)))

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. CreateUser() test was successful.", got, want)
	}
}

func TestReadUser(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	got := readUser(1)
	want := fmt.Sprintf("User details:\n%v\n", db.QueryRow(`SELECT * FROM users WHERE userID = 1`))

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. ReadUser() test was successful.", got, want)
	}
}

func TestUpdateUser(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	// Create New User Data
	userName := "updated user"
	userReadSetting := true
	userWriteSetting := true

	got := updateUser(1, userName, userReadSetting, userWriteSetting)
	want := "The user information has been successfully updated."

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. UpdateUser() test was successful.", got, want)
	}
}

func TestDeleteUser(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	got := deleteUser(1)
	want := "The record for user with ID 1 has been successfully deleted."

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. DeleteUser() test was successful.", got, want)
	}
}

func TestCreateNote(t *testing.T) {
	// Create Database and Tables for New Note Data
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())

	// Connect to the database
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

	// Create Note Data
	test_note := Note{
		Name:            "test note",
		Text:            "test text",
		Completion_Time: "CURRENT_TIMESTAMP()",
		Status:          "completed",
		Delegation:      1,
		Shared_Users:    "[6, 1]",
	}

	got := createNote(test_note.Name, test_note.Text, test_note.Completion_Time, test_note.Status, test_note.Delegation, test_note.Shared_Users)
	want := fmt.Sprintf("Your new note has been successfully added.\nDetails:\n%v\nThere are now %v notes in the database.", test_note, strconv.Itoa(len(Notes)))

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. CreateNote() test was successful.", got, want)
	}
}

func TestReadNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	got := readNote(1)
	want := fmt.Sprintf("Note details:\n%v\n", db.QueryRow(`SELECT * FROM notes WHERE noteID = %d`))

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. ReadNote() test was successful.", got, want)
	}
}

func TestUpdateNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	//Create New Note Data
	noteID := 1
	noteName := "updated note"
	noteText := "updated text"
	noteCompletionTime := ""
	noteStatus := "in-progress"
	noteDelegation := 2
	noteSharedUsers := "[6, 2]"

	got := updateNote(noteID, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
	want := "The user information has been successfully updated."

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. UpdateNote() test was successful.", got, want)
	}
}

func TestDeleteNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	got := deleteNote(1)
	want := "The record for note with ID 1 has been successfully deleted."

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	} else {
		t.Logf("Got %v, wanted %v. DeleteNote() test was successful.", got, want)
	}
}

func TestFindNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	testNum := 0

	for _, note := range noteTests {
		got, _ := findNote(note.input)
		want := note.want
		testNum++
		if !got {
			t.Errorf("FindNote() Test %d: Got %t, wanted %t", testNum, got, want)
		} else {
			t.Logf("FindNote() Test %d: Got %v, wanted %v. Test was successful.", testNum, got, want)
		}
	}
}

func TestAnalyseNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

	// Connect to the database
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

	for _, note := range noteTests {
		got := analyseNote(note.input, note.id)
		want := fmt.Sprintf("The analysis returned %v instances of \"%s\" in the text.", note.count, note.input)
		if got != want {
			t.Errorf("Got %s, wanted %s", got, want)
		} else {
			t.Logf("Got %s, wanted %s. AnalyseNote() Test was successful.", got, want)
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
