package server

import "github.com/gin-gonic/gin"

func loadFileModule(r *gin.RouterGroup, h *handler) {
	r.POST("/upload", h.uploadFile)
	r.GET("/:fileId/download", h.downloadFile)
	r.GET("/:fileId/info", h.getFileInfo)
	r.POST("/:fileId/delete", h.deleteFile)
}

// Files godoc
//
//	@Summary		Upload file
//	@Description	Upload file
//	@Tags			files
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file	true	"File"
//	@Success		200		{object}	uploadResponse
//	@Failure		501		{object}	object
//	@Router			/files/upload [post]
func (h *handler) uploadFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Download file
//	@Description	Download file
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{file}		file
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/download [get]
func (h *handler) downloadFile(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
}

// Files godoc
//
//	@Summary		Get file info
//	@Description	Get file info
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	fileInfoResponse
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/info [get]
func (h *handler) getFileInfo(c *gin.Context) {}

// Files godoc
//
//	@Summary		Delete file
//	@Description	Delete file
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/delete [post]
func (h *handler) deleteFile(c *gin.Context) {}
