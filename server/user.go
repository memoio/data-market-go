package server

import (
	"net/http"

	"github.com/data-market/internal/database"
	"github.com/gin-gonic/gin"
)

func loadUserModule(r *gin.RouterGroup, h *handler) {
	r.GET("/:address/product-list", h.getAddressProductList)
	r.GET("/:address/download-list", h.getAddressDownloadedList)
	r.GET("/:address/purchase-list", h.getAddressPurchasedList)
}

// user godoc
//
//	@Summary		Get user product list
//	@Description	Get user product list
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"address"
//	@Param			state	path		int	true	"state"
//	@Success		200		{object}	object
//	@Router			/user/{address}/productList/{state} [get]
func (h *handler) getAddressProductList(c *gin.Context) {
	ownerAddress := c.Param("address")
	if ownerAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ownerAddress参数必填"})
		return
	}

	// 2. 构建查询
	var files []OwnerFileResponse
	err := h.db.Model(&database.File{}).
		// 显式指定字段映射（数据库列名）
		Select(
			"name",
			"file_did",
			"file_type",
			"category",
			"price",
			"file_size",
			"upload_time",
			"publish_state",
			"publish_time",
			"purchase_count",
			"download_count",
			"view_count",
			"description",
			"e_tag",
		).
		Where("owner_address = ?", ownerAddress). // 使用数据库列名
		Order("upload_time DESC").
		Scan(&files).Error

	// 处理查询错误
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Database query failed",
			"error":   err.Error(),
		})
		return
	}

	// 处理空结果
	if len(files) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "No records found",
			"data":    []interface{}{}, // 返回空数组
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    files,
	})

}

// user godoc
//
//	@Summary		Get user downloaded list
//	@Description	Get user downloaded list
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"address"
//	@Success		200		{object}	object
//	@Router			/user/{address}/downloadedList [get]
func (h *handler) getAddressDownloadedList(c *gin.Context) {
	// 获取请求参数
	userAddress := c.Param("address")
	if userAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "userAddress parameter is required",
		})
		return
	}

	// 执行联表查询
	var results []FileDownloadResponse
	err := h.db.Model(&database.Download{}).
		Select(
			"file_info.name as file_name",
			"file_info.description as file_description",
			"file_info.price as file_price",
			"downloads.download_date as download_time",
		).
		Joins("LEFT JOIN file_info ON downloads.file_id = file_info.id").
		Where("downloads.user_address = ?", userAddress).
		Scan(&results).Error

	// 处理查询错误
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Database query failed",
			"error":   err.Error(),
		})
		return
	}

	// 处理空结果
	if len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "No records found",
			"data":    []interface{}{}, // 返回空数组
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    results,
	})
}

// user godoc
//
//	@Summary		Get user purchased list
//	@Description	Get user purchased list
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"address"
//	@Success		200		{object}	object
//	@Router			/user/{address}/purchasedList [get]
func (h *handler) getAddressPurchasedList(c *gin.Context) {
	// 从URL参数获取用户地址
	userAddress := c.Param("address")
	if userAddress == "" {
		c.JSON(400, gin.H{"error": "user address is required"})
		return
	}

	// 第一步：从FileMemo表中查询该用户的所有文件ID
	var fileMemos []database.FileMemo
	if err := h.db.Where("user_address = ?", userAddress).Find(&fileMemos).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to query file memos"})
		return
	}

	// 如果没有记录，直接返回空数组
	if len(fileMemos) == 0 {
		c.JSON(200, gin.H{"files": []database.File{}})
		return
	}

	// 收集所有唯一的FileID
	fileIDs := make([]uint, 0, len(fileMemos))
	for _, memo := range fileMemos {
		fileIDs = append(fileIDs, memo.FileID)
	}

	// 第二步：从File表中查询这些文件ID对应的完整文件信息
	var files []database.File
	if err := h.db.Where("file_id IN ?", fileIDs).Find(&files).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to query files"})
		return
	}

	// 返回查询结果
	c.JSON(200, gin.H{"files": files})
}
