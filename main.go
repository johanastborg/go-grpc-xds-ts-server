package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
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
	startTime := time.Now()
	ticker := time.NewTicker(20 * time.Millisecond) // 50 Hz
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case t := <-ticker.C:
			elapsed := t.Sub(startTime).Seconds()

			// Generate a sine wave (e.g., 1 Hz frequency) plus some noise (-0.1 to 0.1)
			noise := rand.Float64()*0.2 - 0.1
			value := math.Sin(2*math.Pi*1.0*elapsed) + noise

			point := &telemetry.TelemetryPoint{
				Value:     value,
				Timestamp: t.UnixNano() / int64(time.Millisecond),
			}
			if err := stream.Send(point); err != nil {
				log.Printf("Error sending point: %v", err)
				return err
			}
		}
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
