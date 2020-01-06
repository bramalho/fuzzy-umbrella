package resolvers

import (
	"errors"
	"fuzzy-umbrella/models"

	"github.com/graphql-go/graphql"
)

// GetProduct resolver
func GetProduct(p graphql.ResolveParams) (*models.Product, error) {
	id, ok := p.Args["id"].(string)
	if ok == false {
		return nil, nil
	}
	uid := models.GetUserID(p.Context.Value("user"))
	if uid == "" {
		return nil, errors.New("Invald User ID")
	}

	user, err := models.GetUserByID(uid)
	if err != nil {
		return nil, err
	}

	product, err := models.GetProductByID(id, *user)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProducts resolver
func GetProducts(p graphql.ResolveParams) ([]models.Product, error) {
	uid := models.GetUserID(p.Context.Value("user"))
	if uid == "" {
		return nil, errors.New("Invald User ID")
	}

	user, err := models.GetUserByID(uid)
	if err != nil {
		return nil, err
	}

	products, err := models.GetProducts(*user)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// CreateProduct resolver
func CreateProduct(p graphql.ResolveParams) (*models.Product, error) {
	uid := models.GetUserID(p.Context.Value("user"))
	if uid == "" {
		return nil, errors.New("Invald User ID")
	}

	user, err := models.GetUserByID(uid)
	if err != nil {
		return nil, err
	}

	product, err := models.CreateProduct(&models.Product{
		User:        user,
		Name:        p.Args["name"].(string),
		Description: p.Args["description"].(string),
		Quantity:    p.Args["quantity"].(int),
		Status:      p.Args["status"].(bool),
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}
