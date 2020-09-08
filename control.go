// 51cz project main.go
package main

import (
	"51cz/common"
	"51cz/service/config"
	"51cz/service/dbcomm"
	"51cz/service/users"
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

/*
	首页LOGO
*/
type Logo struct {
	Logo string `json:"logo"`
}

func index(c *gin.Context) {
	log.Println("ss")
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"title": ""})
}

func getLogos(c *gin.Context) {
	common.PrintHead("getLogos")
	var search config.Search
	r := config.New(dbcomm.GetDB(), config.DEBUG)
	if l, err := r.GetList(search); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 400, "err_msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, l)
	}
	common.PrintTail("getLogos")
}

/*
	修改用户信息
*/
func setUser(c *gin.Context) {
	common.PrintHead("setUser")
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err_code": 400, "err_msg": err.Error()})
		return
	}
	r := users.New(dbcomm.GetDB(), users.DEBUG)
	if err := r.InsertEntity(u, nil); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 400, "err_msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"err_code": 0000, "err_msg": "success"})
	}

	common.PrintTail("setUser")
}

/*
	得到用户信息
*/
func getUser(c *gin.Context) {
	common.PrintHead("getUser")
	userId := c.Query("user_id")
	var search users.Search
	search.UserId = userId
	r := users.New(dbcomm.GetDB(), users.DEBUG)
	if u, err := r.Get(search); err != nil {
		c.JSON(http.StatusOK, gin.H{"err_code": 400, "err_msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, u)
	}
	common.PrintTail("getUser")
}
