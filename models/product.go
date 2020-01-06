package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product struct
type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User        *User              `json:"user,omitempty" bson:"user,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Quantity    int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Status      bool               `json:"status,omitempty" bson:"status,omitempty"`
}

// CreateProduct for user
func CreateProduct(product *Product) (*Product, error) {
	product.Status = true

	collection := GetDB().Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)

	return product, nil
}

// GetProductByID for user
func GetProductByID(id string, u User) (*Product, error) {
	product := Product{}
	collection := GetDB().Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	oid, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, Product{ID: oid, User: &u}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// GetProducts for user
func GetProducts(u User) ([]Product, error) {
	var products []Product
	collection := GetDB().Collection("products")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cursor, err := collection.Find(ctx, Product{User: &u})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
