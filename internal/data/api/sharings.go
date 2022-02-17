package api

import (
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
	"github.com/romarq/visualtez-storage/internal/data/dto"
	Repository "github.com/romarq/visualtez-storage/internal/data/repository"
	Service "github.com/romarq/visualtez-storage/internal/data/service"
	LOG "github.com/romarq/visualtez-storage/internal/logger"
)

type SharingsAPI struct {
	SharingsService Service.SharingsService
}

func InitSharingsAPI(db *mongo.Database) SharingsAPI {
	repository := Repository.InitSharingsRepository(db)
	service := Service.InitSharingsService(repository)
	api := SharingsAPI{SharingsService: service}
	return api
}

// GetSharing - Get sharing by hash (`/sharings/:hash`)
// @ID get-sharing-by-hash
// @Description Get sharing by hash
// @Produce json
// @Param hash path string true "Sharing hash"
// @Success 200 {object} dto.Sharing
// @Failure default {object} Error
// @Router /sharings/{hash} [get]
func (api *SharingsAPI) GetSharing(ctx echo.Context) error {
	hash := ctx.Param("hash")

	if hash == "" {
		return HTTPError(ctx, http.StatusBadRequest, "Invalid parameters.")
	}

	content, err := api.SharingsService.GetSharing(hash)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return HTTPError(ctx, http.StatusNotFound, "The requested sharing doesn't exist.")
		}

		LOG.Debug("Failed to get sharing: %v", err)
		return HTTPError(ctx, http.StatusConflict, "Could not fetch sharing.")
	}

	return ctx.JSON(http.StatusOK, content)
}

// InsertSharing - Insert sharing (`/sharings`)
// @ID insert-sharing
// @Description Inserts a new sharing
// @Accept      json
// @Produce     json
// @Param       content body dto.CreateSharing_Params true "Shared content"
// @Success     200 {object} dto.Sharing
// @Failure     default {object} Error
// @Router /sharings [post]
func (api *SharingsAPI) InsertSharing(ctx echo.Context) error {
	request := new(dto.CreateSharing_Params)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	content, err := api.SharingsService.InsertSharing(*request)
	if err != nil {
		LOG.Debug("Failed to insert sharing: %v", err)
		return HTTPError(ctx, http.StatusConflict, "Could not insert sharing.")
	}

	return ctx.JSON(http.StatusOK, content)
}
