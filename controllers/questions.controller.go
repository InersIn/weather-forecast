package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-forecast/models"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var err error

func Start(c echo.Context) error {
	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	var res models.Response
	var data models.QuestionData

	data, err = models.GetQuestions("root", "0")
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	sess.Values["type"] = "root"
	sess.Values["id"] = "0"
	sess.Values["progress"] = "onProgress"
	sess.Save(c.Request(), c.Response())

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = data
	return c.JSON(http.StatusOK, res)
}

func Question(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err == nil && sess.Values["progress"] != nil {
		var res models.Response
		var data models.QuestionData

		var types string = c.QueryParam("type")
		var id string = c.QueryParam("id")

		data, err = models.GetQuestions(types, id)
		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Error"
			res.Data = err
			return c.JSON(http.StatusNotFound, res)
		}

		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = data

		return c.JSON(http.StatusOK, res)

	} else {
		err := models.Error{Status: "Error", Message: "Forbiden Request"}
		return c.JSON(http.StatusForbidden, err)
	}
}

func Answer(c echo.Context) error {
	var json_map []models.AnswerResponse
	var res models.Response

	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range json_map {
		fmt.Println(k, v)
	}

	weather := ForwardChaining(json_map)
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = weather

	return c.JSON(http.StatusOK, res)
}
