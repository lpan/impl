package graph

import (
	"fmt"
	"testing"
)

func Test_Find(t *testing.T) {
	g := DefaultGraph()
	fmt.Println(g.Find("a"))
}

func Test_Find2(t *testing.T) {
	g := DefaultGraph()
	fmt.Println(g.Find2("a"))
}

func Test_Find3(t *testing.T) {
	g := DefaultGraph()
	fmt.Println(g.Find3("a", 2))
}
