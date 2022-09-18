// Code generated by ent, DO NOT EDIT.

package block

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the block type in the database.
	Label = "block"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldAdhocAccountID holds the string denoting the adhoc_account_id field in the database.
	FieldAdhocAccountID = "adhoc_account_id"
	// FieldBlockHash holds the string denoting the block_hash field in the database.
	FieldBlockHash = "block_hash"
	// FieldBlock holds the string denoting the block field in the database.
	FieldBlock = "block"
	// FieldSendID holds the string denoting the send_id field in the database.
	FieldSendID = "send_id"
	// FieldSubtype holds the string denoting the subtype field in the database.
	FieldSubtype = "subtype"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the block in the database.
	Table = "blocks"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "blocks"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
)

// Columns holds all SQL columns for block fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldAdhocAccountID,
	FieldBlockHash,
	FieldBlock,
	FieldSendID,
	FieldSubtype,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// BlockHashValidator is a validator for the "block_hash" field. It is called by the builders before save.
	BlockHashValidator func(string) error
	// SendIDValidator is a validator for the "send_id" field. It is called by the builders before save.
	SendIDValidator func(string) error
	// SubtypeValidator is a validator for the "subtype" field. It is called by the builders before save.
	SubtypeValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
