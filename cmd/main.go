package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/casbin"
	"api-gateway/config"
	"api-gateway/logs"
	"api-gateway/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("API Gateway started successfully!")
	logger := logs.NewLogger()
	logger.Info("API Gateway started successfully!")

	enforcer, err := casbin.CasbinEnforcer(logger)
	if err != nil {
        log.Println("Error initializing casbin enforcer", "error", err.Error())
		logger.Error("Error initializing enforcer", "error", err.Error())
		return
    }

	config := config.Load()
	serviceManager, err := service.NewServiceManager()
	if err != nil {
		log.Println("Error initializing service manager", "error", err.Error())
		logger.Error("Error initializing service manager", "error", err.Error())
		return
	}


	handler := handler.NewHandler(serviceManager.UserService(), serviceManager.HealthSerivce(), serviceManager.LifeStyleService(),serviceManager.MedicalRecordService(), serviceManager.WearableService(), logger, enforcer)
	controller := api.NewController(gin.Default())
	controller.SetupRoutes(*handler, logger)
	controller.StartServer(config)

}