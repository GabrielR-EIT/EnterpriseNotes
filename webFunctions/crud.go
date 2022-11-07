package webFunctions

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

// --- User CRUD HTTP Functions --- //
// Create User HTTP Function
func CreateUser(db *sqlx.DB, newUser interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// //var newUser NewUser
		// var ctx = c.Request.Context()

		// if e := c.ShouldBindJSON(&newUser); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }

		// if _, e := db.ExecContext(ctx, `INSERT INTO Users(userName, userReadSetting, userWriteSetting) VALUES(%s,%t,%t)`, newUser.userName, newUser.userReadSetting, newUser.userWriteSetting); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return

		// }

		// //var newUser User
		// var row = db.QueryRow(`SELECT * FROM Users WHERE userID=%d`, newUser.userID)

		// if e := row.Scan(fmt.Sprint(`%v.userName, %v.userReadSetting, %v.userWriteSetting`, &newUser)); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var res = NewHTTPResponse(http.StatusCreated, newUser)

		// c.Writer.Header().Add("Location", fmt.Sprintf("/Users/%s", newUser.userID))
		// c.JSON(http.StatusCreated, res)
	}
}

// Read User HTTP Function
func ReadUser(db *sqlx.DB, userID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var binding guidBinding
		// var ctx = c.Request.Context()
		// if e := c.ShouldBindUri(&binding); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var row = db.QueryRowContext(ctx, `SELECT * FROM Users WHERE userID = %d`, userID)
		// var user interface{}

		// // if e := row.Scan(fmt.Sprintf(`%v.userID, %v.userName, %v.userReadSetting, %v.userWriteSetting`, &user)); e != nil {
		// // 	if e != sql.ErrNoRows {
		// // 		var res = NewHTTPResponse(http.StatusNotFound, e)
		// // 		c.JSON(http.StatusNotFound, res)
		// // 		return
		// // 	}

		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var res = NewHTTPResponse(http.StatusOK, user)
		// c.JSON(http.StatusOK, res)
	}
}

// Read Users HTTP Function
func ReadUsers(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rows *sqlx.Rows
		var e error
		if rows, e = db.Queryx(`SELECT * FROM Users`); e != nil {
			var res = NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		defer rows.Close()
		var users []interface{}
		for rows.Next() {
			var user interface{}

			if e := rows.Scan(fmt.Sprintf(`%d.userID, %s.userName, %t.userReadSetting, %t.userWriteSetting`, &user)); e != nil {
				var res = NewHTTPResponse(http.StatusInternalServerError, e)
				c.JSON(http.StatusInternalServerError, res)
				return
			}

			users = append(users, user)
		}

		if len(users) == 0 {
			var res = NewHTTPResponse(http.StatusNotFound, sql.ErrNoRows)
			c.JSON(http.StatusNotFound, res)
			return
		}

		var res = NewHTTPResponse(http.StatusOK, users)
		c.JSON(http.StatusOK, res)
	}
}

// Update User HTTP Function
func UpdateUser(db *sqlx.DB, user interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var binding guidBinding
		// var updatedUser UpdatedUser
		// var ctx = c.Request.Context()

		// if e := c.ShouldBindUri(&binding); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }
		// if e := c.ShouldBindJSON(&updatedUser); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }

		// var row = db.QueryRowContext(ctx, `SELECT * FROM Users WHERE userID = %d`, user.userID)

		// if e := row.Scan(fmt.Sprintf(`%v.userName, %v.userReadSetting, %v.userWriteSetting`, &user)); e != nil {
		// 	if e == sql.ErrNoRows {
		// 		var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 		c.JSON(http.StatusNotFound, res)
		// 		return
		// 	}

		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var option = copier.Option{
		// 	IgnoreEmpty: true,
		// 	DeepCopy:    true,
		// }

		// if e := copier.CopyWithOption(&user, &updatedUser, option); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// // var updatedRow = db.QueryRowContext(ctx, `SELECT * FROM Users WHERE userID = %d`, user.userID)
		// // var newUser interface{}

		// // if e := updatedRow.Scan(fmt.Sprintf(`%v.userName, %v.userReadSetting, %v.userWriteSetting`, &newUser)); e != nil {
		// // 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// // 	c.JSON(http.StatusInternalServerError, res)
		// // 	return
		// // }

		// var res = NewHTTPResponse(http.StatusOK, user)
		// c.JSON(http.StatusOK, res)
	}
}

// Delete User HTTP Function
func DeleteUser(db *sqlx.DB, userID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var binding guidBinding
		// var ctx = c.Request.Context()
		// if e := c.ShouldBindUri(&binding); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }

		// var result sql.Result
		// var e error
		// if result, e = db.ExecContext(ctx, `DELETE FROM Users WHERE userID = %v`, userID); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// if nUsers, _ := result.RowsAffected(); nUsers == 0 {
		// 	var res = NewHTTPResponse(http.StatusNotFound, sql.ErrNoRows)
		// 	c.JSON(http.StatusNotFound, res)
		// 	return
		// }

		// c.JSON(http.StatusNoContent, nil)
	}
}

// --- Note CRUD HTTP Functions --- //
// Create Note HTTP Function
func CreateNote(db *sqlx.DB, newNote interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// //var newNote NewNote
		// var ctx = c.Request.Context()

		// if e := c.ShouldBindJSON(&newNote); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }

		// // var createdAt = time.Now().Format(time.RFC3339)
		// if _, e := db.ExecContext(ctx, `INSERT INTO Notes(noteName, noteText, noteCompletionTime, noteStatus, noteDelegation, noteSharedUsers) VALUES(%s, %s, NOW(), %s, %d, %v)`, newNote.noteName, newNote.noteText, newNote.noteStatus, newNote.noteDelegation, newNote.noteSharedUsers); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return

		// }

		// //var newNote Note
		// var row = db.QueryRow(`SELECT * FROM Notes WHERE noteID=%d`, newNote.noteID)

		// if e := row.Scan(fmt.Sprint(`%v.noteName, %v.noteText, %v.noteCompletionTime, %v.noteStatus, %v.noteDelegation, %v.noteSharedUsers`, &newNote)); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var res = NewHTTPResponse(http.StatusCreated, newNote)

		// c.Writer.Header().Add("Location", fmt.Sprintf("/Notes/%s", newNote.noteID))
		// c.JSON(http.StatusCreated, res)
	}
}

// Read Note HTTP Function
func ReadNote(db *sqlx.DB, noteID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var binding guidBinding
		// var ctx = c.Request.Context()
		// if e := c.ShouldBindUri(&binding); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var row = db.QueryRowContext(ctx, `SELECT * FROM Notes WHERE noteID = %d`, noteID)
		// var note interface{}

		// // if e := row.Scan(fmt.Sprintf(`%v.noteName, %v.noteText, %v.noteCompletionTime, %v.noteStatus, %v.noteDelegation, %v.noteSharedUsers`, &note)); e != nil {
		// // 	if e != sql.ErrNoRows {
		// // 		var res = NewHTTPResponse(http.StatusNotFound, e)
		// // 		c.JSON(http.StatusNotFound, res)
		// // 		return
		// // 	}

		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var res = NewHTTPResponse(http.StatusOK, note)
		// c.JSON(http.StatusOK, res)
	}
}

// Read Notes HTTP Function
func ReadNotes(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rows *sqlx.Rows
		var e error
		if rows, e = db.Queryx(`SELECT * FROM Notes`); e != nil {
			var res = NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		defer rows.Close()
		var notes []interface{}
		for rows.Next() {
			var note interface{}

			if e := rows.Scan(fmt.Sprintf(`%d.noteID, %s.noteName, %s.noteText, %s.noteCompletionTime, %s.noteStatus, %d.noteDelegation, %v.noteSharedUsers`, &note)); e != nil {
				var res = NewHTTPResponse(http.StatusInternalServerError, e)
				c.JSON(http.StatusInternalServerError, res)
				return
			}

			notes = append(notes, note)
		}

		if len(notes) == 0 {
			var res = NewHTTPResponse(http.StatusNotFound, sql.ErrNoRows)
			c.JSON(http.StatusNotFound, res)
			return
		}

		var res = NewHTTPResponse(http.StatusOK, notes)
		c.JSON(http.StatusOK, res)
	}
}

// Update Note HTTP Function
func UpdateNote(db *sqlx.DB, note interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var binding guidBinding
		// var updatedNote UpdatedNote
		// var ctx = c.Request.Context()

		// if e := c.ShouldBindUri(&binding); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }
		// if e := c.ShouldBindJSON(&updatedNote); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }

		// var row = db.QueryRowContext(ctx, `SELECT * FROM Notes WHERE noteID = %d`, note.noteID)

		// if e := row.Scan(fmt.Sprintf(`%v.noteName, %v.noteText, %v.noteCompletionTime, %v.noteStatus, %v.noteDelegation`, &note)); e != nil {
		// 	if e == sql.ErrNoRows {
		// 		var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 		c.JSON(http.StatusNotFound, res)
		// 		return
		// 	}

		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// var option = copier.Option{
		// 	IgnoreEmpty: true,
		// 	DeepCopy:    true,
		// }

		// if e := copier.CopyWithOption(&note, &updatedNote, option); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// // var updatedRow = db.QueryRowContext(ctx, `SELECT * FROM Notes WHERE noteID = %d`, note.noteID)
		// // var newNote interface{}

		// // if e := updatedRow.Scan(fmt.Sprintf(`%v.noteName, %v.noteText, %v.noteCompletionTime, %v.noteStatus, %v.noteDelegation`, &newNote)); e != nil {
		// // 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// // 	c.JSON(http.StatusInternalServerError, res)
		// // 	return
		// // }

		// var res = NewHTTPResponse(http.StatusOK, newNote)
		// c.JSON(http.StatusOK, res)
	}
}

// Delete Note HTTP Function
func DeleteNote(db *sqlx.DB, noteID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var binding guidBinding
		// var ctx = c.Request.Context()
		// if e := c.ShouldBindUri(&binding); e != nil {
		// 	var res = NewHTTPResponse(http.StatusBadRequest, e)
		// 	c.JSON(http.StatusBadRequest, res)
		// 	return
		// }

		// var result sql.Result
		// var e error
		// if result, e = db.ExecContext(ctx, `DELETE FROM Notes WHERE noteID = %v`, noteID); e != nil {
		// 	var res = NewHTTPResponse(http.StatusInternalServerError, e)
		// 	c.JSON(http.StatusInternalServerError, res)
		// 	return
		// }

		// if nNotes, _ := result.RowsAffected(); nNotes == 0 {
		// 	var res = NewHTTPResponse(http.StatusNotFound, sql.ErrNoRows)
		// 	c.JSON(http.StatusNotFound, res)
		// 	return
		// }

		// c.JSON(http.StatusNoContent, nil)
	}
}
