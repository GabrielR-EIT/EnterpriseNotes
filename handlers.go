package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//go:embed templates
var tmplEmbed embed.FS

// Middleware to connect the database for each request that uses this
// middleware.
func connectDatabase(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
	}
}

// Return all users as JSON
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Users)
}

// Return all notes as JSON
func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Notes)
}

// func ReadNotes(ctx *gin.Context, db *sqlx.DB, req *http.Request) string {
// 	returnMsg := ""
// 	err := req.ParseForm()
// 	if err != nil {
// 		log.Println("An error occurred when parsing form data.\nGot\n", err)
// 	}
// 	if req.FormValue("selectNote") == "all" {
// 		for _, note := range Notes {
// 			returnMsg += readNote(db, note.ID)
// 		}
// 	} else {
// 		selectedNote, err := strconv.Atoi(req.FormValue("selectNote"))
// 		if err != nil {
// 			log.Println("An error occurred when parsing form data.\nGot\n", err)
// 		}
// 		returnMsg += readNote(db, selectedNote)
// 	}
// 	ctx.HTML(http.StatusOK, "views/users.html", gin.H{"results": returnMsg})
// 	return returnMsg
// }

func handlerReadNotes(ctx *gin.Context) {
	results := ""
	//db := ctx.Value("database").(*sqlx.DB)
	ctx.Request.ParseForm()
	noteID := ctx.Request.FormValue("selectNote")
	fmt.Println(noteID)
	if noteID != "" {
		noteID := ctx.Request.Context().Value("selectNote")
		if noteID == "all" {
			noteNum := 0
			for _, note := range Notes {
				noteNum++
				//results += readNote(db, note.ID)
				//fmt.Println(results)
				results += fmt.Sprintf("Note %d: note%v\n\n", noteNum, note)
			}
		} else {
			//results += readNote(db, noteID.(int))
			//fmt.Println(results)
			for _, note := range Notes {
				if note.ID == noteID {
					//results += readNote(db, note.ID)
					//fmt.Println(results)
					results += fmt.Sprint(note)
				}
			}
		}
	}
	ctx.HTML(http.StatusOK, "views/notes.html", gin.H{"results": results})
}
