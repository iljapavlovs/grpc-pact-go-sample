package pact_sample

import (
	"fmt"
	pb "github.com/iljapavlovs/grpc-pact-go-sample/pact-sample/routeguide"
	"github.com/iljapavlovs/grpc-pact-go-sample/pact-sample/routeguide/server"

	"log"

	"net"
	"testing"

	l "github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestGrpcProvider(t *testing.T) {
	go startProvider()
	l.SetLogLevel("TRACE")

	verifier := provider.NewVerifier()

	err := verifier.VerifyProvider(t, provider.VerifyRequest{
		ProviderBaseURL: "http://localhost:8222",
		Transports: []provider.Transport{
			provider.Transport{
				Protocol: "grpc",
				Port:     8222,
			},
		},
		Provider:        "grpcprovider",
		ProviderVersion: "baf529b",                //os.Getenv("APP_SHA"),
		BrokerURL:       "http://localhost:9292/", //os.Getenv("PACT_BROKER_URL")
		//PactFiles: []string{
		//	filepath.ToSlash(fmt.Sprintf("%s/../pacts/grpcconsumer-grpcprovider.json", dir)),
		//},

		PublishVerificationResults: true,

		ConsumerVersionSelectors: []provider.Selector{
			&provider.ConsumerVersionSelector{
				Tag: "master",
			},
			&provider.ConsumerVersionSelector{
				Tag: "prod",
			},
			&provider.ConsumerVersionSelector{
				Tag: "1.0.0",
			},
			&provider.ConsumerVersionSelector{
				Branch: "main",
			},

			&provider.ConsumerVersionSelector{
				Deployed: true,
			},

			&provider.ConsumerVersionSelector{
				Released: true,
			},
		},
		ProviderBranch: "main",
		ProviderTags:   []string{"1.0.1"},
	})

	assert.NoError(t, err)
}

func startProvider() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8222))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
