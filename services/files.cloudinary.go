package services

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryAgent interface {
	SaveFile(file multipart.File) (*uploader.UploadResult, error)
}

type CloudinaryAgentImp struct {
	cld *cloudinary.Cloudinary
	ctx context.Context
}

func NewCloudAgent(cld *cloudinary.Cloudinary, ctx context.Context) CloudinaryAgent {
	return &CloudinaryAgentImp{
		cld: cld,
		ctx: ctx,
	}
}

func (cu CloudinaryAgentImp) SaveFile(file multipart.File) (*uploader.UploadResult, error) {
	res, err := cu.cld.Upload.Upload(cu.ctx, file, uploader.UploadParams{
		Folder:       "assesment-golang-fileuploader",
		ResourceType: "auto",
	})
	return res, err
}
