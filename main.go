package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Buy products", Completed: false},
	{ID: "3", Item: "Make a haircut", Completed: false},
	{ID: "4", Item: "Feed a cat", Completed: true},
}

func getTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}

func postTodo(ctx *gin.Context) {
	var newTodo todo

	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, todos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	ctx.IndentedJSON(http.StatusFound, todo)
}

func toggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed
	ctx.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)

	router.POST("/todos", postTodo)

	router.GET("/todos/:id", getTodo)

	router.PATCH("/todos/:id", toggleTodoStatus)

	err := router.Run(":9090")
	if err != nil {
		return
	}
}
