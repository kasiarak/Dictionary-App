package graphql

import (
	"github.com/graphql-go/graphql"
)

func CreateSchema() (*graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})

	return &schema, err
}
