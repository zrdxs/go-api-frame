package database

import "database/sql"

// Transaction holds tx structs
type Transaction struct {
	tx       *sql.Tx
	rollback bool
	commit   bool
}

// New return new tx instance
func New(tx *sql.Tx) *Transaction {
	instance := &Transaction{
		tx: tx,
	}

	return instance
}

// Commit the sql tx
func (t *Transaction) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}

	t.commit = true

	return nil
}

// Rollback aborts the sql tx
func (t *Transaction) Rollback() error {
	if !t.commit && !t.rollback {
		err := t.tx.Rollback()
		if err != nil {
			return err
		}

		t.rollback = true
	}

	return nil
}
