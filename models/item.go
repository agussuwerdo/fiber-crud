package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Price int                `json:"price" bson:"price"`
}

type PaginatedItemsResponse struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalPages int    `json:"totalPages"`
	TotalCount int    `json:"totalCount"`
	Items      []Item `json:"items"`
}

type Response struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}
