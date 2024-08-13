package api

import (
	"api-gateway/api/handler"
	// middleware "api-gateway/api/middlerware"
	"api-gateway/config"
	"log/slog"

	"github.com/gin-gonic/gin"

	_ "api-gateway/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	SetupRoutes(handler.Handler, *slog.Logger)
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
// @BasePath /
// @schemes http
// SetupRoutes sets up the routes for the API.
func (c *controllerImpl) SetupRoutes(h handler.Handler, logger *slog.Logger) {
    c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router := c.Router.Group("/api")
    // router.Use(middleware.Check)
    // router.Use(middleware.CheckPermissionMiddleware(enforcer))

    users := router.Group("/user")
    {
        users.GET("/profile/:id", h.GetUserProfile)
        users.PUT("/updateUser/:id", h.UpdateUser)
        users.GET("/email/:email", h.GetUserByEmail)
    }

    health := router.Group("/health")
    {
        health.POST("/generate", h.GenerateHealthRecommendations)
        health.GET("/getRealtimeHealthMonitoring/:user_id", h.GetRealtimeHealthMonitoring)
        health.GET("/getDailyHealthSummary/:id", h.GetDailyHealthSummary)
        health.GET("/getWeeklyHealthSummary/:id", h.GetWeeklyHealthSummary)
    }

    lifestyle := router.Group("/lifestyle")
    {
        lifestyle.POST("/addLifestyleData", h.AddLifeStyleData)
        lifestyle.GET("/getAllLifestyleData/:user_id", h.GetLifeStyleData)
        lifestyle.GET("/getLifestyleById/:id", h.GetLifeStyleDataById)
        lifestyle.PUT("/updateLifestyleData", h.UpdateLifeStyleData)
        lifestyle.DELETE("/deleteLifestyleData/:id", h.DeleteLifeStyleData)
    }

    medicalReport := router.Group("/medicalReport")
    {
        medicalReport.POST("/add", h.AddMedicalReport)
        medicalReport.GET("/get/:user_id", h.GetMedicalReport)
        medicalReport.GET("/getById/:id", h.GetMedicalReportById)
        medicalReport.PUT("/update", h.UpdateMedicalReport)
        medicalReport.DELETE("/delete/:id", h.DeleteMedicalReport)
    }

    wearable := router.Group("/wearable")
    {
        wearable.POST("/add", h.AddWearableData)
        wearable.GET("/get/:user_id", h.GetWearableData)
        wearable.GET("/getById/:id", h.GetWearableDataById)
        wearable.PUT("/update", h.UpdateWearableData)
        wearable.DELETE("/delete/:id", h.DeleteWearableData)
    }
}
