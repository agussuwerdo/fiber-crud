package handlers

import (
	"context"
	"fiber-crud/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var itemsCollection *mongo.Collection

// SetClient sets the MongoDB client and initializes the collection
func SetClient(mongoClient *mongo.Client, dbName, collectionName string) {
	client = mongoClient
	itemsCollection = client.Database(dbName).Collection(collectionName)
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	item.ID = primitive.NewObjectID()
	_, err := itemsCollection.InsertOne(context.TODO(), item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create item"})
	}

	return c.JSON(item)
}

func int64Ptr(i int64) *int64 { return &i }

func GetItems(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	skip := (page - 1) * limit

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	cursor, err := itemsCollection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch items"})
	}

	var items []models.Item
	if err = cursor.All(context.TODO(), &items); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch items"})
	}

	// Add logic to calculate totalPages and totalCount
	totalCount, _ := itemsCollection.CountDocuments(context.TODO(), bson.M{})
	totalPages := (int(totalCount) + limit - 1) / limit // rounding up to the nearest whole page

	return c.JSON(models.PaginatedItemsResponse{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalCount: int(totalCount),
		Items:      items,
	})
}

func UpdateItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
	}

	var updateData models.Item
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	update := bson.M{"$set": bson.M{
		"name":  updateData.Name,
		"price": updateData.Price,
	}}

	_, err = itemsCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update item"})
	}

	return c.JSON(models.Response{
		Message: "Item updated",
		ID:      id.Hex(),
	})
}

func DeleteItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
	}

	_, err = itemsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete item"})
	}

	return c.JSON(models.Response{
		Message: "Item deleted",
		ID:      id.Hex(),
	})
}
