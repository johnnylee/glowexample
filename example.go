package glowexample

import (
	"fmt"
	"github.com/johnnylee/glow"
	"strings"
)

// ----------------------------------------------------------------------------
// Globals will be passed to each node. You might imagine that this would
// contain configuration information.
type Globals struct {
}

// ----------------------------------------------------------------------------
// Nodes.
func StringWriter(gl *Globals, StrOut chan string) {
	for i := 0; i < 8; i++ {
		StrOut <- fmt.Sprintf("string %v", i)
	}
	close(StrOut)
}

func StringUpperer(gl *Globals, StrIn, StrOut chan string) {
	for s := range StrIn {
		StrOut <- strings.ToUpper(s)
	}
	close(StrOut)
}

func StringPrinter(gl *Globals, StrIn chan string) {
	for s := range StrIn {
		fmt.Println(s)
	}
}

// ----------------------------------------------------------------------------
// Graph.
func NewExample() (*glow.Graph, error) {
	gl := new(Globals)
	g := glow.NewGraph(gl)

	g.AddNode(StringWriter, "Writer", "StrOut")
	g.AddNode(StringUpperer, "Upperer", "StrIn", "StrOut")
	g.AddNode(StringPrinter, "Printer", "StrIn")

	g.Connect(4, "Writer:StrOut", "Upperer:StrIn")
	g.Connect(1, "Upperer:StrOut", "Printer:StrIn")
	g.SetForeground("Printer")

	return g, nil
}
