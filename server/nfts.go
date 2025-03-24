package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadNFTModule(r *gin.RouterGroup, h *handler) {
	r.GET("/list", h.nftList)
	r.GET("/mint", h.nftMint)
	r.GET("/approve", h.nftApprove)
	r.GET("/share", h.nftShare)
}

// NFT godoc
//
//	@Summary		Mint NFT
//	@Description	Mint NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/mint [get]
func (h *handler) nftMint(c *gin.Context) {}

// NFT godoc
//
//	@Summary		List NFT
//	@Description	List NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/list [get]
func (h *handler) nftList(c *gin.Context) {
	// query NFT contracts to get nft list
	// question: with what param?
	// or send tx to nft contract from the frontend

	// 1. 从请求参数获取用户地址
	ownerAddress := c.Query("address")
	if ownerAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address is required"})
		return
	}

	// 2. 调用 Alchemy API
	nfts, err := fetchNFTsFromAlchemy(ownerAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch NFTs"})
		return
	}

	// 3. 返回 NFT 列表
	c.JSON(http.StatusOK, gin.H{
		"nfts": nfts,
	})
}

// fetchNFTsFromAlchemy 调用 Alchemy API 获取 NFT 数据
func fetchNFTsFromAlchemy(ownerAddress string) ([]NFT, error) {
	// Alchemy API 端点（替换 YOUR_API_KEY）
	apiURL := fmt.Sprintf("https://eth-mainnet.g.alchemy.com/nft/v2/lY5jDdV8bcHSrF-oOvPjwArwhApbXqh8/getNFTs?owner=%s", ownerAddress)
	// 发送 HTTP GET 请求
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据
	var result struct {
		OwnedNFTs []NFT `json:"ownedNfts"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.OwnedNFTs, nil
}

// NFT godoc
//
//	@Summary		Approve NFT
//	@Description	Approve NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/approve [get]
func (h *handler) nftApprove(c *gin.Context) {}

// NFT godoc
//
//	@Summary		Share NFT
//	@Description	Share NFT
//	@Tags			NFT
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/nft/share [get]
func (h *handler) nftShare(c *gin.Context) {}
