package calculator

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func GetNumber(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	str, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
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
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operation: %s", operation)
	}
}
