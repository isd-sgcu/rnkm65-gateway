package estamp

import (
	"net/http"

	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"

	validate "github.com/isd-sgcu/rnkm65-gateway/src/app/validator"
)

type Handler struct {
	service  IEstampService
	validate *validate.DtoValidator
}

type IContext interface {
	JSON(int, interface{})
	UserID() string
	Bind(interface{}) error
	ID() (string, error)
}

func NewHandler(estampService IEstampService, v *validate.DtoValidator) *Handler {
	return &Handler{
		service:  estampService,
		validate: v,
	}
}

type IEstampService interface {
	FindEventByID(id string) (*proto.FindEventByIDResponse, *dto.ResponseErr)
}

// Get detail of event using event id
// @Summary Get event detail
// @Description Get detail of event using event id
// @Param id path string true "id"
// @Tags event
// @Accept json
// @Produce json
// @Success 200 {object} proto.FindEventByIDResponse OK
// @Failure 400 {object} dto.ResponseBadRequestErr Bad Request
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 500 {object} dto.ResponseInternalErr Internal server error
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Router /estamp/:id [get]
// @Security     AuthToken
func (h *Handler) FindEventByID(ctx IContext) {
	id, err := ctx.ID()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	res, errRes := h.service.FindEventByID(id)

	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
