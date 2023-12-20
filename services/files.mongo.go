package services

import (
	"context"
	"time"

	"github.com/joey1123455/file-upload/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileDB interface {
	SaveFile(
		SecureURL, VersionID, URL, PublicID, FileName, AssetID,
		FileType, Etag, Bytes, Version, Signature string,
	) (*mongo.InsertOneResult, error)

	RetrieveAll() ([]models.ResorceObject, error)
	RetrieveOne(fileName string) (*models.ResorceObject, error)
}

type FileDBImp struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewFileDB(coll *mongo.Collection, ctx context.Context) FileDB {
	return &FileDBImp{
		collection: coll,
		ctx:        ctx,
	}
}

func (fdb FileDBImp) SaveFile(
	secureURL,
	versionID,
	uRL,
	publicID,
	fileName,
	assetID,
	fileType,
	etag,
	bytes,
	version,
	signature string) (*mongo.InsertOneResult, error) {
	var object = &models.ResorceObject{
		CreatedAt: time.Now(),
		AssetID:   assetID,
		Bytes:     bytes,
		Etag:      etag,
		FileName:  fileName,
		URL:       uRL,
		SecureURL: secureURL,
		PublicID:  publicID,
		FileType:  fileType,
		VersionID: versionID,
		Version:   version,
		Signature: signature,
	}

	res, err := fdb.collection.InsertOne(fdb.ctx, object)
	return res, err
}

func (fdb FileDBImp) RetrieveAll() ([]models.ResorceObject, error) {
	files := make([]models.ResorceObject, 0)
	cur, err := fdb.collection.Find(fdb.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var elem models.ResorceObject
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		filteredElem := models.FilteredResponse(&elem)
		files = append(files, filteredElem)
	}
	return files, nil
}

func (fdb FileDBImp) RetrieveOne(fileName string) (*models.ResorceObject, error) {
	var file *models.ResorceObject
	query := bson.M{"file_name": fileName}
	err := fdb.collection.FindOne(fdb.ctx, query).Decode(&file)
	if err != nil {
		return nil, err
	}
	filteredFile := models.FilteredResponse(file)
	return &filteredFile, err
}
