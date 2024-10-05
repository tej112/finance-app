package utils

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

func GetUUIDField() ent.Field {
	return field.UUID("id", uuid.UUID{}).Default(uuid.New)
}

func GetCreatedAtField() ent.Field {
	return field.Time("created_at").Default(time.Now)
}

func GetUpdatedAtField() ent.Field {
	return field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now)
}

func GetAmountField() ent.Field {
	return field.Float("amount").Positive().SchemaType(map[string]string{
		dialect.MySQL:    "decimal(5, 2)",
		dialect.Postgres: "numeric(5, 2)",
	})
}
