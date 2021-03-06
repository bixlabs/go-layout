package todo

import (
	"fmt"
	"github.com/bixlabs/go-layout/todo/interactors"
	"github.com/bixlabs/go-layout/todo/structures"
	"github.com/bixlabs/go-layout/tools"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	"time"
)

type todoRestConfigurator struct {
	handler interactors.TodoOperations
}

func NewTodoRestConfigurator(handler interactors.TodoOperations, router *gin.Engine) {
	configureTodoRoutes(todoRestConfigurator{handler}, router)
}

func configureTodoRoutes(restConfig todoRestConfigurator, router *gin.Engine) {
	router.POST("/todo", restConfig.createTodo)
	router.GET("/todo/:id", restConfig.readTodo)
	router.PUT("/todo", restConfig.updateTodo)
	router.DELETE("/todo/:id", restConfig.deleteTodo)
}

// @Summary Create Todo
// @Description Creates a todo given the correct JSON representation of it.
// @Accept  json
// @Produce  json
// @Param todo body structures.Todo true "Todo structure"
// @Success 200 {object} structures.Todo
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todo [post]
func (config todoRestConfigurator) createTodo(c *gin.Context) {
	var request Request
	var todo *structures.Todo

	if err := c.ShouldBind(&request); err == nil {
		tools.Log().WithFields(logrus.Fields{"Request": request}).Info("A request object was received")
		todo = config.handler.Create(RequestToTodo(request))
		c.String(http.StatusOK, fmt.Sprintf("Create was successful for TODO with name: %s", todo.Name))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (config todoRestConfigurator) readTodo(c *gin.Context) {
	id := c.Param("id")
	todo := config.handler.Read(id)
	c.String(http.StatusOK, fmt.Sprintf("Read was successful for TODO with ID: %s", todo.ID))
}

type Request struct {
	ID          string    `json:"i_am"`
	Name        string    `json:"title"`
	Description string    `json:"the_rest"`
	DueDate     time.Time `json:"when_finish"`
}

func RequestToTodo(request Request) structures.Todo {
	return structures.Todo{ID: request.ID, Name: request.Name, Description: request.Description, DueDate: request.DueDate}
}

func (config todoRestConfigurator) updateTodo(c *gin.Context) {
	var request Request
	var todo *structures.Todo

	if c.ShouldBind(&request) == nil {
		todo = config.handler.Update(RequestToTodo(request))
	} else {
		// handle validation case
		tools.Log().Info("Validation case")
	}

	c.String(http.StatusOK, fmt.Sprintf("Update was successful for TODO with name: %s", todo.Name))
}

func (config todoRestConfigurator) deleteTodo(c *gin.Context) {
	id := c.Param("id")
	success := config.handler.Delete(id)
	c.String(http.StatusOK, fmt.Sprintf("Delete was successful %t", success))
}
