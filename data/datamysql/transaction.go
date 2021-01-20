package datamysql

import (
	"database/sql"
)

// Transaction struct model
type Transaction struct {
	tx       *sql.Tx
	commit   bool
	rollback bool
}

// NewTransaction return a new instance for transaction
func NewTransaction(tx *sql.Tx) *Transaction {
	instance := &Transaction{
		tx: tx,
	}

	return instance
}

func (t *Transaction) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}

	t.commit = true

	return nil
}

func (t *Transaction) Rollback() error {
	if t != nil && !t.commit && !t.rollback {
		err := t.tx.Rollback()
		if err != nil {
			return err
		}

		t.rollback = true
	}

	return nil
}
