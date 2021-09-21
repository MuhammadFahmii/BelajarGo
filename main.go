package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func main() {
	e := echo.New()
	e.GET("/users", getUsers)
	e.POST("/users", postUser)
	e.Start(":8000")
}

func getUsers(c echo.Context) error {
	user := User{Name: "udin", Email: "123@gmail.com"}
	return c.JSON(http.StatusOK, user)
}

func postUser(c echo.Context) error {
	newUser := User{}
	c.Bind(&newUser)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"msg":  "Success create user",
		"user": newUser,
	})
}
