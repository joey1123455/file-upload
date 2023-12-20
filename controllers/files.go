package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joey1123455/file-upload/services"
)

type UploadController struct {
	cloudService services.CloudinaryAgent
	fileService  services.FileDB
}

func NewUploadController(cloud services.CloudinaryAgent, coll services.FileDB) UploadController {
	//add params
	return UploadController{
		cloudService: cloud,
		fileService:  coll,
	}
}

func (uc UploadController) UploadFiles(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	// fmt.Println(file)

	res, err := uc.cloudService.SaveFile(file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(res)
	_, err = uc.fileService.SaveFile(
		res.SecureURL, res.VersionID, res.URL, res.PublicID, res.OriginalFilename,
		res.AssetID, res.Format, res.Etag, fmt.Sprint(res.Bytes), fmt.Sprint(res.Version), res.Signature,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})

}

func (UploadController) GetUploadedFiles(c *gin.Context) {}

func (UploadController) ListUploadedFiles(c *gin.Context) {}
