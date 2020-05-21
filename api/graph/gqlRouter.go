package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/sergushka/ahoi-kino/log"
)

func New() graphql.Schema {
	gqlFields := LoadFields()
	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: gqlFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.GetLogger().Fatalf("Failed to create new GraphQL Schema, %v", err)
	}
	return schema
}
