package controllers

import (
	"golang-echo/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserReq struct {
	ID    int    `json:"id" form:"id" param:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var UserList = make(map[int]models.User, 0)
var LastID int = 1

func GetUser(c echo.Context) (err error) {
	var result []models.User

	for k, _ := range UserList {
		var tempUser models.User
		tempUser = UserList[k]
		result = append(result, tempUser)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success get all users",
		"data":    result,
	})
}

func CreateUser(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := models.User{
		ID:    LastID,
		Name:  req.Name,
		Email: req.Email,
	}
	UserList[LastID] = user
	LastID++
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create user",
		"data":    user,
	})
}

func GetUserByID(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result := UserList[req.ID]
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"data":    result,
	})
}

func UpdateUserByID(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := models.User{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	}
	UserList[req.ID] = user
	result := UserList[req.ID]
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Update user",
		"data":    result,
	})
}

func DeleteUserByID(c echo.Context) (err error) {
	req := new(UserReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	result := map[string]string{
		"response_code": "200",
		"message":       "succsess",
	}
	delete(UserList, req.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"data":    result,
	})
}
