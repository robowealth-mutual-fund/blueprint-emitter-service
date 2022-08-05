package grpcclient

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/config"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-emitter-service/pkg/api/v1"

	"google.golang.org/grpc"
)

// HTTPGRPCClient ...
type HTTPGRPCClient struct {
	Config   config.Configuration
	PingPong apiV1.PingPongServiceClient
}

// Connect ...
func (client *HTTPGRPCClient) Connect() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Error("did not connect: %w", err)
	}

	log.Info("Connect to service on", "0.0.0.0:3000")

	client.PingPong = apiV1.NewPingPongServiceClient(conn)
}

// NewHTTPGRPCClient ...
func NewHTTPGRPCClient(config config.Configuration) *HTTPGRPCClient {
	grpcClient := HTTPGRPCClient{
		Config: config,
	}

	grpcClient.Connect()

	return &grpcClient
}
