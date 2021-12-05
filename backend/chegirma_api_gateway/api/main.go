// Note: please follow rules

package api

import (
	"net/http"

	_ "chegirma_api_gateway/api/docs"
	v1 "chegirma_api_gateway/api/handler/v1"
	"chegirma_api_gateway/config"
	"chegirma_api_gateway/pkg/event"
	"chegirma_api_gateway/pkg/logger"
	"chegirma_api_gateway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log         logger.Logger
	Cfg         config.Config
	Services    services.ServiceManager
	RedisClient services.RedisManager
	Kafka       *event.Kafka
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")

	router.Use(cors.New(corsConfig))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:         opt.Log,
		Cfg:         opt.Cfg,
		Services:    opt.Services,
		RedisClient: opt.RedisClient,
		Kafka:       opt.Kafka,
	})
	routesV1 := router.Group("/v1")
	routesV1.Use()
	{
		routesV1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    "chegirma api_gateway",
			})
		})
		//Category endpoints
		routesV1.POST("/category", handlerV1.CreateCategory)
		routesV1.GET("/category/:category_id", handlerV1.GetCategory)
		routesV1.GET("/category", handlerV1.GetAllCategories)
		routesV1.PUT("/category/:category_id", handlerV1.UpdateCategory)

		//Product endpoints
		routesV1.POST("/product", handlerV1.CreateProduct)
		routesV1.GET("/product/:product_id", handlerV1.GetProduct)
		routesV1.GET("/product", handlerV1.GetAllProducts)
		routesV1.PUT("/product/:product_id", handlerV1.UpdateProduct)

		//User endpoints
		routesV1.POST("/user", handlerV1.CreateUser)
		routesV1.GET("/user/:user_id", handlerV1.GetUser)
		routesV1.GET("/user", handlerV1.GetAllUsers)
		routesV1.PUT("/user/:user_id", handlerV1.UpdateUser)
		routesV1.POST("/user-exists", handlerV1.UserExists)
	}

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router

}
