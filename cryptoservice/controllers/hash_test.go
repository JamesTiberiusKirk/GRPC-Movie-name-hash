package controllers

import (
	"context"
	"testing"

	"github.com/JamesTiberiusKirk/moviehash/common/hashmovieservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_grpcController(t *testing.T) {
	controller := NewHashMovieNameServer()
	ctx := context.Background()

	t.Run("successful run on", func(t *testing.T) {
		t.Run("normal input", func(t *testing.T) {
			in := &hashmovieservice.MovieNameRequest{
				Name: "Jin-Roh",
			}
			expected := "5e6b15d562170141acd8052539d91fdaf9229c3ec704b8a20e1f8a60a05edeeb"

			actual, err := controller.HashName(ctx, in)
			require.NoError(t, err)
			assert.Equal(t, expected, actual.GetHash())
		})
		t.Run("empty input", func(t *testing.T) {
			in := &hashmovieservice.MovieNameRequest{}
			expected := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

			actual, err := controller.HashName(ctx, in)
			require.NoError(t, err)
			assert.Equal(t, expected, actual.GetHash())
		})
	})
	t.Run("error on bad movie input", func(t *testing.T) {
		in := &hashmovieservice.MovieNameRequest{
			Name: "Fast And Furious",
		}
		expected := "Nope...nopenopenopenope"

		actual, err := controller.HashName(ctx, in)
		require.Nil(t, actual)
		statusErr, ok := status.FromError(err)
		require.Equal(t, true, ok)
		assert.Equal(t, expected, statusErr.Message())
		assert.Equal(t, codes.InvalidArgument, statusErr.Code())
	})
}
