package schema

import (
	"entgo.io/ent"
)

// CreditCard holds the schema definition for the CreditCard entity.
type CreditCard struct {
	ent.Schema
}

// Fields of the CreditCard.
func (CreditCard) Fields() []ent.Field {
	// return []ent.Field{
	// 	utils.GetUUIDField(),
	// }
	return nil
}

// Edges of the CreditCard.
func (CreditCard) Edges() []ent.Edge {
	return nil
}
