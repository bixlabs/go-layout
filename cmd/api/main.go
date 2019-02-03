package main

import (
	"github.com/bixlabs/go-layout/api"
	"github.com/bixlabs/go-layout/todo/use_cases"
	"github.com/bixlabs/go-layout/tools"
)

func main() {
	tools.InitializeLogger()
	todoOperations := use_cases.NewTodoOperationsHandler()
	api.NewTodoRestConfigurator(todoOperations)
}


