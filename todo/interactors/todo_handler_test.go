package interactors

import (
	"github.com/bixlabs/go-layout/todo/structures"
	"github.com/bixlabs/go-layout/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	var operationHandler TodoOperations
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Todo CRUD use cases", func() {

		// Runs at the beginning of all tests
		g.Before(func() {
			operationHandler = NewTodoOperationsHandler()
		})

		// Runs before each test
		g.BeforeEach(func() {
			operationHandler = NewTodoOperationsHandler()
		})

		// Runs after each test
		g.AfterEach(func() {
			operationHandler = nil
		})

		// Runs after all tests
		g.After(func() {
			operationHandler = nil
		})

		// Passing Tests
		g.It("Should create a todo ", func() {
			todo := structures.Todo{ID: "1"}
			result := operationHandler.Create(todo)
			Expect(todo.ID).To(Equal(result.ID))
		})

		g.It("Should read a todo ", func() {
			id := "1"
			result := operationHandler.Read("1")
			Expect(id).To(Equal(result.ID))
		})

		g.It("Should update a todo ", func() {
			todo := structures.Todo{ID: "1"}
			result := operationHandler.Update(todo)
			Expect(todo.ID).To(Equal(result.ID))
		})

		g.It("Should delete a todo ", func() {
			id := "1"
			result := operationHandler.Delete(id)
			Expect(true).To(Equal(result))
		})

		// Pending Test
		g.It("Should delete todo")

		// Exclude Test
		g.Xit("Should delete a todo ", func() {
			id := "1"
			result := operationHandler.Delete(id)
			Expect(true).To(Equal(result))
		})

		// We can use describe inside of other describes
		g.Describe("A Failing case", func() {
			// Failing Test
			g.It("Should delete a todo", func() {
				id := "1"
				result := operationHandler.Delete(id)
				Expect(true).To(Equal(result))
			})
		})
	})
}
