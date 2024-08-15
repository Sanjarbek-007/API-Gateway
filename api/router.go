package api

import (
 "api-gateway/api/handler"
 middleware "api-gateway/api/middlerware"
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
// @title Api Gateway
// @version 1.0
// @description This is a sample server for Api-gateway Service
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
// @schemes http
func (c *controllerImpl) SetupRoutes(h handler.Handler, logger *slog.Logger) {
    c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router := c.Router.Group("/api")
    router.Use(middleware.Check)
    router.Use(middleware.CheckPermissionMiddleware(h.Enforcer))

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
        health.GET("/getDailyHealthSummary/:date", h.GetDailyHealthSummary)
        health.GET("/getWeeklyHealthSummary/:start_date/:end_date", h.GetWeeklyHealthSummary)
    }

    lifestyle := router.Group("/lifestyle")
    {
        lifestyle.POST("/addLifestyleData", h.AddLifeStyleData)
        lifestyle.GET("/getAllLifestyleData", h.GetLifeStyleData)
        lifestyle.GET("/getLifestyleById/:id", h.GetLifeStyleDataById)
        lifestyle.PUT("/updateLifestyleData", h.UpdateLifeStyleData)
        lifestyle.DELETE("/deleteLifestyleData/:id", h.DeleteLifeStyleData)
    }

    medicalReport := router.Group("/medicalReport")
    {
        medicalReport.POST("/add", h.AddMedicalReport)
        medicalReport.GET("/get", h.GetMedicalReport)
        medicalReport.GET("/getById/:id", h.GetMedicalReportById)
        medicalReport.PUT("/update", h.UpdateMedicalReport)
        medicalReport.DELETE("/delete/:id", h.DeleteMedicalReport)
    }

    wearable := router.Group("/wearable")
    {
        wearable.POST("/add", h.AddWearableData)
        wearable.GET("/get", h.GetWearableData)
        wearable.GET("/getById/:id", h.GetWearableDataById)
        wearable.PUT("/update", h.UpdateWearableData)
        wearable.DELETE("/delete/:id", h.DeleteWearableData)
    }

    notifications := router.Group("/notifications")
    {
        notifications.GET("/getAll", h.GetAllNotifications)
        notifications.GET("/new", h.GetAndMarkNotificationAsRead)
    }
}