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
		"user_id": &graphql.ArgumentConfig{ // get this from middleware
			Type: graphql.String,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// todo find a better way
		return resolvers.GetProduct(params)
	},
}

// GetProductsQuery query
var GetProductsQuery = graphql.Field{
	Type: graphql.NewList(types.Product),
	Args: graphql.FieldConfigArgument{
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetProducts(params)
	},
}
