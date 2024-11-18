package server

import "github.com/gin-gonic/gin"

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
//	@Success		200		{object}	object
//	@Router			/user/{address}/productList [get]
func (h *handler) getAddressProductList(c *gin.Context) {
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
