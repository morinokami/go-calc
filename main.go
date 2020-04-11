package main

import (
	"os"

	"github.com/morinokami/go.calc/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
