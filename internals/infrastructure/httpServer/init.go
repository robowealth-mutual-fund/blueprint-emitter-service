package httpServer

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/config"
	userCtrl "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/controller/user"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-emitter-service/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
)

type Server struct {
	Config   config.Configuration
	Server   *runtime.ServeMux
	HttpMux  *http.ServeMux
	UserCtrl *userCtrl.Controller
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {
	_ = apiV1.RegisterUserServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
}

func NewServer(config config.Configuration, rmux *runtime.ServeMux, httpMux *http.ServeMux,
	userCtrl *userCtrl.Controller,
) *Server {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	s := &Server{
		Config:   config,
		Server:   rmux,
		HttpMux:  httpMux,
		UserCtrl: userCtrl,
	}
	s.Configure(context.Background(), opts)
	return s
}
