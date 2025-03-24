package server

import (
	"net/http"

	"github.com/data-market/docs"
	"github.com/data-market/internal/database"
	"github.com/data-market/internal/logs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var logger = logs.Logger("server")

type handler struct {
	endpoint string
	db       *gorm.DB
}

// Path: server/server.go
func StartServer(port string) *http.Server {
	// init gin
	gin.SetMode(gin.ReleaseMode)

	// new engine
	r := gin.Default()

	// for form file
	r.MaxMultipartMemory = 100 << 20 // 100 MB

	// welcome handler
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to Data Market")
	})

	// for swagger
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// handler for requests
	h := &handler{}
	h.db = database.G_DB

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
