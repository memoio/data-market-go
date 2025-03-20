package server

import (
	"net/http"

	"github.com/data-market/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
	endpoint string
}

// Path: server/server.go
func StartServer(port string) *http.Server {
	// init gin
	gin.SetMode(gin.ReleaseMode)

	// new engine
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to Data Market")
	})

	// for swagger
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// handler for requests
	h := &handler{}

	// todo: add endpoint and proxy address for handler

	// register handler for all requests
	loadFileModule(r.Group("/files"), h)
	loadUserModule(r.Group("/user"), h)
	loadNFTModule(r.Group("/nft"), h)
	loadMarketModule(r.Group("/market"), h)

	return &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
}
