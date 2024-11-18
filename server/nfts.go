package server

import "github.com/gin-gonic/gin"

func loadNFTModule(r *gin.RouterGroup, h *handler) {
	r.GET("/list", h.nftList)
	r.GET("/mint", h.nftMint)
	r.GET("/approve", h.nftApprove)
	r.GET("/share", h.nftShare)
}

// NFT godoc
//	@Summary		Mint NFT
//	@Description	Mint NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/mint [get]
func (h *handler) nftMint(c *gin.Context) {}

// NFT godoc
//	@Summary		List NFT
//	@Description	List NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/list [get]
func (h *handler) nftList(c *gin.Context) {}

// NFT godoc
//	@Summary		Approve NFT
//	@Description	Approve NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/approve [get]
func (h *handler) nftApprove(c *gin.Context) {}

// NFT godoc
//	@Summary		Share NFT
//	@Description	Share NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/share [get]
func (h *handler) nftShare(c *gin.Context) {}
