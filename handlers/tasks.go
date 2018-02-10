package handlers

import (
	_ "database/sql"
	_ "net/http"
	_ "strconv"
	"models"

	_ "github.com/labstack/echo"
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task
		// Map imcoming JSON body to the new Task
		c.Bind(&task)
		// Add a task using our new model
		id, err := models.PutTask(db, task.Name)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := models.DeleteTask(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
