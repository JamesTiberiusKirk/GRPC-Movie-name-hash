//go:generate mockgen -package controllers -destination hashmoviename_mock.go -source hashmoviename.go
package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/JamesTiberiusKirk/moviehash/common/hashmovieservice"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type hashMovieService interface {
	HashName(ctx context.Context, in *hashmovieservice.MovieNameRequest,
		opts ...grpc.CallOption) (*hashmovieservice.HashedNameReply, error)
}

type HashMovieNameController struct {
	hashMovieService hashMovieService
}

func NewHashMovieNameController(hms hashMovieService) HashMovieNameController {
	return HashMovieNameController{
		hashMovieService: hms,
	}
}

type MovieNameUserRequest struct {
	Name string `json:"name"`
}

type MovieNameUserResponse struct {
	Hashed string `json:"hashed"`
}

const (
	errProvideName = "Please provie a movie name, make it a good one"
	errInternal    = "Internal server error ¯\\_(ツ)_/¯"
)

func (ctrl *HashMovieNameController) HashMovieNameHandler(c echo.Context) error {
	movieName := MovieNameUserRequest{}
	err := c.Bind(&movieName)
	if err != nil {
		log.Printf("Error from the user %v", err)
		return c.String(http.StatusBadRequest, errProvideName)
	}
	log.Printf("%+v", movieName)

	if movieName.Name == "" {
		return c.String(http.StatusBadRequest, errProvideName)
	}

	in := &hashmovieservice.MovieNameRequest{Name: movieName.Name}

	reply, err := ctrl.hashMovieService.HashName(c.Request().Context(), in)
	if err != nil {
		grpcErr, ok := status.FromError(err)

		if ok && grpcErr.Code() == codes.InvalidArgument {
			log.Printf("GRPC invalid argument %s", grpcErr.Message())
			return c.JSON(http.StatusBadRequest, grpcErr.Message())
		}

		log.Printf("Error from the server %v", err)
		return c.String(http.StatusInternalServerError, errInternal)
	}

	log.Printf("GRPC reply %s", reply.GetHash())

	hashed := MovieNameUserResponse{
		Hashed: reply.GetHash(),
	}
	return c.JSON(http.StatusOK, hashed)
}
