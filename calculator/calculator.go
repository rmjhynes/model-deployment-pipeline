package calculator

import (
	"bufio"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"strconv"
	"strings"
)

// Used for performing postgres database operations
type Repository struct {
	conn *pgx.Conn
}

// Take database connection string and initialise Repository
func NewRepository(ctx context.Context, connStr string) (*Repository, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return &Repository{
		conn: conn,
	}, nil
}

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

// Write calculation details to DB
func (r Repository) SaveCalculation(ctx context.Context, calculation Calculation) (Calculation, error) {
	err := r.conn.QueryRow(ctx,
		"INSERT INTO calculations (number1, number2, operator, result) VALUES ($1, $2, $3, $4) RETURNING id",
		calculation.Number1, calculation.Number2, calculation.Operator, calculation.Result).Scan(&calculation.Id)
	return calculation, err
}

// Get all rows from DB
func (r Repository) GetCalculations(ctx context.Context) (Calculation, error) {
	var calculation Calculation
	query := "SELECT * FROM calculations"
	err := r.conn.QueryRow(ctx, query).Scan(&calculation.Id, &calculation.Number1, &calculation.Number2, &calculation.Operator, &calculation.Result)
	if err != nil {
		return Calculation{}, err
	}
	return calculation, nil
}
