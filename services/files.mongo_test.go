package services

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestFileDB_SaveFile(t *testing.T) {
	ctx := context.TODO()
	err := godotenv.Load("../app.env")
	assert.Nil(t, err)
	t.Log(err)
	uri := os.Getenv("MONGO_URI")
	mongoconn := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, mongoconn)
	assert.Nil(t, err)
	coll := client.Database("golang_mongodb").Collection("test")

	// Create a FileDBImp
	fdb := NewFileDB(coll, ctx)
	// Call the SaveFile method
	res, err := fdb.SaveFile("secureURL", "versionID", "URL", "publicID", "fileName", "assetID", "fileType", "etag", "bytes", "version", "signature")
	// Check if the method returns the expected result or error
	assert.NotNil(t, res)
	assert.Nil(t, err)
	defer coll.Drop(ctx)
}

func TestFileDB_RetrieveAll(t *testing.T) {
	ctx := context.TODO()
	err := godotenv.Load("../app.env")
	assert.Nil(t, err)
	t.Log(err)
	uri := os.Getenv("MONGO_URI")
	mongoconn := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, mongoconn)
	assert.Nil(t, err)
	coll := client.Database("golang_mongodb").Collection("test")

	// Create a FileDBImp
	fdb := NewFileDB(coll, ctx)
	res, err := fdb.RetrieveAll()
	// Check if the method returns the expected result or error
	assert.NotNil(t, res)
	assert.Nil(t, err)
	defer coll.Drop(ctx)
}

func TestFileDB_RetrieveOne(t *testing.T) {
	ctx := context.TODO()
	err := godotenv.Load("../app.env")
	assert.Nil(t, err)
	t.Log(err)
	uri := os.Getenv("MONGO_URI")
	mongoconn := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, mongoconn)
	assert.Nil(t, err)
	coll := client.Database("golang_mongodb").Collection("test")

	// Create a FileDBImp
	fdb := NewFileDB(coll, ctx)
	res, err := fdb.SaveFile("secureURL", "versionID", "URL", "publicID", "fileName", "assetID", "fileType", "etag", "bytes", "version", "signature")
	// Check if the method returns the expected result or error
	assert.NotNil(t, res)
	assert.Nil(t, err)
	fRes, err := fdb.RetrieveOne("fileName")
	// Check if the method returns the expected result or error
	assert.NotNil(t, fRes)
	assert.Nil(t, err)
	defer coll.Drop(ctx)
}
