package main

import (
	"embed"
	"log"
	"net/http"
	"strconv"

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

// Create User Handler
func handlerCreateUser(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	inputName := ctx.PostForm("inputName")
	inputReadSetting := ctx.GetBool(ctx.PostForm("inputReadSetting"))
	inputWriteSetting := ctx.GetBool(ctx.PostForm("inputWriteSetting"))
	results += createUser(db, inputName, inputReadSetting, inputWriteSetting)

	// Update the form data
	ctx.HTML(http.StatusOK, "views/users.html", gin.H{"users": Users, "results": results})
	ctx.Redirect(http.StatusFound, "views/users.html")
}

// Create Note Handler
func handlerCreateNote(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	inputName := ctx.PostForm("inputName")
	inputText := ctx.PostForm("inputText")
	inputStatus := ctx.PostForm("inputStatus")
	inputDelegation := ctx.GetInt(ctx.PostForm("inputDelegation"))
	inputSharedUsers := ctx.PostForm("inputSharedUsers")
	results += createNote(db, inputName, inputText, inputStatus, inputDelegation, inputSharedUsers)

	// Update the form data
	ctx.HTML(http.StatusOK, "views/notes.html", gin.H{"notes": Notes, "statuses": Statuses, "users": Users, "results": results})
	ctx.Redirect(http.StatusFound, "views/notes.html")
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
	ctx.HTML(http.StatusOK, "views/users.html", gin.H{"users": Users, "results": results})
	ctx.Redirect(http.StatusFound, "views/users.html")
}

// Update User Handler
func handlerUpdateUser(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	inputID := ctx.GetInt(ctx.Request.FormValue("selectUser"))
	inputName := ctx.Request.FormValue("inputName")
	inputReadSetting := ctx.GetBool(ctx.Request.FormValue("inputReadSetting"))
	inputWriteSetting := ctx.GetBool(ctx.Request.FormValue("inputWriteSetting"))

	// Check if Delete Option is Selected
	switch inputDelete := ctx.GetBool(ctx.Request.FormValue("inputDelete")); inputDelete {
	case true:
		results += deleteUser(db, inputID)
	default:
		results += updateUser(db, inputID, inputName, inputReadSetting, inputWriteSetting)
	}

	// Update the form data
	ctx.HTML(http.StatusOK, "views/users.html", gin.H{"users": Users, "results": results})
	ctx.Redirect(http.StatusFound, "views/users.html")
}

// Update Note Handler
func handlerUpdateNote(ctx *gin.Context) {
	results := ""
	db := ctx.Value("database").(*sqlx.DB)

	// Parse the HTTP Response for Selected User
	ctx.Request.ParseForm()
	inputID := ctx.GetInt(ctx.Request.FormValue("selectNote"))
	inputName := ctx.Request.FormValue("inputName")
	inputText := ctx.Request.FormValue("inputText")
	inputStatus := ctx.Request.FormValue("inputStatus")
	inputDelegation := ctx.GetInt(ctx.Request.FormValue("inputDelegation"))
	inputSharedUsers := ctx.Request.FormValue("inputSharedUsers")

	// Check if Delete Option is Selected
	switch inputDelete := ctx.GetBool(ctx.Request.FormValue("inputDelete")); inputDelete {
	case true:
		results += deleteNote(db, inputID)
	default:
		results += updateNote(db, inputID, inputName, inputText, inputStatus, inputDelegation, inputSharedUsers)
	}

	// Update the form data
	ctx.HTML(http.StatusOK, "views/notes.html", gin.H{"notes": Notes, "statuses": Statuses, "users": Users, "results": results})
	ctx.Redirect(http.StatusFound, "views/notes.html")
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
	ctx.HTML(http.StatusOK, "views/notes.html", gin.H{"notes": Notes, "statuses": Statuses, "users": Users, "results": results})
	ctx.Redirect(http.StatusFound, "views/notes.html")
}
