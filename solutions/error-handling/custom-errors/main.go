package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Custom error types
var (
	ErrInvalidAmount        = errors.New("invalid amount")
	ErrInvalidAccount       = errors.New("invalid account")
	ErrInsufficientFunds    = errors.New("insufficient funds")
	ErrUnsupportedCurrency  = errors.New("unsupported currency")
	ErrDuplicateTransaction = errors.New("duplicate transaction")
	ErrTransactionNotFound  = errors.New("transaction not found")
	ErrAccountNotFound      = errors.New("account not found")
	ErrNetworkError         = errors.New("network error")
)

// TransactionError wraps transaction-related errors with additional context
type TransactionError struct {
	Err     error
	TxID    string
	Amount  float64
	From    string
	To      string
	Context string
}

func (e *TransactionError) Error() string {
	return fmt.Sprintf("%s: %v (tx: %s, amount: %.2f, from: %s, to: %s)",
		e.Context, e.Err, e.TxID, e.Amount, e.From, e.To)
}

func (e *TransactionError) Unwrap() error {
	return e.Err
}

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
	// Validate amount
	if tx.Amount <= 0 {
		return &TransactionError{
			Err:     ErrInvalidAmount,
			TxID:    tx.ID,
			Amount:  tx.Amount,
			From:    tx.From,
			To:      tx.To,
			Context: "amount validation failed",
		}
	}

	// Validate accounts
	if tx.From == "" || tx.To == "" {
		return &TransactionError{
			Err:     ErrInvalidAccount,
			TxID:    tx.ID,
			Amount:  tx.Amount,
			From:    tx.From,
			To:      tx.To,
			Context: "account validation failed",
		}
	}

	// Check for duplicate transaction
	if _, exists := p.transactions[tx.ID]; exists {
		return &TransactionError{
			Err:     ErrDuplicateTransaction,
			TxID:    tx.ID,
			Amount:  tx.Amount,
			From:    tx.From,
			To:      tx.To,
			Context: "duplicate transaction detected",
		}
	}

	// Check currency
	if tx.Currency != "USD" {
		return &TransactionError{
			Err:     ErrUnsupportedCurrency,
			TxID:    tx.ID,
			Amount:  tx.Amount,
			From:    tx.From,
			To:      tx.To,
			Context: "currency validation failed",
		}
	}

	// Check sufficient funds
	balance, exists := p.balances[tx.From]
	if !exists {
		return &TransactionError{
			Err:     ErrAccountNotFound,
			TxID:    tx.ID,
			Amount:  tx.Amount,
			From:    tx.From,
			To:      tx.To,
			Context: "source account not found",
		}
	}

	if balance < tx.Amount {
		return &TransactionError{
			Err:     ErrInsufficientFunds,
			TxID:    tx.ID,
			Amount:  tx.Amount,
			From:    tx.From,
			To:      tx.To,
			Context: "insufficient funds",
		}
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
	tx, exists := p.transactions[id]
	if !exists {
		return nil, &TransactionError{
			Err:     ErrTransactionNotFound,
			TxID:    id,
			Context: "transaction lookup failed",
		}
	}
	return tx, nil
}

// GetBalance retrieves the balance for an account
func (p *PaymentProcessor) GetBalance(account string) (float64, error) {
	balance, exists := p.balances[account]
	if !exists {
		return 0, &TransactionError{
			Err:     ErrAccountNotFound,
			From:    account,
			Context: "balance lookup failed",
		}
	}
	return balance, nil
}

// SimulateNetworkError simulates a network error
func SimulateNetworkError() error {
	if rand.Float32() < 0.3 {
		return ErrNetworkError
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
}
