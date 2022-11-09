package main

import (
	"embed"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//go:embed templates
var tmplEmbed embed.FS

// Middleware to connect the database for each request that uses this
// middleware.
func connectDatabase(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("database", db)
	}
}

// Load Data Handler
func handlerLoadData(ctx *gin.Context, results string, location string) {
	db := ctx.Value("database").(*sqlx.DB)

	foundUsers := []User{}
	sqlQuery := `SELECT * FROM users;`
	err := db.Select(&foundUsers, sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to return data from the users table.\nGot %s\n", err)
	}

	foundNotes := []Note{}
	sqlQuery = `SELECT * FROM notes;`
	err = db.Select(&foundNotes, sqlQuery)
	if err != nil {
		log.Printf("An error occurred when trying to return data from the notes table.\nGot %s\n", err)
	}

	if location == "views/users.html" {
		ctx.HTML(http.StatusOK, "views/users.html", gin.H{"users": foundUsers, "results": results})
	}
	if location == "views/notes.html" {
		ctx.HTML(http.StatusOK, "views/notes.html", gin.H{"notes": foundNotes, "statuses": Statuses, "users": foundUsers, "results": results})
	}
	ctx.Redirect(http.StatusFound, location)
}

// --- CRUD Handlers --- //
// Create User Handler
func handlerCreateUser(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	inputName := ctx.PostForm("inputName")
	inputReadSetting := (ctx.PostForm("inputReadSetting") == "on")
	inputWriteSetting := (ctx.PostForm("inputWriteSetting") == "on")
	results += createUser(db, inputName, inputReadSetting, inputWriteSetting)

	// Update the form data
	handlerLoadData(ctx, results, "views/users.html")
}

// Read Users Handler
func handlerReadUsers(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	userID := ctx.Request.FormValue("selectUser")
	// Check if a user has been selected
	if userID != "" {
		// Read all users
		if userID == "all" {
			for _, user := range Users {
				results += readUser(db, user.ID)
			}
		} else {
			// Read a single user
			userID, err := strconv.Atoi(userID)
			if err != nil {
				log.Println("An error occurred when parsing form data.\nGot\n", err)
			}
			results += readUser(db, userID)
		}
	}

	// Update the form data
	handlerLoadData(ctx, results, "views/users.html")
}

// Update User Handler
func handlerUpdateUser(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	inputID, err := strconv.Atoi(ctx.Request.FormValue("selectUser"))
	if err != nil {
		log.Println("An error occurred when parsing form data.\nGot\n", err)
	}
	inputName := ctx.PostForm("inputName")
	inputReadSetting := (ctx.PostForm("inputReadSetting") == "on")
	inputWriteSetting := (ctx.PostForm("inputWriteSetting") == "on")

	// Check if Delete Option is Selected
	switch inputDelete := (ctx.PostForm("inputDelete") == "on"); inputDelete {
	case true:
		// Delete the User
		results += deleteUser(db, inputID)
	default:
		// Update the User
		results += updateUser(db, inputID, inputName, inputReadSetting, inputWriteSetting)
	}

	// Update the form data
	handlerLoadData(ctx, results, "views/users.html")
}

// Create Note Handler
func handlerCreateNote(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)
	var err interface{}

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	inputName := ctx.PostForm("inputName")
	inputText := ctx.PostForm("inputText")
	inputStatus := ctx.PostForm("inputStatus")
	inputDelegation, err := strconv.Atoi(ctx.PostForm("inputDelegation"))
	if err != nil {
		log.Println("An error occurred when parsing form data.\nGot\n", err)
	}
	inputSharedUsers, err := ctx.GetPostFormArray("inputSharedUsers")
	if err == false {
		log.Println("An error occurred when parsing form data.\nGot\n", err)
	}
	// Add Formatting to the Shared Users Input
	sharedUsersString := "[" + strings.Join(inputSharedUsers, `,`) + `]`
	results += createNote(db, inputName, inputText, inputStatus, inputDelegation, sharedUsersString)

	// Update the form data
	handlerLoadData(ctx, results, "views/notes.html")
}

// Read Notes Handler
func handlerReadNotes(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected Note
	ctx.Request.ParseForm()
	noteID := ctx.Request.FormValue("selectNote")

	// Check if a note has been selected
	if noteID != "" {
		// Read all notes
		if noteID == "all" {
			for _, note := range Notes {
				results += readNote(db, note.ID)
			}
		} else {
			// Read a single note
			noteID, err := strconv.Atoi(noteID)
			if err != nil {
				log.Println("An error occurred when parsing form data.\nGot\n", err)
			}
			results += readNote(db, noteID)
		}
	}
	// Update the form data
	handlerLoadData(ctx, results, "views/notes.html")
}

// Update Note Handler
func handlerUpdateNote(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)
	var err interface{}

	// Parse the HTTP Response for Selected Note
	ctx.Request.ParseForm()
	inputID, err := strconv.Atoi(ctx.Request.FormValue("selectNote"))
	if err != nil {
		log.Println("An error occurred when parsing form data.\nGot\n", err)
	}
	inputName := ctx.PostForm("inputName")
	inputText := ctx.PostForm("inputText")
	inputStatus := ctx.PostForm("inputStatus")
	inputDelegation := ctx.GetInt(ctx.PostForm("inputDelegation"))
	inputSharedUsers, err := ctx.GetPostFormArray("inputSharedUsers")
	if err == false {
		log.Println("An error occurred when parsing form data.\nGot\n", err)
	}
	// Add Formatting to the Shared Users Input
	sharedUsersString := "[" + strings.Join(inputSharedUsers, `,`) + `]`

	// Check if Delete Option is Selected
	switch inputDelete := (ctx.PostForm("inputDelete") == "on"); inputDelete {
	case true:
		// Delete the Note
		results += deleteNote(db, inputID)
	default:
		// Update the Note
		results += updateNote(db, inputID, inputName, inputText, inputStatus, inputDelegation, sharedUsersString)
	}

	// Update the form data
	handlerLoadData(ctx, results, "views/notes.html")
}
