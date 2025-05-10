//go:build integration

package calculator_test

import (
	"context"
	"localhost/calculator/calculator"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestSaveCalculation(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("..", "testdata", "init-db.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	calculationRepo, err := calculator.NewRepository(ctx, connStr)
	assert.NoError(t, err)

	c, err := calculationRepo.SaveCalculation(ctx, calculator.Calculation{
		Number1:  5,
		Number2:  7,
		Operator: "+",
		Result:   12,
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)

	calculation, err := calculationRepo.GetCalculations(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, calculation)
}
