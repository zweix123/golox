package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/zweix123/golox/internal/lox"
)

func runFile(path string) {
	fmt.Println("Running file:", path)
	// read file
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(74)
	}
	err = lox.Run(string(bytes))
	if err != nil {
		fmt.Println(err)
		os.Exit(65)
	}
}

func runPrompt() {
	fmt.Println("Running prompt")
	// read input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Exiting prompt")
				break
			}
			fmt.Println("Error reading input:", err)
			continue
		}
		err = lox.Run(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else { // len(os.Args) == 1
		runPrompt()
	}
}
