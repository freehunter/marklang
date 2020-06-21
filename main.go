package main

import (
	"fmt"
	"os"
	"marklang/repl"
)

func main() {
	fmt.Printf("Marklang v0.1")
	repl.Start(os.Stdin, os.Stdout)
}