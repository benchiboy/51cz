// 51cz project main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	log.Println("ss")
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"title": ""})
}

func test(c *gin.Context) {
	log.Println("ss")

	c.JSON(http.StatusOK, "")
}
