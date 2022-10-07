package main

import (
	_ "database/sql"
	"fmt"
	_ "log"
	"strconv"
	"strings"

	//_ "github.com/YZakizon/geeksbeginner/golang-gin-template/src/web"

	_ "github.com/gin-gonic/gin"
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
	ID              int      `json:"note_id"`
	Name            string   `json:"name"`
	Text            string   `json:"text"`
	Completion_Time string   `json:"completion_date"`
	Status          string   `json:"status"`
	Delegation      string   `json:"delegation"`
	Shared_Users    []string `json:"shared_users"`
}

type Association struct {
	ID         int    `json:"association_id"`
	userID     int    `json:"user_id"`
	noteID     int    `json:"note_id"`
	permission string `json:"permission"`
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
var optionSelect = true

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
func createUser() string {
	outputMsg := ""

	user_name := ""
	var user_read_setting, user_write_setting bool
	fmt.Scanln(&user_name)
	fmt.Scanln(&user_read_setting)
	fmt.Scanln(&user_write_setting)

	// Create struct for new note
	new_user := User{
		Name:          user_name,
		Read_Setting:  user_read_setting,
		Write_Setting: user_write_setting,
	}

	//sqlQuery := fmt.Sprintf(`INSERT INTO users VALUE (%v, %v, %v);`, new_user.Name, new_user.Read_Setting, new_user.Write_Setting)
	// _, err := db.Exec(sqlQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("An error occurred when trying to create the user.")
	// }

	// Add new note struct to Notes slice
	Users = append(Users, new_user)
	outputMsg = fmt.Sprintf("A new user has been successfully added.\nDetails:\n%v\nThere are now %v users.\n", new_user, strconv.Itoa(len(Users)))
	return outputMsg
}

// Read User Function
func readUser(user *User) string {
	outputMsg := fmt.Sprintf("ID: %v\nName: %v\nRead Setting: %v\nWrite Setting: %v\n", user.ID, user.Name, user.Read_Setting, user.Write_Setting)
	return outputMsg
}

// Update User Function
func updateUser(user *User) string {
	input, outputMsg := "", ""
	println("Details:\n", user, "\nDo you wish to edit:\n1. Name\n2. Read Setting\n3. Write Setting\nOr type N to cancel.")
	fmt.Scanln(input)
	switch input {
	case "1":
		fmt.Scanln(input)
		user.Name = input
		//sqlQuery := fmt.Sprintf(`UPDATE users SET name = %v WHERE ID = %v;`, input, user.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the user.")
		// }
		outputMsg = fmt.Sprintf("The name for this user has been changed to '%v'", input)
		return outputMsg
	case "2":
		fmt.Scanln(input)
		//user.Read_Setting = strconv.ParseBool(input)
		//sqlQuery := fmt.Sprintf(`UPDATE users SET read_setting = %v WHERE ID = %v;`, input, user.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the user.")
		// }
		outputMsg = fmt.Sprintf("The read setting for this user has been changed to '%v'", input)
		return outputMsg
	case "3":
		fmt.Scanln(input)
		//user.Write_Setting = strconv.ParseBool(input)
		//sqlQuery := fmt.Sprintf(`UPDATE users SET write_setting = %v WHERE ID = %v;`, input, user.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the user.")
		// }
		outputMsg = fmt.Sprintf("The write setting for this user has been changed to '%v'", input)
		return outputMsg
	default:
		return ""
	}
}

// Delete User Function
func deleteUser(user User) string {
	outputMsg := ""
	//sqlQuery := fmt.Sprintf(`DELETE FROM users WHERE ID = %v;`, user.ID)
	// _, err := db.Exec(sqlQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("An error occurred when trying to delete the user.")
	// }

	//set the values of the user to null
	user = User{}
	outputMsg = fmt.Sprintf("The information for user '%v' has been deleted.", user.Name)
	return outputMsg
}

// Create Note Function
func createNote() string {
	outputMsg := ""
	note_name, note_text, note_time, note_status, note_delegation, note_users := "", "", "", "", "", []string{}
	fmt.Scanln(&note_name)
	fmt.Scanln(&note_text)
	fmt.Scanln(&note_time)
	fmt.Scanln(&note_status)
	fmt.Scanln(&note_delegation)
	fmt.Scanln(&note_users)

	// Create struct for new note
	new_note := Note{
		Name:            note_name,
		Text:            note_text,
		Completion_Time: note_time,
		Status:          note_status,
		Delegation:      note_delegation,
		Shared_Users:    note_users,
	}

	//sqlQuery := fmt.Sprintf(`INSERT INTO notes VALUE (%v, %v, %v, %v, %v, %v);`, new_note.Name, new_note.Text, new_note.Completion_Time, new_note.Status, new_note.Delegation, new_note.Shared_Users)
	// _, err := db.Exec(sqlQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("An error occurred when trying to create the note.")
	// }

	// Add new note struct to Notes slice
	Notes = append(Notes, new_note)
	outputMsg = fmt.Sprintf("Your new note has been successfully added.\nDetails:\n%v\nThere are now %v notes.", new_note, strconv.Itoa(len(Notes)))
	return outputMsg

	// Create an association between the user and note
}

// Read Note Function
func readNote(note *Note) string {
	// println(note)
	outputMsg := fmt.Sprintf("ID: %v\nName: %v\nText: '%v'\nCompletion Time: %v\nStatus: %v\nDelegation: %v\nShared Users: %v\n", note.ID, note.Name, note.Text, note.Completion_Time, note.Status, note.Delegation, note.Shared_Users)
	return outputMsg
}

// Update Note Function
func updateNote(note *Note) string {
	input, outputMsg := "", ""
	println("Details:\n", note, "\nDo you wish to edit:\n1. Name\n2. Text\n3. Completion Time\n4. Status\n5. Delegation\n6. Shared Users\nOr type N to cancel.")
	fmt.Scanln(input)
	switch input {
	case "1":
		fmt.Scanln(input)
		note.Name = input
		//sqlQuery := fmt.Sprintf(`UPDATE notes SET name = %v WHERE ID = %v;`, input, note.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the note.")
		// }
		outputMsg = fmt.Sprintf("The name for this note has been changed to '%v'\n", input)
		return outputMsg
	case "2":
		fmt.Scanln(input)
		note.Text = input
		//sqlQuery := fmt.Sprintf(`UPDATE notes SET text = %v WHERE ID = %v;`, input, note.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the note.")
		// }
		outputMsg = fmt.Sprintf("The text for this note has been changed to '%v'\n", input)
		return outputMsg
	case "3":
		fmt.Scanln(input)
		note.Completion_Time = input
		//sqlQuery := fmt.Sprintf(`UPDATE notes SET completion_time = %v WHERE ID = %v;`, input, note.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the note.")
		// }
		outputMsg = fmt.Sprintf("The completion time for this note has been changed to '%v'\n", input)
		return outputMsg
	case "4":
		fmt.Scanln(input)
		note.Status = input
		//sqlQuery := fmt.Sprintf(`UPDATE notes SET status = %v WHERE ID = %v;`, input, note.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the note.")
		// }
		outputMsg = fmt.Sprintf("The status for this note has been changed to '%v'\n", input)
		return outputMsg
	case "5":
		fmt.Scanln(input)
		note.Delegation = input
		//sqlQuery := fmt.Sprintf(`UPDATE notes SET delegation = %v WHERE ID = %v;`, input, note.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the note.")
		// }
		outputMsg = fmt.Sprintf("The delegation for this note has been changed to '%v'\n", input)
		return outputMsg
	case "6":
		fmt.Scanln(input)
		// note.Shared_Users = input
		//sqlQuery := fmt.Sprintf(`UPDATE notes SET shared_users = %v WHERE ID = %v;`, input, note.ID)
		// _, err := db.Exec(sqlQuery)
		// if err != nil {
		// 	log.Fatal(err)
		// 	fmt.Println("An error occurred when trying to update the note.")
		// }
		outputMsg = fmt.Sprintf("The shared users for this note have been changed to '%v'\n", input)
		return outputMsg
	default:
		return ""
	}
}

// Delete Note Function
func deleteNote(note Note) string {
	outputMsg := ""
	//sqlQuery := fmt.Sprintf(`DELETE FROM notes WHERE ID = %v;`, note.ID)
	// _, err := db.Exec(sqlQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("An error occurred when trying to delete the note.")
	// }

	//set the values of the note to null
	note = Note{}
	outputMsg = fmt.Sprintf("The information for note '%v' has been deleted.", note.Name)
	return outputMsg
}

// Find Note Function
func findNote(input string) (bool, string) {
	result := false
	outputMsg := ""
	for _, aNote := range Notes {
		if strings.Contains(aNote.Text, input) {
			result = true
			outputMsg = fmt.Sprintf("The text \"%v\" was found in note '%v'.\n\nDetails:\n%v\n", input, aNote.Name, aNote)
			return result, outputMsg
		}
	}
	return result, ""
}

// Analyse Note Function
func analyseNote(input string, note *Note) (int, string) {
	result := 0
	j := 0
	pattern := len(input)
	noteText := len(note.Text)
	outputMsg := ""
	for result < noteText-pattern+1 {
		for j = 0; j < pattern-1; j++ {
			if input[result+j] != note.Text[j] { //match letter from substring
				break // failed match, continue searching
			}
		}

		if j == pattern-1 {
			if result == 1 {
				outputMsg = fmt.Sprintln("The analysis returned %v instance of \"%v\" in the note '%v'\n", strconv.Itoa(result), input, note.Name)
			} else {
				outputMsg = fmt.Sprintln("The analysis returned %v instances of \"%v\" in the note '%v'\n", strconv.Itoa(result), input, note.Name)
			}
			return result, outputMsg
		} else if j == 0 {
			result++
		} else {
			result = result + j
		}
	}
	outputMsg = fmt.Sprintln("The analysis returned zero instances of \"%v\" in the note '%v'\n", input, note.Name)
	return -1, outputMsg //not found so return no position
}

// Select Pattern Function
func selectPattern() string {
	pattern := ""
	if optionSelect {
		var inputNum int
		println("Please select a pattern option:\n\n1. A sentence with a given prefix and/or suffix\n2. A phone number with a given area code and optionally a consecutive sequence of numbers that are part of that number\n3. An email address on a domain that is only partially provided\n4. Text that contains at least three of the following case-insensitive words: meeting, minutes, agenda, action, attendees, apologies\n5. A word in all capitals of three characters or more\nOr enter 'r' to return.")
		fmt.Scanln(inputNum)
		switch inputNum {
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
		case 1, 2, 3, 4, 5:
			i := 0
			for k, v := range Patterns {
				i++
				if i == inputNum {
					var inputStr string
					pattern := Patterns[k]
					patternDesc := Patterns[v]
					fmt.Printf("\n\nPlease enter a string that matches the pattern: %v", patternDesc)
					fmt.Scanln(inputStr)

					// Check for erroneous value
					switch isValid, returnMsg := validatePattern(pattern, inputStr); {
					case isValid:
						pattern = inputStr
						return pattern
					default:
						fmt.Println(returnMsg)
						continue
					}
				}
			}
		default:
			return ""
		}
	}
	return pattern
}

// Select Option Function
func selectOption() {
	if optionSelect {
		input := ""
		println("Please select an option:\n\n1. Users\n2. Notes")
		fmt.Scanln(input)
		switch input {
		case "1":
			println("\n\nPlease select an option:\n\n1.Create User\n2.Read User\n3.Update User\n4.Delete User\nOr enter 'r' to return.")
			fmt.Scanln(input)
			switch input {
			case "1":
				createUser()
			case "2":
				//readUser()
			case "3":
				//updateUser()
			case "4":
				//deleteUser()
			default:
				return
			}
		default:
			return
		}
	}
}

// --- Main ---//
func main() {
	go fmt.Println("Test")
	fmt.Print(CreateDB())
	fmt.Print(CreateTables())
	fmt.Print(PopulateTables())
	selectOption()
}
