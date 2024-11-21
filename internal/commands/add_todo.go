package commands

import (
	"time"
	"todo/internal/drivers"
	"todo/internal/models"
)

func AddTodo(title string, deadline time.Time) {
    todo := models.Todo {
        Title: title,
        Deadline: deadline,
    }

    todos := []models.Todo{todo}
    drivers.WriteCsv(todos)
}
