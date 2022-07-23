package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JamesTiberiusKirk/moviehash/common/hashmovieservice"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_apiController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockhashMovieService(ctrl)
	e := echo.New()

	handler := NewHashMovieNameController(mockService)

	t.Run("successful run on", func(t *testing.T) {
		t.Run("normal input", func(t *testing.T) {
			data := "{\"name\":\"Jin-Roh\"}"
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			in := &hashmovieservice.MovieNameRequest{Name: "Jin-Roh"}
			out := &hashmovieservice.HashedNameReply{Hash: "testhash"}
			mockService.EXPECT().
				HashName(gomock.Any(), in).
				Return(out, nil).
				Times(1)

			err := handler.HashMovieNameHandler(c)
			require.NoError(t, err)
			assert.Equal(t, http.StatusOK, c.Response().Status)

			res := rec.Result()
			defer res.Body.Close()
			buf := new(strings.Builder)
			io.Copy(buf, res.Body)

			assert.Equal(t, "{\"hashed\":\"testhash\"}\n", string(buf.String()))
		})
	})
	t.Run("error on", func(t *testing.T) {
		t.Run("badly formatted input", func(t *testing.T) {
			data := "movie name"
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.HashMovieNameHandler(c)
			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, c.Response().Status)

			res := rec.Result()
			defer res.Body.Close()
			buf := new(strings.Builder)
			io.Copy(buf, res.Body)

			assert.Equal(t, "Please provie a movie name, make it a good one", string(buf.String()))
		})
		t.Run("empty input", func(t *testing.T) {
			data := "{}"
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.HashMovieNameHandler(c)
			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, c.Response().Status)

			res := rec.Result()
			defer res.Body.Close()
			buf := new(strings.Builder)
			io.Copy(buf, res.Body)

			assert.Equal(t, "Please provie a movie name, make it a good one", string(buf.String()))
			t.Run("bad movie input", func(t *testing.T) {
				data := "{\"name\":\"Bad Movie\"}"
				req := httptest.NewRequest(echo.POST, "/", strings.NewReader(data))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				in := &hashmovieservice.MovieNameRequest{Name: "Bad Movie"}
				outErr := status.Error(codes.InvalidArgument, "Bad movie error")
				mockService.EXPECT().
					HashName(gomock.Any(), in).
					Return(nil, outErr).
					Times(1)

				err := handler.HashMovieNameHandler(c)
				require.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, c.Response().Status)

				res := rec.Result()
				defer res.Body.Close()
				buf := new(strings.Builder)
				io.Copy(buf, res.Body)

				assert.Equal(t, "\"Bad movie error\"\n", string(buf.String()))
			})
			t.Run("grpc error", func(t *testing.T) {
				data := "{\"name\":\"Jin-Roh\"}"
				req := httptest.NewRequest(echo.POST, "/", strings.NewReader(data))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				in := &hashmovieservice.MovieNameRequest{Name: "Jin-Roh"}
				outErr := status.Error(codes.Internal, "error")
				mockService.EXPECT().
					HashName(gomock.Any(), in).
					Return(nil, outErr).
					Times(1)

				err := handler.HashMovieNameHandler(c)
				require.NoError(t, err)
				assert.Equal(t, http.StatusInternalServerError, c.Response().Status)

				res := rec.Result()
				defer res.Body.Close()
				buf := new(strings.Builder)
				io.Copy(buf, res.Body)

				assert.Equal(t, "Internal server error ¯\\_(ツ)_/¯", string(buf.String()))
			})
		})
	})
}
