package services

import (
	"context"
	"mime/multipart"
	"os"
	"testing"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCloudinaryAgentImp_SaveFile(t *testing.T) {
	cwd, err := os.Getwd()
	assert.Nil(t, err)
	file, err := os.Open(cwd + "/test/bg.jpg")
	assert.Nil(t, err)
	defer file.Close()
	var mf multipart.File = file
	err = godotenv.Load("../app.env")
	assert.Nil(t, err)
	cloudApi := os.Getenv("CLOUDINARY")
	cld, err := cloudinary.NewFromURL(cloudApi)
	ctx := context.TODO()
	assert.Nil(t, err)
	upload := NewCloudAgent(cld, ctx)
	res, err := upload.SaveFile(mf)
	assert.Nil(t, err)
	t.Log(err)
	assert.NotNil(t, res)
	t.Log(res)
}
