package api

import (
	"fmt"
	. "github.com/bixlabs/go-layout/todo/structures"
	. "github.com/bixlabs/go-layout/todo/use_cases"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TodoRestConfigurator struct {
	handler TodoOperations
	Port   string       `env:"WEB_SERVER_PORT" envDefault:"3000"`
}

func NewTodoRestConfigurator(handler TodoOperations) {
	todoRestConfig := TodoRestConfigurator{handler, ""}

	err := env.Parse(&todoRestConfig)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Content-Type of "application/json" must be used for this endpoint handler
	router.POST("/todo", todoRestConfig.createTodo)
	router.GET("/todo/:id", todoRestConfig.readTodo)
	// Content-Type of "application/json" must be used for this endpoint handler
	router.PUT("/todo", todoRestConfig.updateTodo)
	router.DELETE("/todo/:id", todoRestConfig.deleteTodo)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	err = router.Run(fmt.Sprintf(":%s", todoRestConfig.Port))

	if err != nil {
		panic(err)
	}
	// router.Run(":3000") for a hard coded port
}


func (config TodoRestConfigurator) createTodo(c *gin.Context) {
	var request TodoRequest
	var todo *Todo

	if err := c.ShouldBind(&request); err == nil {
		fmt.Printf("%s", request)
		todo = config.handler.Create(TodoPostToBusinessTodo(request))
		c.String(http.StatusOK, fmt.Sprintf("Create was successful for TODO with name: %s", todo.Name))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (config TodoRestConfigurator) readTodo(c *gin.Context) {
	id := c.Param("id")
	todo := config.handler.Read(id)
	c.String(http.StatusOK, fmt.Sprintf("Read was successful for TODO with ID: %s", todo.ID))
}

type TodoRequest struct {
	ID string `json:"i_am"`
	Name string `json:"title"`
	Description string `json:"the_rest"`
	DueDate time.Time `json:"when_finish"`
}

func TodoPostToBusinessTodo(request TodoRequest) Todo {
	return Todo{ID: request.ID, Name: request.Name, Description: request.Description, DueDate: request.DueDate}
}

func (config TodoRestConfigurator) updateTodo(c *gin.Context) {
	var request TodoRequest
	var todo *Todo

	if c.ShouldBind(&request) == nil {
		todo = config.handler.Update(TodoPostToBusinessTodo(request))
	} else {
		// handle validation case
	}

	c.String(http.StatusOK, fmt.Sprintf("Update was successful for TODO with name: %s", todo.Name))
}

func (config TodoRestConfigurator) deleteTodo(c *gin.Context) {
	id := c.Param("id")
	success := config.handler.Delete(id)
	c.String(http.StatusOK, fmt.Sprintf("Delete was successful %t", success))
}
