package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var messages = make(map[int]Message)
var nextID = 1

func Gethandler(c echo.Context) error {
	var slMassages []Message

	for _, msg := range messages {
		slMassages = append(slMassages, msg)
	}
	return c.JSON(http.StatusOK, &slMassages)
}

func Posthandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Cannot bind message",
		})
	}
	message.ID = nextID
	nextID++

	messages[message.ID] = message
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was successfully added",
	})
}

func PutchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Cannot convert id to int",
		})
	}
	var updateMessage Message
	if err := c.Bind(&updateMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Could not update message",
		})
	}

	if _, exists := messages[id]; !exists {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Message was not found",
		})
	}
	updateMessage.ID = id
	messages[id] = updateMessage

	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Message was updated",
	})
}

func DeletechHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Cannot convert id to int",
		})
	}
	if _, exists := messages[id]; !exists {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Message was not found",
		})
	}
	delete(messages, id)
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "Message was deleted",
	})
}

func main() {
	e := echo.New()
	e.GET("/messages", Gethandler)
	e.POST("/messages", Posthandler)
	e.PATCH("/messages/:id", PutchHandler)
	e.DELETE("/messages/:id", DeletechHandler)
	e.Start(":8080")
}
