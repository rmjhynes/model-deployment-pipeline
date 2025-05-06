package calculator

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
	num1, err := GetNumber(reader, "Enter first number: ")
	num2, err := GetNumber(reader, "Enter second number: ")

	fmt.Print("Choose operation (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	// Calculate result
	result, err := Calculate(num1, num2, op)
	if err != nil {
		log.Fatalf("Calculation error: %v", err)
	}

	fmt.Printf("Result: %.2f\n", result)
}

func GetNumber(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	str, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid number: %v", err)
	}
	return val, err
}

func Calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Invalid operation: %s", operation)
	}
}
