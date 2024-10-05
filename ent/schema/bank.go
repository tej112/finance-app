package schema

import (
	"finance/utils"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bank holds the schema definition for the Bank entity.
type Bank struct {
	ent.Schema
}

// Fields of the Bank.
func (Bank) Fields() []ent.Field {
	return []ent.Field{
		utils.GetUUIDField(),
		field.String("name"),
		field.String("account_number").Optional(),
		field.String("ifsc_code").Optional(),
		field.String("branch_name").Optional(),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
		utils.GetCreatedAtField(),
		utils.GetUpdatedAtField(),
	}
}

// Edges of the Bank.
func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("debits", Bank_Transfer.Type),
		edge.To("credits", Bank_Transfer.Type),
		edge.To("transactions", Transaction.Type),
	}
}
