package main

import (
	"fmt"
	"github.com/nade-harlow/go-test-exercise/section1/Logger"
	"github.com/nade-harlow/go-test-exercise/section1/slices"
	"github.com/nade-harlow/go-test-exercise/section1/variables"
)

func main() {
	fmt.Println(variables.Swap(1, 2))
	fmt.Println(slices.SumEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	// Create instances of FileLogger and ConsoleLogger
	fileLogger := &Logger.FileLogger{Filename: "log.txt"}
	consoleLogger := &Logger.ConsoleLogger{}

	fileLogger.Log("This is a message to the file logger.")

	consoleLogger.Log("This is a message to the console logger.")
}
