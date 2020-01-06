package types

import (
	"fuzzy-umbrella/utils"

	"github.com/graphql-go/graphql"
)

// Product type
var Product = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: utils.ObjectID,
			},
			"user": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
