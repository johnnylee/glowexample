package main

import (
	"fmt"
	"github.com/johnnylee/glowexample"
)

func main() {
	g, _ := glowexample.NewExample()
	fmt.Println(g.DotString())
}
