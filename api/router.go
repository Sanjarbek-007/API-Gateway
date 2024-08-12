package api

import (
	"api-gateway/api/handler"
	middleware "api-gateway/api/middlerware"
	"api-gateway/config"
	"log/slog"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	_ "api-gateway/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	SetupRoutes(handler.Handler, *slog.Logger, *casbin.Enforcer)
	StartServer(config.Config) error
}

type controllerImpl struct {
	Port   string
	Router *gin.Engine
}

func NewController(router *gin.Engine) Controller {
	return &controllerImpl{Router: router}
}

func (c *controllerImpl) StartServer(cfg config.Config) error {
	c.Port = cfg.HTTP_PORT
	return c.Router.Run(c.Port)
}

// @title Api Getaway
// @version 1.0
// @description api gateway service
// @host localhost:8080
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
// @schemes http
func (c *controllerImpl) SetupRoutes(h handler.Handler, logger *slog.Logger, enforcer *casbin.Enforcer) {

	c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := c.Router.Group("/api")
	router.Use(middleware.IsAuthenticated(), middleware.LogMiddleware(logger), middleware.Authorize(enforcer))

	users := router.Group("/users")
	{
		users.GET("/getUserProfile/:id", h.GetUserProfile)
		users.PUT("/updateUser/:id", h.UpdateUser)
		users.GET("/getUserByEmail/:email", h.GetUserByEmail)
		users.DELETE("/deleteUser/:id", h.DeleteUser)
	}
}
