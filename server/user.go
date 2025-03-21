package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadUserModule(r *gin.RouterGroup, h *handler) {
	r.GET("/:address/productList", h.getAddressProductList)
	r.GET("/:address/downloadedList", h.getAddressDownloadedList)
	r.GET("/:address/purchasedList", h.getAddressPurchasedList)
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

	// response data
	var files []File

	// 执行数据库查询
	result := h.db.Where("owner_address = ?", ownerAddress).Find(&files)

	// 处理查询错误
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "Database query failed",
			Data:    nil,
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "Success",
		Data:    files,
	})

	c.JSON(200, gin.H{})
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
	c.JSON(200, gin.H{})
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
	c.JSON(200, gin.H{})
}
