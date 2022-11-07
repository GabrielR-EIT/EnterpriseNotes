package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/users", users)
	r.GET("/notes", notes)
}

func users(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/users.html",
		gin.H{
			"title": "Users",
		},
	)
}

func notes(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/notes.html",
		gin.H{},
	)
}
