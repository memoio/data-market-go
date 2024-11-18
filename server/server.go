package server

import (
	"net/http"

	"github.com/data-market/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
}

// Path: server/server.go
func StartServer(port string) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to Data Market")
	})

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := &handler{}

	loadFileModule(r.Group("/files"), h)
	loadUserModule(r.Group("/user"), h)
	loadNFTModule(r.Group("/nft"), h)
	loadMarketModule(r.Group("/market"), h)
	return &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
}
