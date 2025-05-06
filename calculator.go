package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get user input
	num1 := getNumber(reader, "Enter first number: ")
	num2 := getNumber(reader, "Enter second number: ")

	fmt.Print("Choose operation (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	// Calculate result
	result, err := calculate(num1, num2, op)
	if err != nil {
		log.Fatalf("Calculation error: %v", err)
	}

	fmt.Printf("Result: %.2f\n", result)
}

func getNumber(reader *bufio.Reader, prompt string) float64 {
	fmt.Print(prompt)
	str, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		log.Fatalf("Invalid number: %v", err)
	}
	return val
}

func calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operation: %s", operation)
	}
}
