package model

type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operation_name"`
	Variables     map[string]interface{} `json:"variables"`
}
