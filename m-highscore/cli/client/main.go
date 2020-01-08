package main

import (
	"flag"
	pbhighscore "gRPC/m-apis/m-highscore/v1"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-highscore gRPC service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("address", *addressPtr, "failed to close connection")
		}
	}()

	//get client
	c := pbhighscore.NewGameClient(conn) //returns client

	if c == nil {
		log.Info().Msg("client nil")
	}

	//calling rpc
	r, err := c.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("Highscore form m-highscore microservice")
	} else {
		log.Error().Msg("Couldn't get highscore")
	}
}
