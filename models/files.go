package models

import "time"

type ResorceObject struct {
	SecureURL string    `json:"secure_url" bson:"secure_url"`
	VersionID string    `json:"version_id" bson:"version_id"`
	URL       string    `json:"url" bson:"url"`
	PublicID  string    `json:"public_id" bson:"public_id"`
	FileName  string    `json:"file_name" bson:"file_name"`
	AssetID   string    `json:"asset_id" bson:"asset_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	FileType  string    `json:"file_type" bson:"file_type"`
	Etag      string    `json:"etag" bson:"etag"`
	Bytes     string    `json:"bytes" bson:"bytes"`
	Version   string    `json:"version" bson:"version"`
	Signature string    `json:"signature" bson:"signature"`
}
