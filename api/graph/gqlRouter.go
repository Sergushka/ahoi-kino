package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/sergushka/ahoi-kino/log"
)

func New() graphql.Schema {
	gqlFields := LoadFields()
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: gqlFields}

	mutationFields := LoadMutationFields()
	mutationQuery := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(mutationQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.GetLogger().Fatalf("Failed to create new GraphQL Schema, %v", err)
	}
	return schema
}
