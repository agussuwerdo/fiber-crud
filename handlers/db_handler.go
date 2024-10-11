package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB initializes the MongoDB client and returns it.
func InitMongoDB() (*mongo.Client, error) {
	// Load environment variables from .env file, unless in production
	if strings.ToLower(os.Getenv("ENVIRONMENT")) != "production" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
	}

	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASS")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	mongoDB := os.Getenv("MONGO_DB")

	if mongoDB == "" {
		log.Fatal("MONGO_DB not set in environment")
	}

	// Build the MongoDB URI based on the environment
	var mongoURI string
	if strings.ToLower(mongoHost) == "localhost" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	} else {
		mongoURI = fmt.Sprintf("mongodb+srv://%s/?retryWrites=true&w=majority", mongoHost)
	}

	fmt.Println("mongoURI", mongoURI)

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// If not localhost, add authentication credentials
	if mongoUser != "" && mongoPass != "" {
		credential := options.Credential{
			Username: mongoUser,
			Password: mongoPass,
		}
		clientOptions = clientOptions.SetAuth(credential)
	}

	// Create a new MongoDB client
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
