package server

import "github.com/gin-gonic/gin"

func loadMarketModule(r *gin.RouterGroup, h *handler) {
	r.GET("/:address/transaction-list", h.getMarketAddressTransactionList)
	r.POST("/purchase", h.postMarketPurchase)

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
