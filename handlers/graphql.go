package handlers

import (
	"encoding/json"
	"fuzzy-umbrella/mutations"
	"fuzzy-umbrella/queries"
	u "fuzzy-umbrella/utils"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type request struct {
	Query string `json:"query"`
}

// GraphQL handler
var GraphQL = func(w http.ResponseWriter, r *http.Request) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queries.NewRoot().Query,
		Mutation: mutations.NewRoot().Query,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	if r.Body == nil {
		u.Respond(w, u.Message(false, "Invalid graphql query in request body"))
		return
	}

	var req request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.Respond(w, u.Message(false, "Error parsing JSON request body"))
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: req.Query,
		Context:       r.Context(),
	})

	json.NewEncoder(w).Encode(result)
}
