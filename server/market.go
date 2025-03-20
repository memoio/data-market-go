package server

import (
	"fmt"
	"time"

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

	type TransactionRecord struct {
		FileName     string    `json:"fileName"`
		Description  string    `json:"description"`
		Price        string    `json:"price"`
		BuyTime      time.Time `json:"buyTime"`
		BuyerAddress string    `json:"buyerAddress"`
	}

	var records []TransactionRecord

	// query db
	err := database.G_DB.Model(&database.FileMemo{}).
		Select(`file_info.name AS file_name,
                file_info.description,
                file_info.price,
                file_memo.buy_time,
                file_memo.user_address AS buyer_address`).
		Joins("INNER JOIN file_info ON file_memo.file_id = file_info.id").
		Where("file_memo.user_address = ?", buyer).
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
