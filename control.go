// 51cz project main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Comments struct {
	ItemNo     string `json:"item_no"`
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id"`
	Text       string `json:"text"`
}

type Items struct {
	ItemNo    string `json:"item_no"`
	ItemTitle string `json:"item_title"`
	ItemDesc  string `json:"item_desc"`
}

type Orders struct {
	OrderNo string `json:"order_no"`
	UserId  string `json:"user_id"`
	ItemNo  string `json:"item_no"`
	TestNo  string `json:"test_no"`
	Answer  string `json:"answer"`
}

func index(c *gin.Context) {
	log.Println("ss")
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"title": ""})
}

func test(c *gin.Context) {
	log.Println("ss")
	c.JSON(http.StatusOK, "")
}

/*

 */
func getQuestionItems(c *gin.Context) {
	log.Println("ss")
	c.JSON(http.StatusOK, "")
}

/*

 */
func getQuestion(c *gin.Context) {
	log.Println("ss")
	c.JSON(http.StatusOK, "")
}

/*

 */
func getNextQuestion(c *gin.Context) {
	log.Println("ss")
	c.JSON(http.StatusOK, "")
}

/*

 */
func putQuestionAnswer(c *gin.Context) {
	log.Println("ss")
	c.JSON(http.StatusOK, "")
}
