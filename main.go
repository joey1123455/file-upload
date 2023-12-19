package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
	cld    *cloudinary.Cloudinary
)

func main() {

}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	cloudApi := os.Getenv("CLOUDINARY")

	client, err = connectDB(uri)
	if err != nil {
		log.Fatal(err)
	}
	cld, err = cloudinary.NewFromURL(cloudApi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cld, client)

	fmt.Println("Connected to MongoDB")
}
