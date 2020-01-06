package queries

import (
	"fuzzy-umbrella/resolvers"
	"fuzzy-umbrella/types"

	"github.com/graphql-go/graphql"
)

// GetProductQuery query
var GetProductQuery = graphql.Field{
	Type: types.Product,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetProduct(params)
	},
}

// GetProductsQuery query
var GetProductsQuery = graphql.Field{
	Type: graphql.NewList(types.Product),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetProducts(params)
	},
}
