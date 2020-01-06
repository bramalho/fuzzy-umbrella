package resolvers

import (
	"errors"
	"fuzzy-umbrella/models"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProduct resolver
func GetProduct(p graphql.ResolveParams) (*models.Product, error) {
	id, ok := p.Args["id"].(string)
	if ok == false {
		return nil, nil
	}
	user, err := models.GetUserByID(p.Context.Value("user").(primitive.ObjectID))
	if err != nil {
		return nil, errors.New("Invald User ID")
	}

	product, err := models.GetProductByID(id, *user)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProducts resolver
func GetProducts(p graphql.ResolveParams) ([]models.Product, error) {
	user, err := models.GetUserByID(p.Context.Value("user").(primitive.ObjectID))
	if err != nil {
		return nil, errors.New("Invald User ID")
	}

	products, err := models.GetProducts(*user)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// CreateProduct resolver
func CreateProduct(p graphql.ResolveParams) (*models.Product, error) {
	user, err := models.GetUserByID(p.Context.Value("user").(primitive.ObjectID))
	if err != nil {
		return nil, errors.New("Invald User ID")
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
