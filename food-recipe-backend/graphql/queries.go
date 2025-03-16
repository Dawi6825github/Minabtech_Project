package graphql

import "github.com/hasura/go-graphql-client"

var client = graphql.NewClient("http://localhost:8081/v1/graphql", nil)

type Recipe struct {
	ID           int
	Title        string
	Ingredients  string
	Instructions string
}

var GetRecipesQuery struct {
	Recipes []Recipe `graphql:"recipes"`
}
