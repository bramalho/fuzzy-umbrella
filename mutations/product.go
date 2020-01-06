package mutations

import (
	"fuzzy-umbrella/resolvers"
	"fuzzy-umbrella/types"

	"github.com/graphql-go/graphql"
)

// CreateProductMutation mutation
var CreateProductMutation = graphql.Field{
	Type: types.Product,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"quantity": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		product, err := resolvers.CreateProduct(params)
		return product, err
	},
}
