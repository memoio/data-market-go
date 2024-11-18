package server

import "github.com/gin-gonic/gin"

func loadFileModule(r *gin.RouterGroup, h *handler) {
	r.POST("/upload", h.uploadFile)
	r.GET("/:fileId/download", h.downloadFile)
	r.POST("/:fileId/collect", h.collectFile)
	r.POST("/:fileId/uncollect", h.uncollectFile)
	r.GET("/:fileId/info", h.getFileInfo)
	r.POST("/:fileId/updateInfo", h.updateFileInfo)
	r.POST("/:fileId/purchase", h.purchaseFile)
	r.GET("/:fileId/share", h.shareFile)
	r.POST("/:fileId/upProduct", h.upProduct)
	r.POST("/:fileId/downProduct", h.downProduct)
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
//	@Summary		Collect file
//	@Description	Collect file
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/collect [post]
func (h *handler) collectFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Uncollect file
//	@Description	Uncollect file
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/uncollect [post]
func (h *handler) uncollectFile(c *gin.Context) {}

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
//	@Summary		Update file info
//	@Description	Update file info
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/updateInfo [post]
func (h *handler) updateFileInfo(c *gin.Context) {}

// Files godoc
//
//	@Summary		Purchase file
//	@Description	Purchase file
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/purchase [post]
func (h *handler) purchaseFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Share file
//	@Description	Share file
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/share [get]
func (h *handler) shareFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Up product
//	@Description	Up product
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/upProduct [post]
func (h *handler) upProduct(c *gin.Context) {}

// Files godoc
//
//	@Summary		Down product
//	@Description	Down product
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/downProduct [post]
func (h *handler) downProduct(c *gin.Context) {}

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
