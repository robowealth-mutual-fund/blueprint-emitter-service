package grpcserver

import (
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/controller"
	userCtrl "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/controller/user"
	grpcErrors "github.com/robowealth-mutual-fund/shared-utility/grpc_errors"
	validatorUtils "github.com/robowealth-mutual-fund/shared-utility/validator"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	Config     config.Configuration
	Server     *grpc.Server
	HealthCtrl *controller.HealthZController
	UserCtrl   *userCtrl.Controller
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	userCtrl *userCtrl.Controller,
	validator *validatorUtils.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		grpcErrors.UnaryServerInterceptor(),
		validatorUtils.UnaryServerInterceptor(validator),
	)

	s := &Server{
		Server:     grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:     config,
		HealthCtrl: healthCtrl,
		UserCtrl:   userCtrl,
	}

	s.Configure()

	return s
}
