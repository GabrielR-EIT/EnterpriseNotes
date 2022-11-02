package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/jmoiron/sqlx"
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
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println(err)
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
		t.Errorf("createUser() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("createUser() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestReadUser(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

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

	var user User
	db.QueryRowx(`SELECT * FROM users WHERE userID = 1`).StructScan(&user)
	got := readUser(1)
	want := fmt.Sprintf("User details:\n%v\n", user)

	if got != want {
		t.Errorf("readUser() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("readUser() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestUpdateUser(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

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

	// Create New User Data
	userName := "updated user"
	userReadSetting := true
	userWriteSetting := true

	got := updateUser(1, userName, userReadSetting, userWriteSetting)
	want := "The user information has been successfully updated."

	if got != want {
		t.Errorf("updateUser() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("updateUser() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestDeleteUser(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

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

	got := deleteUser(1)
	want := "The record for user with ID 1 has been successfully deleted."

	if got != want {
		t.Errorf("deleteUser() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("deleteUser() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestCreateNote(t *testing.T) {
	// Create Database and Tables for New Note Data
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())

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

	// Create Note Data
	test_note := Note{
		Name:            "test note",
		Text:            "test text",
		Completion_Time: "CURRENT_TIMESTAMP",
		Status:          "completed",
		Delegation:      1,
		Shared_Users:    "[6, 1]",
	}

	got := createNote(test_note.Name, test_note.Text, test_note.Completion_Time, test_note.Status, test_note.Delegation, test_note.Shared_Users)
	want := fmt.Sprintf("Your new note has been successfully added.\nDetails:\n%v\nThere are now %v notes in the database.", test_note, strconv.Itoa(len(Notes)))

	if got != want {
		t.Errorf("createNote() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("createNote() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestReadNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

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
	db.QueryRowx(`SELECT * FROM notes WHERE noteID = 1`).StructScan(&note)
	got := readNote(1)
	want := fmt.Sprintf("Note details:\n%v\n", note)

	if got != want {
		t.Errorf("readNote() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("readNote() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestUpdateNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

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

	//Create New Note Data
	noteID := 1
	noteName := "updated note"
	noteText := "updated text"
	noteCompletionTime := "CURRENT_TIMESTAMP"
	noteStatus := "in-progress"
	noteDelegation := 2
	noteSharedUsers := "[6, 2]"

	got := updateNote(noteID, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
	want := "The user information has been successfully updated."

	if got != want {
		t.Errorf("updateNote() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("updateNote() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestDeleteNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

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

	got := deleteNote(1)
	want := "The record for note with ID 1 has been successfully deleted."

	if got != want {
		t.Errorf("deleteNote() Test: Got %v, wanted %v", got, want)
	} else {
		t.Logf("deleteNote() Test: Got %v, wanted %v. Test was successful.", got, want)
	}
}

func TestFindNote(t *testing.T) {
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())

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

	for _, note := range noteTests {
		got, _ := findNote(note.input)
		want := note.want
		if !got {
			t.Errorf("findNote() Test %d: Got %t, wanted %t", note.id, got, want)
		} else {
			t.Logf("findNote() Test %d: Got %v, wanted %v. Test was successful.", note.id, got, want)
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
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	for _, note := range noteTests {
		got := analyseNote(note.input, note.id)
		want := fmt.Sprintf("The analysis returned %v instances of \"%s\" in the text.", note.count, note.input)
		if got != want {
			t.Errorf("analyseNote() Test %d: Got %s, wanted %s", note.id, got, want)
		} else {
			t.Logf("analyseNote() Test %d: Got %s, wanted %s. Test was successful.", note.id, got, want)
		}
	}
}

func TestValidatePattern(t *testing.T) {
	type patternTest struct {
		input   string
		pattern string
		id      int
		want    bool
	}

	var patternTests = []patternTest{
		{"pre", "[a-zA-z]+", 1, true},
		{"123", "[a-zA-z]+", 2, false},
		{"01189998819991197253", "[0-9\\W]", 3, true},
		{"this is not a phone number", "[0-9\\W]", 4, false},
		{"robing7@student.eit.ac.nz", "@{1}", 5, true},
		{"01100101 01101101 01100001 01101001 01101100", "@{1}", 6, false},
		{"action meeting apologies", "meeting|minutes|agenda|action|attendees|apologies{3,}", 7, true},
		{"action meeting sorry", "meeting|minutes|agenda|action|attendees|apologies{3,}", 8, false},
		{"UPPERCASE", "[A-Z]{3,}", 9, true},
		{"UPPERcase", "[A-Z]{3,}", 10, false},
	}

	for _, pattern := range patternTests {
		got, _ := validatePattern(pattern.pattern, pattern.input)
		want := pattern.want
		if got != want {
			t.Errorf("validatePattern() Test %d: Got %t, wanted %t", pattern.id, got, want)
		} else {
			t.Logf("validatePattern() Test %d: Got %t, wanted %t. Test was successful.", pattern.id, got, want)
		}
	}
}

func TestValidateStatus(t *testing.T) {
	type statusTest struct {
		input string
		id    int
		want  bool
	}

	var statusTests = []statusTest{
		{"none", 1, true},
		{"in progress", 2, true},
		{"completed", 3, true},
		{"cancelled", 4, true},
		{"delegated", 5, true},
		{"invalid", 6, false},
	}

	for _, status := range statusTests {
		got, _ := validateStatus(status.input)
		want := status.want
		if got != want {
			t.Errorf("validateStatus() Test %d: Got %t, wanted %t", status.id, got, want)
		} else {
			t.Logf("validateStatus() Test %d: Got %t, wanted %t. Test was successful.", status.id, got, want)
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
