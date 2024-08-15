package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"
	"log/slog"
	"github.com/casbin/casbin/v2"
)

type Handler struct {
	User user.UsersClient
	Health health.HealthCheckClient
	Lifestyle health.LifeStyleClient
	Mecdical health.MedicalRecordClient
	Wearable health.WearableClient
	Logger *slog.Logger
	Enforcer *casbin.Enforcer
}

func NewHandler(user user.UsersClient, healthClient health.HealthCheckClient, lifeStyleClient health.LifeStyleClient, medicalRecordClient health.MedicalRecordClient, wearableClient health.WearableClient, logger *slog.Logger, Enforcer *casbin.Enforcer) *Handler {
	return &Handler{
        User:         user,
        Health:  healthClient,
        Lifestyle: lifeStyleClient,
        Mecdical: medicalRecordClient,
        Wearable: wearableClient,
        Logger:        logger,
		Enforcer: Enforcer,
    }
}