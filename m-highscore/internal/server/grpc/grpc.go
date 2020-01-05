package grpc

import (
	"context"
	pbhighscore "gRPC/m-apis/m-highscore/v1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

//Grpc struct with address and server
type Grpc struct {
	address string       // grpc service will listen on to this service
	srv     *grpc.Server // server which will be required to set up grpc connection
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
