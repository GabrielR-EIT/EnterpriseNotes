package main

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
	Delegation      int    `json:"delegation"`
	Shared_Users    string `json:"shared_users"`
}

type Association struct {
	ID         int    `json:"association_id"`
	UserID     int    `json:"user_id"`
	NoteID     int    `json:"note_id"`
	Permission string `json:"permission"`
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
