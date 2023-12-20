package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/joey1123455/file-upload/config"
	"github.com/joey1123455/file-upload/controllers"
	"github.com/joey1123455/file-upload/routes"
	"github.com/joey1123455/file-upload/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	cld            *cloudinary.Cloudinary
	ctx            context.Context
	server         *gin.Engine
	fileCollection *mongo.Collection

	cloudService services.CloudinaryAgent
	fileService  services.FileDB

	uploadController controllers.UploadController
	uploadRouter     *routes.UploadRouter
)

func main() {
	defer client.Disconnect(ctx)
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "healthy"})
	})

	uploadRouter.Initiate(router)
	log.Fatal(server.Run(":8000"))
}

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	client, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	fileCollection = client.Database("golang_mongodb").Collection("files")
	cld, err = cloudinary.NewFromURL(config.CloudKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	server = gin.Default()
	cloudService = services.NewCloudAgent(cld, ctx)
	fileService = services.NewFileDB(fileCollection, ctx)
	uploadController = controllers.NewUploadController(cloudService, fileService)
	uploadRouter = routes.NewUploadRouter(uploadController)
}
