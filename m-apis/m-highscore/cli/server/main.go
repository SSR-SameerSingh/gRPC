//to make sure that the server is up and running from "internal" folder

package main

import (
	"flag"
	// grpcSetup "gRPC/m-highscore/internal/server/grpc"
	"gRPC/m-highscore/internal/server/grpc"

	"github.com/rs/zerolog/log"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address where connect with m-highscore service")
	flag.Parse()

	//initial grpc struct, bu calling bew server function
	s := grpc.NewServer(*addressPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc m-highscore")
	}
}
