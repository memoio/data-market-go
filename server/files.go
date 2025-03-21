package server

import (
	"net/http"
	"time"

	"github.com/data-market/internal/database"
	"github.com/gin-gonic/gin"
)

func loadFileModule(r *gin.RouterGroup, h *handler) {
	r.POST("/upload", h.uploadFile)
	r.GET("/:fileId/download", h.downloadFile)
	r.GET("/:fileId/info", h.getFileInfo)
	r.POST("/:fileId/delete", h.deleteFile)
	r.POST("/:fileId/collect", h.collectFile)
	r.POST("/:fileId/uncollect", h.uncollectFile)
	r.POST("/:fileId/updateInfo", h.updateFileInfo)
	r.POST("/:fileId/purchase", h.purchaseFile)
	r.GET("/:fileId/share", h.shareFile)
	r.POST("/:fileId/upProduct", h.upProduct)
	r.POST("/:fileId/downProduct", h.downProduct)
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
func (h *handler) uploadFile(c *gin.Context) {
	// todo:
	// call putObjectHandle interface in middleware to upload file to mefs
	// the result of putObjectHandle is the cid of uploaded file
	// use this cid to generate a mfiledid
	// call registerMfileDid(string memory mfileDid, string memory _encode, FileType _ftype, string memory _controller, uint256 _price, string[] memory _keywords)
	// params needed: mfiledid, encode, ftype(private/public), controller, price, keywords

	// client, err := ethclient.DialContext(context.TODO(), c.endpoint)
	// if err != nil {
	// 	return err
	// }
	// defer client.Close()

	// // todo: get proxyAddr from instance

	// proxyIns, err := proxy.NewProxy(c.proxyAddr, client)
	// if err != nil {
	// 	return err
	// }

	// tx, err := proxyIns.RegisterMfileDid(c.didTransactor, c.did.Identifier, encode, ftype, controller.Identifier, price, keywords)
	// if err != nil {
	// 	return err
	// }

}

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
func (h *handler) getFileInfo(c *gin.Context) {
	fid := c.Param("fileId")
	if fid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fid is required"})
		return
	}

	var file database.File

	// 使用 GORM 查询数据库
	result := h.db.First(&file, fid)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, file)
}

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

// Files godoc
//
//	@Summary		Collect file
//	@Description	Collect file
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Param 			userAddr 	query 		string  false 	"user"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/collect [post]
func (h *handler) collectFile(c *gin.Context) {
	fid := c.Param("fileId")
	userAddress := c.Query("userAddr")

	// 查询文件信息
	var file database.File
	if err := h.db.First(&file, fid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 查询 MemoDID 信息
	var memoDID database.MemoDID
	if err := h.db.Where("user_address = ?", userAddress).First(&memoDID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "MemoDID not found for the user"})
		return
	}

	// 创建收藏记录
	collection := database.Collection{
		UserAddress: userAddress,
		FileID:      file.FileID,
		MemoDID:     memoDID.MemoDID,
		CollectTime: time.Now(),
	}

	// 写入数据库
	if err := h.db.Create(&collection).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to collection"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Added to collection successfully", "data": collection})
}

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
