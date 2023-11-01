package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	Tel   string `json:"tel"`
	Email string `json:"email"`
}

type Struct_L struct {
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Contact   *Contact   `json:"contact"`
	Account_D *Account_D `json:"account_details"`
}

type Account_D struct {
	FirstName  string `json:"first"`
	MiddleName string `json:"middle"`
	LastName   string `json:"last"`
}

var Login []Struct_L

func EnterRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Login)
}

func CreateRequest(c *gin.Context) {
	var newLogin Struct_L
	err := c.BindJSON(&newLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Login = append(Login, newLogin)
	c.JSON(http.StatusCreated, newLogin)
}

func EditRequest(c *gin.Context) {
	username := c.Param("username")
	var updatedLogin Struct_L
	err := c.BindJSON(&updatedLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	for _, existingLogin := range Login {
		if existingLogin.Username == username {
			existingLogin.Contact = updatedLogin.Contact
			existingLogin.Account_D = updatedLogin.Account_D
			c.JSON(http.StatusOK, existingLogin)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Username not found"})
}

func DeleteRequest(c *gin.Context) {
	username := c.Param("username")
	for i, existingLogin := range Login {
		if existingLogin.Username == username {
			Login = append(Login[:i], Login[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Username not found"})
}

func main() {
	router := gin.Default()

	router.GET("/Login", EnterRequest)

	router.POST("/Register", CreateRequest)

	router.PUT("/Account/Edit/:username", EditRequest)

	router.DELETE("Account/Delete/:username", DeleteRequest)

	router.Run()
}
