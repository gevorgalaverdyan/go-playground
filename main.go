package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

type Error struct {
	Message string `json:"message"`
}

var todos = []todo{
	{ID: "1", Item: "clean", Completed: false},
	{ID: "2", Item: "read", Completed: false},
	{ID: "3", Item: "sing", Completed: false},
}

// http://universities.hipolabs.com/search?country=Canada

// info abt http context ... body, header
func getTodos(c *gin.Context) {
	//change status of req
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodos(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, Error{Message: "Bind issue"})
		return
	}

	for _, x := range todos {
		if x.ID == newTodo.ID {
			c.IndentedJSON(http.StatusBadRequest, Error{Message: "Duplicate ID"})
			return
		}
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, x := range todos {
		if x.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, Error{Message: "Todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func updateTodoStatus(c *gin.Context) {
    id := c.Param("id")
    var updatedTodo todo

    if err := c.BindJSON(&updatedTodo); err != nil {
        c.IndentedJSON(http.StatusBadRequest, Error{Message: "Invalid request body"})
        return
    }

    todo, err := getTodoById(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, Error{Message: "Todo not found"})
        return
    }

    // Update the todo status
    todo.Completed = updatedTodo.Completed
    c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "OK")
	})

	router.GET("/todos", getTodos)

	router.POST("/todos", addTodos)

	router.GET("/todo/:id", getTodo)

	router.PATCH("/todo/:id", updateTodoStatus)

	router.Run("localhost:9090")
}
