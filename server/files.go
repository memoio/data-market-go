package server

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
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
	r.POST("/:fileId/update-info", h.updateFileInfo)
	//r.POST("/:fileId/purchase", h.purchaseFile)
	r.GET("/:fileId/share", h.shareFile)
	r.POST("/:fileId/publish", h.publish)
	r.POST("/:fileId/unpublish", h.unpublish)
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

	// 获取客户端的 Authorization 头
	authHeader := c.GetHeader("Authorization")

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// 打开文件
	fileReader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer fileReader.Close()

	// 创建一个缓冲区，用于构造 multipart/form-data 请求体
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件字段
	filePart, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create form file"})
		return
	}
	// 拷贝文件数据到文件字段中
	_, err = io.Copy(filePart, fileReader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to copy file data"})
		return
	}

	// 从表单获取参数
	sign := c.PostForm("sign")
	area := c.PostForm("area")

	// 添加其他表单字段
	_ = writer.WriteField("sign", sign) // 替换为实际的 sign 值
	_ = writer.WriteField("area", area) // 替换为实际的 area 值

	// 关闭 multipart writer
	writer.Close()

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", "http://backend:8008/putObject", &requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 将客户端的 Authorization 头添加到中间件的请求中
	req.Header.Set("Authorization", authHeader)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send request to backend"})
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response from backend"})
		return
	}

	// 返回中间件服务的响应
	c.Data(resp.StatusCode, "application/json", respBody)

	// 	代码说明
	// 获取用户地址和文件：

	// 使用 c.PostForm("address") 获取用户地址。

	// 使用 c.FormFile("file") 获取上传的文件。

	// 构造 multipart/form-data 请求体：

	// 使用 multipart.NewWriter 创建一个 multipart/form-data 格式的请求体。

	// 使用 writer.CreateFormFile 添加文件字段。

	// 使用 writer.WriteField 添加其他表单字段（如 address、sign、area）。

	// 创建 HTTP 请求：

	// 使用 http.NewRequest 创建一个 POST 请求，目标地址为 http://backend:8008/putObject。

	// 设置请求头 Content-Type 为 multipart/form-data。

	// 发送请求并处理响应：

	// 使用 http.Client 发送请求。

	// 读取中间件服务的响应，并将其返回给客户端。

	// 通过调用did合约的register方法来为文件注册did
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

	// todo: make filedid with cid; calc filedid hash and store file info into db;
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
	// get fileid
	fid := c.Param("fileId")

	logger.Debug("file id:", fid)

	// get filedid from file table
	var file database.File

	// 使用 GORM 查询数据库
	result := h.db.First(&file, fid)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	fileDID := file.FileDID

	logger.Debug("file did:", fileDID)

	// get cid from file did
	parts := strings.Split(fileDID, ":")
	cid := parts[len(parts)]
	logger.Debug("cid:", cid)

	// 构造中间件的URL
	middlewareURL := fmt.Sprintf("http://localhost:8080//ipfs/getObject/%s", cid)

	// 从请求中获取必要的参数
	sign := c.Query("sign")

	// 获取客户端的 Authorization 头
	authHeader := c.GetHeader("Authorization")

	// 构造请求
	req, err := http.NewRequest("GET", middlewareURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// 添加必要的参数和头信息
	query := req.URL.Query()
	query.Add("sign", sign)
	req.URL.RawQuery = query.Encode()

	// 将客户端的 Authorization 头添加到中间件的请求中
	req.Header.Set("Authorization", authHeader)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to middleware"})
		return
	}
	defer resp.Body.Close()

	// 检查中间件返回的状态码
	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Middleware returned an error"})
		return
	}

	// 将中间件返回的文件流直接转发给客户端
	extraHeaders := map[string]string{
		"Content-Disposition": resp.Header.Get("Content-Disposition"),
	}

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, extraHeaders)

	// 解释代码
	// 构造中间件的URL：你需要指定中间件的地址和路径。

	// 获取参数：从客户端的请求中获取 cid、和 sign 参数。

	// 	获取客户端的 Authorization 头：

	// 使用 c.GetHeader("Authorization") 获取客户端请求中的 Authorization 头。

	// 这个头包含了访问令牌（token）。

	// 将 Authorization 头转发到中间件：

	// 在构造中间件的请求时，将客户端的 Authorization 头添加到请求中：req.Header.Set("Authorization", authHeader)。

	// 这样中间件在收到请求后，会通过它的内置中间件完成 token 到 address 的转换，并设置到它的 context 中。

	// 构造请求：使用 http.NewRequest 创建一个新的 HTTP 请求，并将参数添加到查询字符串中。

	// 发送请求：使用 http.Client 发送请求到中间件。

	// 处理响应：检查中间件返回的状态码，如果状态码不是 200，则返回错误。否则，将中间件返回的文件流直接转发给客户端。
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
func (h *handler) deleteFile(c *gin.Context) {
	fid := c.Param("fileId")

	// 删除记录
	result := h.db.Delete(&database.File{}, fid)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	// 检查是否删除了记录
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

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
//	@Param 			userAddr 	query 		string  false 	"user"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/uncollect [post]
func (h *handler) uncollectFile(c *gin.Context) {
	fid := c.Param("fileId")
	userAddress := c.Query("userAddr")

	// 删除记录
	result := h.db.Where("file_id = ? AND user_address = ?", fid, userAddress).Delete(&database.Collection{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete from collection"})
		return
	}

	// 检查是否删除了记录
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Deleted from collection successfully"})
}

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
func (h *handler) updateFileInfo(c *gin.Context) {
	// 获取 file_id
	fid := c.Param("fileId")

	// 查询数据库中是否存在该记录
	var file database.File
	if err := h.db.First(&file, fid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 从表单中获取字段值
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// 更新记录
	if err := h.db.Save(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update file"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "File updated successfully", "data": file})
}

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
func (h *handler) shareFile(c *gin.Context) {
	// todo: send grantRead tx from the frontend
}

// Files godoc
//
//	@Summary		Publish
//	@Description	Publish
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/publish [post]
func (h *handler) publish(c *gin.Context) {
	// 获取 file_id
	fileID := c.Param("file_id")

	// 查询当前文件的 publish_state
	var file database.File
	if err := h.db.Select("publish_state").First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 判断文件是否已经上架
	if file.PublishState == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "File is already published"})
		return
	}

	// 更新 publish_state 为 1（已上架），并记录上架时间
	result := h.db.Model(&database.File{}).Where("file_id = ?", fileID).Updates(map[string]interface{}{
		"publish_state": 1,
		"publish_time":  time.Now(), // 记录上架时间
	})

	// 检查是否更新成功
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update publish state"})
		return
	}

	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "File published successfully"})
}

// Files godoc
//
//	@Summary		Un publish
//	@Description	Un publish
//	@Tags			files
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/files/{fileId}/unpublish [post]
func (h *handler) unpublish(c *gin.Context) {
	// 获取 file_id
	fid := c.Param("file_id")

	// 更新 publish_state 为 2（已下架）
	result := h.db.Model(&database.File{}).Where("file_id = ?", fid).Update("publish_state", 2)

	// 检查是否更新成功
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update publish state"})
		return
	}

	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "File unpublished successfully"})
}
