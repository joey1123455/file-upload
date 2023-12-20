package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joey1123455/file-upload/controllers"
)

type UploadRouter struct {
	controller controllers.UploadController
}

func NewUploadRouter(cont controllers.UploadController) *UploadRouter {
	return &UploadRouter{
		controller: cont,
	}
}

func (rc *UploadRouter) Initiate(rg *gin.RouterGroup) {
	router := rg.Group("/files")

	router.POST("/upload", rc.controller.UploadFiles)
}
