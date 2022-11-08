package main

// Create Structs
type User struct {
	ID            int    `json:"user_id" db:"userid"`
	Name          string `json:"name" db:"username"`
	Read_Setting  bool   `json:"read_setting" db:"userreadsetting"`
	Write_Setting bool   `json:"write_setting" db:"userwritesetting"`
}

type Note struct {
	ID              int    `json:"note_id" db:"noteid"`
	Name            string `json:"name" db:"notename"`
	Text            string `json:"text" db:"notetext"`
	Completion_Time string `json:"completion_date" db:"notecompletiontime"`
	Status          string `json:"status" db:"notestatus"`
	Delegation      int    `json:"delegation" db:"notedelegation"`
	Shared_Users    string `json:"shared_users" db:"notesharedusers"`
}

type Association struct {
	ID         int    `json:"association_id" db:"associationid"`
	UserID     int    `json:"user_id" db:"userid"`
	NoteID     int    `json:"note_id" db:"noteid"`
	Permission string `json:"permission" db:"associationpermission"`
}

// Create Struct for Test Data
type Data struct {
	Users        []User        `json:"users"`
	Notes        []Note        `json:"notes"`
	Associations []Association `json:"associations"`
}

// Create Slices
var Users []User
var Notes []Note
var Associations []Association
var Statuses = [5]string{
	"none",
	"in progress",
	"completed",
	"cancelled",
	"delegated",
}
