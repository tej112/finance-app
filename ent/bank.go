// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"finance/ent/bank"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Bank is the model entity for the Bank schema.
type Bank struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// AccountNumber holds the value of the "account_number" field.
	AccountNumber string `json:"account_number,omitempty"`
	// IfscCode holds the value of the "ifsc_code" field.
	IfscCode string `json:"ifsc_code,omitempty"`
	// BranchName holds the value of the "branch_name" field.
	BranchName string `json:"branch_name,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BankQuery when eager-loading is set.
	Edges        BankEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BankEdges holds the relations/edges for other nodes in the graph.
type BankEdges struct {
	// Debits holds the value of the debits edge.
	Debits []*Bank_Transfer `json:"debits,omitempty"`
	// Credits holds the value of the credits edge.
	Credits []*Bank_Transfer `json:"credits,omitempty"`
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DebitsOrErr returns the Debits value or an error if the edge
// was not loaded in eager-loading.
func (e BankEdges) DebitsOrErr() ([]*Bank_Transfer, error) {
	if e.loadedTypes[0] {
		return e.Debits, nil
	}
	return nil, &NotLoadedError{edge: "debits"}
}

// CreditsOrErr returns the Credits value or an error if the edge
// was not loaded in eager-loading.
func (e BankEdges) CreditsOrErr() ([]*Bank_Transfer, error) {
	if e.loadedTypes[1] {
		return e.Credits, nil
	}
	return nil, &NotLoadedError{edge: "credits"}
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e BankEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[2] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bank) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bank.FieldMetadata:
			values[i] = new([]byte)
		case bank.FieldName, bank.FieldAccountNumber, bank.FieldIfscCode, bank.FieldBranchName:
			values[i] = new(sql.NullString)
		case bank.FieldCreatedAt, bank.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case bank.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bank fields.
func (b *Bank) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bank.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case bank.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case bank.FieldAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_number", values[i])
			} else if value.Valid {
				b.AccountNumber = value.String
			}
		case bank.FieldIfscCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ifsc_code", values[i])
			} else if value.Valid {
				b.IfscCode = value.String
			}
		case bank.FieldBranchName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field branch_name", values[i])
			} else if value.Valid {
				b.BranchName = value.String
			}
		case bank.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case bank.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case bank.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bank.
// This includes values selected through modifiers, order, etc.
func (b *Bank) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryDebits queries the "debits" edge of the Bank entity.
func (b *Bank) QueryDebits() *BankTransferQuery {
	return NewBankClient(b.config).QueryDebits(b)
}

// QueryCredits queries the "credits" edge of the Bank entity.
func (b *Bank) QueryCredits() *BankTransferQuery {
	return NewBankClient(b.config).QueryCredits(b)
}

// QueryTransactions queries the "transactions" edge of the Bank entity.
func (b *Bank) QueryTransactions() *TransactionQuery {
	return NewBankClient(b.config).QueryTransactions(b)
}

// Update returns a builder for updating this Bank.
// Note that you need to call Bank.Unwrap() before calling this method if this Bank
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bank) Update() *BankUpdateOne {
	return NewBankClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bank entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bank) Unwrap() *Bank {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bank is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bank) String() string {
	var builder strings.Builder
	builder.WriteString("Bank(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("account_number=")
	builder.WriteString(b.AccountNumber)
	builder.WriteString(", ")
	builder.WriteString("ifsc_code=")
	builder.WriteString(b.IfscCode)
	builder.WriteString(", ")
	builder.WriteString("branch_name=")
	builder.WriteString(b.BranchName)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", b.Metadata))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Banks is a parsable slice of Bank.
type Banks []*Bank
