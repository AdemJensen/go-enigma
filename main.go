package main

import "fmt"

func main() {
	app := NewCommandLineInterface()
	err := app.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
