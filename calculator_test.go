package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func TestGetNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{"valid integer", "42\n", 42, false},
		{"valid float", "3.14\n", 3.14, false},
		{"negative", "-5.5\n", -5.5, false},
		{"invalid", "abc\n", 0, true},
		{"empty", "\n", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock input
			reader := bufio.NewReader(strings.NewReader(tt.input))

			// Call function
			got, err := getNumber(reader, "Test prompt:")

			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("getNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check value
			if !tt.wantErr && got != tt.want {
				t.Errorf("getNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		operation string
		want      float64
		wantErr   string
	}{
		{"addition", 5.0, 10.0, "+", 15.0, ""},
		{"subtraction", 7.36, 5.3, "-", 2.06, ""},
		{"multiplication", 3.0, 7.0, "*", 21.0, ""},
		{"division", 8.0, 4.0, "/", 2.0, ""},
    {"division by zero", 13.0, 0, "/", 0, "Division by zero"},
    {"invalid operation", 3.0, 4.0, "abc", 0, "Invalid operation"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, tt.operation)

			// Check error cases
			if tt.wantErr != "" {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("Error %q should contain %q", err.Error(), tt.wantErr)
				}
				return
			}

			// Check success cases
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if !almostEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

// Floating-point comparison helper
func almostEqual(a, b float64) bool {
    const tolerance = 1e-8
    return math.Abs(a-b) < tolerance
}
