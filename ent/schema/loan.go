package schema

import (
	"finance/utils"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Loan holds the schema definition for the Loan entity.
type Loan struct {
	ent.Schema
}

// Fields of the Loan.
func (Loan) Fields() []ent.Field {
	return []ent.Field{
		utils.GetUUIDField(),
		field.Time("due_date"),
		field.Float("emi_amount").Positive(),
		field.Time("emi_schedule"),
		field.UUID("bank_id", uuid.UUID{}),
		utils.GetCreatedAtField(),
		utils.GetUpdatedAtField(),
	}
}

// Edges of the Loan.
func (Loan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bank", Bank.Type).
			Ref("loans").
			Unique().
			Field("bank_id").
			Required(),
		edge.To("transactions", Transaction.Type),
	}
}
