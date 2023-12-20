package services

import (
	"context"
	"time"

	"github.com/joey1123455/file-upload/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileDB interface {
	SaveFile(
		SecureURL,
		VersionID,
		URL,
		PublicID,
		FileName,
		AssetID,
		FileType,
		Etag,
		Bytes,
		Version,
		Signature string) (*mongo.InsertOneResult, error)
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
