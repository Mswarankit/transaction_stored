package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	service := NewTransactionService()

	// Test creating a valid transaction
	err := service.CreateTransaction(1, 5000, "cars", nil)
	assert.Nil(t, err)

	// Test creating a transaction with a parent
	parentID := int64(1)
	err = service.CreateTransaction(2, 10000, "shopping", &parentID)
	assert.Nil(t, err)

	// Test creating a duplicate transaction
	err = service.CreateTransaction(1, 6000, "electronics", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "transaction already exists", err.Error())
}

func TestGetTransaction(t *testing.T) {
	service := NewTransactionService()
	service.CreateTransaction(1, 5000, "cars", nil)

	// Test getting an existing transaction
	transaction, err := service.GetTransaction(1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), transaction.ID)
	assert.Equal(t, float64(5000), transaction.Amount)
	assert.Equal(t, "cars", transaction.Type)

	// Test getting a non-existent transaction
	transaction, err = service.GetTransaction(999)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "transaction not found", err.Error())
}

func TestGetTransactionsByType(t *testing.T) {
	service := NewTransactionService()
	service.CreateTransaction(1, 5000, "cars", nil)
	service.CreateTransaction(2, 10000, "cars", nil)
	service.CreateTransaction(3, 15000, "shopping", nil)

	// Test getting transactions of an existing type
	carTransactions := service.GetTransactionsByType("cars")
	assert.Equal(t, 2, len(carTransactions))
	assert.Contains(t, carTransactions, int64(1))
	assert.Contains(t, carTransactions, int64(2))

	// Test getting transactions of a non-existent type
	electronicsTransactions := service.GetTransactionsByType("electronics")
	assert.Equal(t, 0, len(electronicsTransactions))
}

func TestCalculateSum(t *testing.T) {
	service := NewTransactionService()
	service.CreateTransaction(1, 5000, "cars", nil)
	parentID := int64(1)
	service.CreateTransaction(2, 10000, "shopping", &parentID)
	service.CreateTransaction(3, 15000, "electronics", &parentID)

	// Test sum of a parent transaction
	sum := service.CalculateSum(1)
	assert.Equal(t, float64(30000), sum)

	// Test sum of a child transaction
	sum = service.CalculateSum(2)
	assert.Equal(t, float64(10000), sum)

	// Test sum of a non-existent transaction
	sum = service.CalculateSum(999)
	assert.Equal(t, float64(0), sum)
}
