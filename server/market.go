package server

import (
	"fmt"

	"github.com/data-market/internal/database"
	"github.com/gin-gonic/gin"
)

func loadMarketModule(r *gin.RouterGroup, h *handler) {
	// the buyer's tx-list
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
	// get param
	buyer := c.Param("address")

	var records []TransactionRecord

	// query db
	err := h.db.Model(&database.Access{}).
		Select(`file_info.name AS file_name,
                file_info.description,
                file_info.price,
                access.add_time,
                access.user_address AS buyer_address`).
		Joins("INNER JOIN file_info ON access.file_id = file_info.file_id").
		Where("access.user_address = ?", buyer).
		Scan(&records).Error
	if err != nil {
		c.JSON(500, gin.H{
			"result": fmt.Sprintf("query db failed: %v", err),
		})
	}

	// response result
	c.JSON(200, gin.H{
		"transaction-list": records,
	})
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
