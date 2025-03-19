package server

import "github.com/gin-gonic/gin"

func loadMarketModule(r *gin.RouterGroup, h *handler) {
	r.GET("/:address/transactionlist", h.getMarketAddressTransactionList)
	r.POST("/purchase", h.postMarketPurchase)

	r.POST("/:fileId/collect", h.collectFile)
	r.POST("/:fileId/uncollect", h.uncollectFile)
	r.POST("/:fileId/updateInfo", h.updateFileInfo)
	r.POST("/:fileId/purchase", h.purchaseFile)
	r.GET("/:fileId/share", h.shareFile)
	r.POST("/:fileId/upProduct", h.upProduct)
	r.POST("/:fileId/downProduct", h.downProduct)

}

// market godoc
//
//	@Summary		Get user transaction list
//	@Description	Get user transaction list
//	@Tags			market
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"address"
//	@Success		200		{object}	object
//	@Router			/market/{address}/transactionList [get]
func (h *handler) getMarketAddressTransactionList(c *gin.Context) {
	c.JSON(200, gin.H{})
}

// market godoc
//
//	@Summary		Purchase
//	@Description	Purchase
//	@Tags			market
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"address"
//	@Success		200		{object}	object
//	@Router			/market/purchase [post]
func (h *handler) postMarketPurchase(c *gin.Context) {
	c.JSON(200, gin.H{})
}

// Files godoc
//
//	@Summary		Collect file
//	@Description	Collect file
//	@Tags			market
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/market/{fileId}/collect [post]
func (h *handler) collectFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Uncollect file
//	@Description	Uncollect file
//	@Tags			market
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/market/{fileId}/uncollect [post]
func (h *handler) uncollectFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Update file info
//	@Description	Update file info
//	@Tags			market
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/market/{fileId}/updateInfo [post]
func (h *handler) updateFileInfo(c *gin.Context) {}

// Files godoc
//
//	@Summary		Purchase file
//	@Description	Purchase file
//	@Tags			market
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/market/{fileId}/purchase [post]
func (h *handler) purchaseFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Share file
//	@Description	Share file
//	@Tags			market
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/market/{fileId}/share [get]
func (h *handler) shareFile(c *gin.Context) {}

// Files godoc
//
//	@Summary		Up product
//	@Description	Up product
//	@Tags			market
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/market/{fileId}/upProduct [post]
func (h *handler) upProduct(c *gin.Context) {}

// Files godoc
//
//	@Summary		Down product
//	@Description	Down product
//	@Tags			market
//	@Produce		json
//	@Param			fileId	path		string	true	"File ID"
//	@Success		200		{object}	object
//	@Failure		501		{object}	object
//	@Router			/market/{fileId}/downProduct [post]
func (h *handler) downProduct(c *gin.Context) {}
