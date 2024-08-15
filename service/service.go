package service

import (
	"api-gateway/config"
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() user.UsersClient
	HealthSerivce() health.HealthCheckClient
	LifeStyleService() health.LifeStyleClient
	MedicalRecordService() health.MedicalRecordClient
	WearableService() health.WearableClient
}

type serviceManagerImpl struct {
	userClient    user.UsersClient
	healthClient   health.HealthCheckClient
	lifeStyleClient health.LifeStyleClient
	medicalRecordClient health.MedicalRecordClient
	werableClient health.WearableClient
}

func (s *serviceManagerImpl) UserService() user.UsersClient {
	return s.userClient
}

func (s *serviceManagerImpl) HealthSerivce() health.HealthCheckClient {
	return s.healthClient
}

func (s *serviceManagerImpl) LifeStyleService() health.LifeStyleClient {
    return s.lifeStyleClient
}

func (s *serviceManagerImpl) MedicalRecordService() health.MedicalRecordClient {
    return s.medicalRecordClient
}

func (s *serviceManagerImpl) WearableService() health.WearableClient {
    return s.werableClient
}


func NewServiceManager() (ServiceManager, error) {
	connUser, err := grpc.Dial(
		"l-auth-service"+config.Load().GRPC_USER_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connHealth, err := grpc.Dial(
		"health"+config.Load().GRPC_PRODUCT_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceManagerImpl{
		userClient:         user.NewUsersClient(connUser),
		healthClient:       health.NewHealthCheckClient(connHealth),
		lifeStyleClient:    health.NewLifeStyleClient(connHealth),
		medicalRecordClient: health.NewMedicalRecordClient(connHealth),
		werableClient:      health.NewWearableClient(connHealth),
	}, nil
}

