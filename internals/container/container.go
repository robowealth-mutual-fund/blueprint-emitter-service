package container

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/controller"
	userCtrl "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/controller/user"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/infrastructure/database"
	grpcServer "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/infrastructure/grpcServer"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/infrastructure/httpServer"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/infrastructure/jaeger"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/codec"
	usrEmitter "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/emitter/user"
	kafkastreams "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/kafka_stream"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/postgres"
	userSvc "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/service/user"
	userWrp "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/service/user/wrapper"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/utils"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/utils/logrus"
	"github.com/robowealth-mutual-fund/shared-utility/validator"
)

type Container struct {
	container *dig.Container
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}

func (c *Container) Configure() error {

	servicesConstructors := []interface{}{
		config.NewConfiguration,
		grpcServer.NewServer,
		http.NewServeMux,
		httpServer.NewServer,
		runtime.NewServeMux,
		jaeger.NewJaeger,
		logrus.NewLog,
		utils.NewUtils,
		utils.NewCustomValidator,
		validator.NewCustomValidator,

		database.NewServerBase,
		postgres.NewRepository,

		kafkastreams.NewEmitter,
		usrEmitter.NewUserEmitter,

		controller.NewHealthZController,

		userSvc.NewService,
		userCtrl.NewController,
		userWrp.NewWrapper,
		codec.NewCodec,
	}
	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	appConfig := config.NewConfiguration()
	jaeger.NewJaeger(appConfig)
	logrus.NewLog()
	return nil
}

func (c *Container) Start() error {
	log.Info("Start Container")
	if err := c.container.Invoke(func(s *grpcServer.Server, h *httpServer.Server, v *validator.CustomValidator) {
		go func() {
			_ = h.Start()
		}()
		s.Start()

	}); err != nil {
		log.Errorf("%s", err)

		return err
	}

	return nil
}

//MigrateDB ...
func (c *Container) MigrateDB() error {
	log.Info("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}
