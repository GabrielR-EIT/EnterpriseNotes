package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"log"
	_ "log"
	"strconv"

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
	Shared_Users []int `json:"shared_users"`
}

type Association struct {
	ID         int    `json:"association_id"`
	UserID     int    `json:"user_id"`
	NoteID     int    `json:"note_id"`
	Permission string `json:"permission"`
}

// Create Maps
var Patterns = map[string]string{`[a-zA-z]+`: "A sentence with a given prefix and/or suffix",
	`[0-9\W]`: "A phone number with a given area code and optionally a consecutive sequence of numbers that are part of that number",
	`@{1}`:    "An email address on a domain that is only partially provided",
	`meeting|minutes|agenda|action|attendees|apologies{3,}`: "Text that contains at least three of the following case-insensitive words: meeting, minutes, agenda, action, attendees, apologies",
	`[A-Z]{3,}`: "A word in all capitals of three characters or more"}

// Create Slices
var Users []User
var Notes []Note
var Associations []Association

// Set Global Variables
//var optionSelect = true

// --- Functions ---//
// Get Index Of Item Function
// func getIndex(s []struct{}, i struct{}) int {
// 	for k, v := range s {
// 		if i == v {
// 			return k
// 		}
// 	}
// 	return -1
// }

// // Remove From Slice Function
// func removeFromSlice(s []struct{}, i int) []struct{} {
// 	var newSlice = make([]struct{}, 0)
// 	newSlice = append(newSlice, s[:i]...)
// 	return append(newSlice, s[i+1:]...)
// }

// Create User Function
func createUser(userName string, userReadSetting bool, userWriteSetting bool) string {
	returnMsg := ""

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

	sqlQuery := fmt.Sprintf(`INSERT INTO users VALUE (%s, %t, %t);`, userName, userReadSetting, userWriteSetting)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when trying to create the user.\n"
	}

	// user_name := ""
	// var user_read_setting, user_write_setting bool
	// fmt.Scanln(&user_name)
	// fmt.Scanln(&user_read_setting)
	// fmt.Scanln(&user_write_setting)

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
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := fmt.Sprintf(`SELECT * FROM users WHERE userID = %d`, userID)
	queryRow := db.QueryRow(sqlQuery)
	if queryRow != nil {
		log.Fatal(queryRow)
		returnMsg += "An error occurred when reading user information.\n"
		return returnMsg
	}

	returnMsg += fmt.Sprintf("User details:\n%v\n", queryRow)
	return returnMsg
}

// Update User Function
func updateUser(userID int, userName string, userReadSetting bool, userWriteSetting bool) string {
	returnMsg := ""

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

	// println("Details:\n", user, "\nDo you wish to edit:\n1. Name\n2. Read Setting\n3. Write Setting\nOr type N to cancel.")
	// fmt.Scanln(input)
	// switch input {
	// case "1":
	// 	fmt.Scanln(input)
	// 	user.Name = input
	// 	//sqlQuery := fmt.Sprintf(`UPDATE users SET name = %v WHERE ID = %v;`, input, user.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the user.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The name for this user has been changed to '%v'", input)
	// 	return returnMsg
	// case "2":
	// 	fmt.Scanln(input)
	// 	//user.Read_Setting = strconv.ParseBool(input)
	// 	//sqlQuery := fmt.Sprintf(`UPDATE users SET read_setting = %v WHERE ID = %v;`, input, user.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the user.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The read setting for this user has been changed to '%v'", input)
	// 	return returnMsg
	// case "3":
	// 	fmt.Scanln(input)
	// 	//user.Write_Setting = strconv.ParseBool(input)
	// 	//sqlQuery := fmt.Sprintf(`UPDATE users SET write_setting = %v WHERE ID = %v;`, input, user.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the user.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The write setting for this user has been changed to '%v'", input)
	// 	return returnMsg
	// default:
	// 	return ""
	// }

	sqlQuery := fmt.Sprintf(`UPDATE users SET userName = %s, userReadSetting = %t, userWriteSetting = %t WHERE userID = %d`, userName, userReadSetting, userWriteSetting, userID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when updating the user information.\n"
		return returnMsg
	}

	returnMsg += fmt.Sprintf("The user information has been successfully updated.")
	return returnMsg
}

// Delete User Function
func deleteUser(userID int) string {
	returnMsg := ""

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

	//sqlQuery := fmt.Sprintf(`DELETE FROM users WHERE ID = %v;`, user.ID)
	// _, err := db.Exec(sqlQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("An error occurred when trying to delete the user.")
	// }

	//set the values of the user to null
	// user = User{}
	// returnMsg += fmt.Sprintf("The information for user '%v' has been deleted.", user.Name)
	// return returnMsg

	sqlQuery := fmt.Sprintf(`DELETE * FROM users WHERE ID = %v;`, userID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when trying to delete the user.\n"
	}

	//set the values of the user to null to remove it from the slice
	Users[userID-1] = User{}
	returnMsg += fmt.Sprintf("The record for user with ID '%v' has been successfully deleted.", userID)
	return returnMsg
}

// Create Note Function
func createNote(noteName string, noteText string, noteCompletionTime string, noteStatus string, noteDelegation int, noteSharedUsers []int) string {
	returnMsg := ""
	// note_name, note_text, note_time, note_status, note_delegation, note_users := "", "", "", "", "", []string{}
	// fmt.Scanln(&note_name)
	// fmt.Scanln(&note_text)
	// fmt.Scanln(&note_time)
	// fmt.Scanln(&note_status)
	// fmt.Scanln(&note_delegation)
	// fmt.Scanln(&note_users)

	//sqlQuery := fmt.Sprintf(`INSERT INTO notes VALUE (%v, %v, %v, %v, %v, %v);`, new_note.Name, new_note.Text, new_note.Completion_Time, new_note.Status, new_note.Delegation, new_note.Shared_Users)
	// _, err := db.Exec(sqlQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("An error occurred when trying to create the note.")
	// }

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

	sqlQuery := fmt.Sprintf(`INSERT INTO notes VALUES ('%s', '%s', '%s', '%s', '%s', '%v')`, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when creating a new note.\n"
		return returnMsg
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
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := fmt.Sprintf(`SELECT * FROM notes WHERE noteID = %d`, noteID)
	queryRow := db.QueryRow(sqlQuery)
	if queryRow != nil {
		log.Fatal(queryRow)
		returnMsg += "An error occurred when reading the note.\n"
		return returnMsg
	}
	// println(note)
	returnMsg += fmt.Sprintf("Note details:\n%v\n", queryRow)
	return returnMsg
}

// Update Note Function
func updateNote(noteID int, noteName string, noteText string, noteCompletionTime string, noteStatus string, noteDelegation int, noteSharedUsers []int) string {
	returnMsg := ""

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

	// println("Details:\n", note, "\nDo you wish to edit:\n1. Name\n2. Text\n3. Completion Time\n4. Status\n5. Delegation\n6. Shared Users\nOr type N to cancel.")
	// fmt.Scanln(input)
	// switch input {
	// case "1":
	// 	fmt.Scanln(input)
	// 	note.Name = input
	// 	//sqlQuery := fmt.Sprintf(`UPDATE notes SET name = %v WHERE ID = %v;`, input, note.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the note.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The name for this note has been changed to '%v'\n", input)
	// 	return returnMsg
	// case "2":
	// 	fmt.Scanln(input)
	// 	note.Text = input
	// 	//sqlQuery := fmt.Sprintf(`UPDATE notes SET text = %v WHERE ID = %v;`, input, note.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the note.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The text for this note has been changed to '%v'\n", input)
	// 	return returnMsg
	// case "3":
	// 	fmt.Scanln(input)
	// 	note.Completion_Time = input
	// 	//sqlQuery := fmt.Sprintf(`UPDATE notes SET completion_time = %v WHERE ID = %v;`, input, note.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the note.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The completion time for this note has been changed to '%v'\n", input)
	// 	return returnMsg
	// case "4":
	// 	fmt.Scanln(input)
	// 	note.Status = input
	// 	//sqlQuery := fmt.Sprintf(`UPDATE notes SET status = %v WHERE ID = %v;`, input, note.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the note.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The status for this note has been changed to '%v'\n", input)
	// 	return returnMsg
	// case "5":
	// 	fmt.Scanln(input)
	// 	note.Delegation = input
	// 	//sqlQuery := fmt.Sprintf(`UPDATE notes SET delegation = %v WHERE ID = %v;`, input, note.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the note.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The delegation for this note has been changed to '%v'\n", input)
	// 	return returnMsg
	// case "6":
	// 	fmt.Scanln(input)
	// 	// note.Shared_Users = input
	// 	//sqlQuery := fmt.Sprintf(`UPDATE notes SET shared_users = %v WHERE ID = %v;`, input, note.ID)
	// 	// _, err := db.Exec(sqlQuery)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// 	fmt.Println("An error occurred when trying to update the note.")
	// 	// }
	// 	returnMsg += fmt.Sprintf("The shared users for this note have been changed to '%v'\n", input)
	// 	return returnMsg
	// default:
	// 	return ""
	// }

	sqlQuery := fmt.Sprintf(`UPDATE notes SET noteName = %s, noteText = %s, noteCompletionTime = %v, noteStatus = %s, noteDelegation = %s, noteSharedUsers = %v WHERE noteID = %d`, noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers, noteID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when updating the note.\n"
		return returnMsg
	}
	// println(note)
	returnMsg += fmt.Sprintf("The note has been successfully updated.")

	return returnMsg
}

// Delete Note Function
func deleteNote(noteID int) string {
	returnMsg := ""

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

	sqlQuery := fmt.Sprintf(`DELETE * FROM notes WHERE ID = %v;`, noteID)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when trying to delete the note.\n"
	}

	//set the values of the note to null to remove it from the slice
	Notes[noteID-1] = Note{}
	returnMsg += fmt.Sprintf("The record for note with ID '%v' has been successfully deleted.", noteID)
	return returnMsg
}

// Find Note Function
func findNote(inputPattern string) (bool, string) {
	result := false
	returnMsg := ""

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

	sqlQuery := fmt.Sprintf(`SELECT * FROM notes WHERE noteText LIKE %v;`, inputPattern)
	queryRows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when trying to find note text matching the given pattern.\n"
	}
	// for _, aNote := range Notes {
	// 	if strings.Contains(aNote.Text, inputPattern) {
	// 		result = true
	// 		returnMsg += fmt.Sprintf("The text \"%v\" was found in note '%v'.\n\nDetails:\n%v\n", input, aNote.Name, aNote)
	// 		return result, returnMsg
	// 	}
	// }
	returnMsg += fmt.Sprintf("At least one match was successfully found for that pattern. Result:\n%v\n", queryRows)
	return result, returnMsg
}

// Analyse Note Function
func analyseNote(inputPattern string, noteID int) (int, string) {
	result := 0
	// j := 0
	//pattern := len(input)
	// noteText := len(note.Text)
	returnMsg := ""

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

	// for result < noteText-pattern+1 {
	// 	for j = 0; j < pattern-1; j++ {
	// 		if input[result+j] != note.Text[j] { //match letter from substring
	// 			break // failed match, continue searching
	// 		}
	// 	}

	// 	if j == pattern-1 {
	// 		if result == 1 {
	// 			returnMsg += fmt.Sprintln("The analysis returned %v instance of \"%v\" in the note '%v'\n", strconv.Itoa(result), input, note.Name)
	// 		} else {
	// 			returnMsg += fmt.Sprintln("The analysis returned %v instances of \"%v\" in the note '%v'\n", strconv.Itoa(result), input, note.Name)
	// 		}
	// 		return result, returnMsg
	// 	} else if j == 0 {
	// 		result++
	// 	} else {
	// 		result = result + j
	// 	}
	// }

	sqlQuery := fmt.Sprintf(`SELECT regexp_matches("noteText", '%s') FROM notes WHERE "noteID" = %d;`, inputPattern, noteID)
	queryRows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
		returnMsg += "An error occurred when trying to find text matching the given pattern.\n"
	}

	returnMsg += fmt.Sprintln("The analysis returned %d instances of \"%v\" in the text.\nResult: %v\n", result, inputPattern, queryRows)
	return -1, returnMsg //not found so return no position
}

// Select Pattern Function
// func selectPattern() string {
// 	pattern := ""
// 	if optionSelect {
// 		var inputNum int
// 		println("Please select a pattern option:\n\n1. A sentence with a given prefix and/or suffix\n2. A phone number with a given area code and optionally a consecutive sequence of numbers that are part of that number\n3. An email address on a domain that is only partially provided\n4. Text that contains at least three of the following case-insensitive words: meeting, minutes, agenda, action, attendees, apologies\n5. A word in all capitals of three characters or more\nOr enter 'r' to return.")
// 		fmt.Scanln(inputNum)
// 		switch inputNum {
// case "1":
// 	println("\n\nPlease enter a string that matches the pattern: A sentence with a given prefix and/or suffix")
// 	fmt.Scanln(input)
// 	pattern = input
// 	return pattern
// case "2":
// 	pattern := `[0-9\W]`
// 	println("\n\nPlease enter a string that matches the pattern: A phone number with a given area code and optionally a consecutive sequence of numbers that are part of that number")
// 	fmt.Scanln(input)

// 	// Check for erroneous value
// 	switch validatePattern(pattern, input).isValid {
// 	case true:
// 		pattern = input
// 		return pattern
// 	default:
// 		fmt.print(validatePattern(pattern, input).returnMsg)
// 	continue
// case "3":
// 	pattern := `@{1}`
// 	println("\n\nPlease enter a string that matches the pattern: An email address on a domain that is only partially provided")
// 	fmt.Scanln(input)
// 	_, err := regexp.MatchString(`@{1}`, input)
// 	pattern = input
// 	return pattern
// case "4":
// 	pattern := `meeting|minutes|agenda|action|attendees|apologies{3,}`
// 	println("\n\nPlease enter a string that matches the pattern: Text that contains at least three of the following case-insensitive words: meeting, minutes, agenda, action, attendees, apologies")
// 	fmt.Scanln(input)
// 	_, err := regexp.MatchString(`meeting|minutes|agenda|action|attendees|apologies{3,}`, input)
// 	pattern = input
// 	return pattern
// case "5":
// 	pattern := `[A-Z]{3,}`
// 	println("\n\nPlease enter a string that matches the pattern: A word in all capitals of three characters or more")
// 	fmt.Scanln(input)
// 	_, err := regexp.MatchString(`[A-Z]{3,}`, input)
// 	pattern = input
// 	return pattern
// 		case 1, 2, 3, 4, 5:
// 			i := 0
// 			for k, v := range Patterns {
// 				i++
// 				if i == inputNum {
// 					var inputStr string
// 					pattern := Patterns[k]
// 					patternDesc := Patterns[v]
// 					fmt.Printf("\n\nPlease enter a string that matches the pattern: %v", patternDesc)
// 					fmt.Scanln(inputStr)

// 					// Check for erroneous value
// 					switch isValid, returnMsg := validatePattern(pattern, inputStr); {
// 					case isValid:
// 						pattern = inputStr
// 						return pattern
// 					default:
// 						fmt.Println(returnMsg)
// 						continue
// 					}
// 				}
// 			}
// 		default:
// 			return ""
// 		}
// 	}
// 	return pattern
// }

// Select Option Function
// func selectOption() {
// 	if optionSelect {
// 		input := ""
// 		println("Please select an option:\n\n1. Users\n2. Notes")
// 		fmt.Scanln(input)
// 		switch input {
// 		case "1":
// 			println("\n\nPlease select an option:\n\n1.Create User\n2.Read User\n3.Update User\n4.Delete User\nOr enter 'r' to return.")
// 			fmt.Scanln(input)
// 			switch input {
// 			case "1":
// 				//createUser()
// 			case "2":
// 				//readUser()
// 			case "3":
// 				//updateUser()
// 			case "4":
// 				//deleteUser()
// 			default:
// 				return
// 			}
// 		default:
// 			return
// 		}
// 	}
// }

// --- Main ---//
func main() {
	go StartServer()
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())
	//selectOption()
}
