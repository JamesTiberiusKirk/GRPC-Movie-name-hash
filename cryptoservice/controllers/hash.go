package controllers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/JamesTiberiusKirk/moviehash/common/hashmovieservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HashMovieNameServer struct {
	hashmovieservice.UnimplementedHashMovieNameServer
}

func NewHashMovieNameServer() *HashMovieNameServer {
	return &HashMovieNameServer{}
}

const (
	badMovieName  = "Fast And Furious"
	badMovieReply = "Nope...nopenopenopenope"
)

// HashName - hashses the name of a movie
func (h *HashMovieNameServer) HashName(ctx context.Context,
	in *hashmovieservice.MovieNameRequest) (*hashmovieservice.HashedNameReply, error) {
	log.Printf("%v", in)

	movieName := in.GetName()

	if movieName == badMovieName {
		return nil, status.Error(codes.InvalidArgument, badMovieReply)
	}

	hash := sha256.Sum256([]byte(movieName))
	hashString := fmt.Sprintf("%x", hash)

	return &hashmovieservice.HashedNameReply{
		Hash: hashString,
	}, nil
}
