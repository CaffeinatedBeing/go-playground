package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Transaction represents a payment transaction
type Transaction struct {
	ID        string
	Amount    float64
	Currency  string
	From      string
	To        string
	Timestamp time.Time
	Status    string
}

// PaymentProcessor handles payment transactions
type PaymentProcessor struct {
	transactions map[string]*Transaction
	balances     map[string]float64
}

// NewPaymentProcessor creates a new payment processor
func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{
		transactions: make(map[string]*Transaction),
		balances:     make(map[string]float64),
	}
}

// ProcessTransaction processes a payment transaction
func (p *PaymentProcessor) ProcessTransaction(tx *Transaction) error {
	// BUG: Basic error handling without context
	if tx.Amount <= 0 {
		return fmt.Errorf("invalid amount")
	}

	if tx.From == "" || tx.To == "" {
		return fmt.Errorf("invalid account")
	}

	// BUG: No error handling for insufficient funds
	if p.balances[tx.From] < tx.Amount {
		return fmt.Errorf("insufficient funds")
	}

	// BUG: No error handling for currency conversion
	if tx.Currency != "USD" {
		return fmt.Errorf("unsupported currency")
	}

	// BUG: No error handling for duplicate transaction
	if _, exists := p.transactions[tx.ID]; exists {
		return fmt.Errorf("duplicate transaction")
	}

	// Process the transaction
	p.balances[tx.From] -= tx.Amount
	p.balances[tx.To] += tx.Amount
	tx.Status = "completed"
	p.transactions[tx.ID] = tx

	return nil
}

// GetTransaction retrieves a transaction by ID
func (p *PaymentProcessor) GetTransaction(id string) (*Transaction, error) {
	// BUG: Basic error handling without context
	tx, exists := p.transactions[id]
	if !exists {
		return nil, fmt.Errorf("transaction not found")
	}
	return tx, nil
}

// GetBalance retrieves the balance for an account
func (p *PaymentProcessor) GetBalance(account string) (float64, error) {
	// BUG: Basic error handling without context
	balance, exists := p.balances[account]
	if !exists {
		return 0, fmt.Errorf("account not found")
	}
	return balance, nil
}

// SimulateNetworkError simulates a network error
func SimulateNetworkError() error {
	if rand.Float32() < 0.3 {
		return fmt.Errorf("network error")
	}
	return nil
}

func main() {
	processor := NewPaymentProcessor()

	// Set initial balances
	processor.balances["account1"] = 1000.0
	processor.balances["account2"] = 500.0

	// Create a transaction
	tx := &Transaction{
		ID:        "tx1",
		Amount:    100.0,
		Currency:  "USD",
		From:      "account1",
		To:        "account2",
		Timestamp: time.Now(),
	}

	// Process the transaction
	err := processor.ProcessTransaction(tx)
	if err != nil {
		fmt.Printf("Error processing transaction: %v\n", err)
		return
	}

	// Try to process the same transaction again
	err = processor.ProcessTransaction(tx)
	if err != nil {
		fmt.Printf("Error processing duplicate transaction: %v\n", err)
	}

	// Try to process a transaction with insufficient funds
	tx2 := &Transaction{
		ID:        "tx2",
		Amount:    2000.0,
		Currency:  "USD",
		From:      "account1",
		To:        "account2",
		Timestamp: time.Now(),
	}
	err = processor.ProcessTransaction(tx2)
	if err != nil {
		fmt.Printf("Error processing transaction with insufficient funds: %v\n", err)
	}

	// Try to process a transaction with unsupported currency
	tx3 := &Transaction{
		ID:        "tx3",
		Amount:    100.0,
		Currency:  "EUR",
		From:      "account1",
		To:        "account2",
		Timestamp: time.Now(),
	}
	err = processor.ProcessTransaction(tx3)
	if err != nil {
		fmt.Printf("Error processing transaction with unsupported currency: %v\n", err)
	}

	// Try to get a non-existent transaction
	_, err = processor.GetTransaction("nonexistent")
	if err != nil {
		fmt.Printf("Error getting non-existent transaction: %v\n", err)
	}

	// Try to get balance for non-existent account
	_, err = processor.GetBalance("nonexistent")
	if err != nil {
		fmt.Printf("Error getting balance for non-existent account: %v\n", err)
	}

	// Simulate network errors
	for i := 0; i < 5; i++ {
		err := SimulateNetworkError()
		if err != nil {
			fmt.Printf("Network error occurred: %v\n", err)
		}
	}
}
