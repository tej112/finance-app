package schema

import (
	"finance/utils"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Bank_Transfer holds the schema definition for the Bank_Transfer entity.
type Bank_Transfer struct {
	ent.Schema
}

// Fields of the Bank_Transfer.
func (Bank_Transfer) Fields() []ent.Field {
	return []ent.Field{
		utils.GetUUIDField(),
		utils.GetAmountField(),
		field.UUID("from_bank_id", uuid.UUID{}),
		field.UUID("to_bank_id", uuid.UUID{}),
		field.Time("transaction_date"),
		field.String("notes").Optional(),
		utils.GetCreatedAtField(),
		utils.GetUpdatedAtField(),
	}
}

// Edges of the Bank_Transfer.
func (Bank_Transfer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("from_bank", Bank.Type).
			Ref("debits").
			Unique().
			Field("from_bank_id").
			Required(),

		edge.From("to_bank", Bank.Type).
			Ref("credits").
			Unique().
			Field("to_bank_id").
			Required(),
		edge.To("transactions", Transaction.Type),
		edge.To("loans", Loan.Type),
	}
}
