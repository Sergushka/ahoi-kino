package graph

import (
	"github.com/graphql-go/graphql"
)

func LoadFields() graphql.Fields {
	fields := movieFields
	return fields
}

func LoadMutationFields() graphql.Fields {
	fields := movieMutationFields
	return fields
}
