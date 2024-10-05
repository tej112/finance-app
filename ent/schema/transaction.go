package schema

import (
	"finance/utils"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		utils.GetUUIDField(),
		utils.GetAmountField(),
		field.UUID("bank_id", uuid.UUID{}),
		field.UUID("category_id", uuid.UUID{}).Optional(),
		field.UUID("transfer_id", uuid.UUID{}).Optional(),
		field.Enum("transaction_type").Values(utils.GetTransactionTypeValues()...),
		field.Enum("transfer_type").Nillable().Values(utils.GetTransferTypeValues()...),
		field.Time("transaction_date").Default(time.Now),
		field.String("notes").Optional(),
		utils.GetCreatedAtField(),
		utils.GetUpdatedAtField(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bank", Bank.Type).
			Ref("transactions").
			Unique().
			Field("bank_id").
			Required().
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("category", Category.Type).
			Ref("transactions").
			Unique().
			Field("category_id").
			Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.From("transfer", Bank_Transfer.Type).
			Ref("transactions").
			Unique().
			Field("transfer_id").
			Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}
