// Package repository contains data access logic.
package repository

import "context"

// Example repository interface.
// Replace with your actual repositories.

// UserRepository defines the user repository interface.
type UserRepository interface {
	// Add your repository methods here
	// Example:
	// GetByID(ctx context.Context, id string) (*model.User, error)
	// Create(ctx context.Context, user *model.User) error
	// Update(ctx context.Context, user *model.User) error
	// Delete(ctx context.Context, id string) error
}

// userRepository implements UserRepository.
type userRepository struct {
	// Add your dependencies here (database connection, etc.)
}

// NewUserRepository creates a new user repository.
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// Ensure interface implementation at compile time.
var _ UserRepository = (*userRepository)(nil)

// Context type for transaction handling (example).
type ctxKey string

const txKey ctxKey = "tx"

// WithTx returns a context with a transaction.
func WithTx(ctx context.Context, tx interface{}) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

// TxFromContext retrieves the transaction from context.
func TxFromContext(ctx context.Context) interface{} {
	return ctx.Value(txKey)
}
