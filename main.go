package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/johanastborg/go-grpc-xds-ts-server/telemetry"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type streamServer struct {
	telemetry.UnimplementedStreamServiceServer
}

func (s *streamServer) GetLiveStream(_ *emptypb.Empty, stream telemetry.StreamService_GetLiveStreamServer) error {
	value := 0.0
	for {
		point := &telemetry.TelemetryPoint{
			Value:     value,
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
		}
		if err := stream.Send(point); err != nil {
			log.Printf("Error sending point: %v", err)
			return err
		}
		value += 1.0
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	telemetry.RegisterStreamServiceServer(grpcServer, &streamServer{})

	fmt.Printf("gRPC server listening on port %s...\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
