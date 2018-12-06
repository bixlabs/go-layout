package test

import (
	"testing"
	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Numbers", func() {
		// Passing Test
		g.It("Should add two numbers ", func() {
			g.Assert(1+1).Equal(2)
		})
		// Failing Test
		g.It("Should match equal numbers", func() {
			g.Assert(4).Equal(4)
		})
		// Pending Test
		g.It("Should substract two numbers")
		// Excluded Test
		g.Xit("Should add two numbers ", func() {
			g.Assert(3+1).Equal(4)
		})
	})
}