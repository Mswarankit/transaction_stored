package transactions

import (
	"errors"
	"sync"
)

type Transaction struct {
	ID       int64   `json:"id"`
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"`
	ParentID *int64  `json:"parent_id,omitempty"`
}

type TransactionService struct {
	transactions map[int64]*Transaction
	typesIndex   map[string][]int64
	mu           sync.RWMutex
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		transactions: make(map[int64]*Transaction),
		typesIndex:   make(map[string][]int64),
	}
}

func (s *TransactionService) CreateTransaction(id int64, amount float64, transactionType string, parentID *int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.transactions[id]; exists {
		return errors.New("transaction already exists")
	}

	transaction := &Transaction{
		ID:       id,
		Amount:   amount,
		Type:     transactionType,
		ParentID: parentID,
	}

	s.transactions[id] = transaction
	s.typesIndex[transactionType] = append(s.typesIndex[transactionType], id)

	return nil
}

func (s *TransactionService) GetTransaction(id int64) (*Transaction, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	transaction, exists := s.transactions[id]
	if !exists {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}

func (s *TransactionService) GetTransactionsByType(transactionType string) []int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.typesIndex[transactionType]
}

func (s *TransactionService) CalculateSum(id int64) float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var sum float64
	visited := make(map[int64]bool)

	var dfs func(int64)
	dfs = func(currentID int64) {
		if visited[currentID] {
			return
		}
		visited[currentID] = true

		if transaction, exists := s.transactions[currentID]; exists {
			sum += transaction.Amount
			for _, t := range s.transactions {
				if t.ParentID != nil && *t.ParentID == currentID {
					dfs(t.ID)
				}
			}
		}
	}

	dfs(id)
	return sum
}
