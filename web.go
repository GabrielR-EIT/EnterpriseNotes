package main

import (
	"enterprise-notes/controller"
	_ "html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User defined constants here
// const serverPort = ":8080"
// const myPage = "/Enterprise_Notes/"

// Return all notes as JSON
func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Notes)
}

// Start New Server Function
func StartServer() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/notes", GetNotes)

	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/scss", "./static/scss")
	r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	r.LoadHTMLGlob("templates/**/*")
	controller.Router(r)

	log.Println("Server started")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
