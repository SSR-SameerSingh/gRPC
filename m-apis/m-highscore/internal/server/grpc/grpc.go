package grpc

import (
	"context"
	pbhighscore "gRPC/m-apis/m-highscore/v1"

	"net"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

//Grpc struct with address and server
type Grpc struct {
	address string       // grpc service will listen on to this service
	srv     *grpc.Server // server which will be required to set up grpc connection
}

//NewServer to start server
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

//HighScore to keep track of
var HighScore = 999999999999.0

//SetHighScore implementing rpc, sethighscore, gethighscore
func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in  m-highscore is called.")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

//GetHighScore implementing rpc, sethighscore, gethighscore
func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in  m-highscore is called.")
	// HighScore = input.HighScore
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

//ListenAndServe which wil do serving part for us
func (g *Grpc) ListenAndServe() error {
	// 1. Open a tcp port to listen on, requires address, grpc struct has address
	lis, err := net.Listen("tcp", g.address) //opens a tcp port
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	// 2. Initialize the server
	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	// 3. RegisterGrpc struct with the server
	// pbhighscore.RegisterGameServer(g.srv, srv pbhighscore.GameServer)
	//because Grpc struct implements pbhighscore.GameServer interface we can change above line to following line
	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("Starting gRPC server for m-highscore microservice.")

	//actually start serving, listener, returns error
	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-highscore microservice.")
	}
	return nil
}
