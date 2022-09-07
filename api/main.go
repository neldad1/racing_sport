package main

import (
	"context"
	"flag"
	"net/http"

	"git.neds.sh/matty/entain/api/proto/racing"
	"git.neds.sh/matty/entain/api/proto/sports"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	apiEndpoint        = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	grpcEndpointRacing = flag.String("grpc-endpoint-racing", "localhost:9000", "gRPC racing server endpoint")
	grpcEndpointSports = flag.String("grpc-endpoint-sports", "localhost:7000", "gRPC sports server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("failed running api server: %s", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Register Racing handler
	if err := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*grpcEndpointRacing,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	// Register Sports handler
	if err := sports.RegisterSportsHandlerFromEndpoint(
		ctx,
		mux,
		*grpcEndpointSports,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	log.Infof("API server listening on: %s", *apiEndpoint)

	return http.ListenAndServe(*apiEndpoint, mux)
}
