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

// // Return all users as JSON
// func getUsers(ctx *gin.Context) {
// 	ctx.IndentedJSON(http.StatusOK, Users)
// }

// // Return all notes as JSON
// func getNotes(ctx *gin.Context) {
// 	ctx.IndentedJSON(http.StatusOK, Notes)
// }

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
