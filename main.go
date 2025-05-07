package main

import (
	"bufio"
	"fmt"
	"localhost/calculator/calculator"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get user input
	num1, err := calculator.GetNumber(reader, "Enter first number: ")
	if err != nil {
		log.Fatalf("invalid first number")
	}

	num2, err := calculator.GetNumber(reader, "Enter second number: ")
	if err != nil {
		log.Fatalf("invalid second number")
	}

	fmt.Print("Choose operation (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	// Calculate result
	result, err := calculator.Calculate(num1, num2, op)
	if err != nil {
		log.Fatalf("Calculation error: %v", err)
	}

	fmt.Printf("Result: %.2f\n", result)
}
