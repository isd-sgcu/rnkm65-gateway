package baan

import (
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"net/http"
)

type Handler struct {
	service     IService
	userService IUserService
}

type IService interface {
	FindAll() ([]*proto.Baan, *dto.ResponseErr)
	FindOne(string) (*proto.Baan, *dto.ResponseErr)
}

type IUserService interface {
	FindOne(string) (*proto.User, *dto.ResponseErr)
}

type IContext interface {
	ID() (string, error)
	UserID() string
	JSON(_ int, v interface{})
}

func NewHandler(service IService, userService IUserService) *Handler {
	return &Handler{service: service, userService: userService}
}

// FindAll is a function that get all baans
// @Summary Get all baans
// @Description Return the array of baan dto if successfully
// @Tags baan
// @Accept json
// @Produce json
// @Success 200 {object} []proto.Baan
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Security     AuthToken
// @Router /baan [get]
func (h *Handler) FindAll(c IContext) {
	result, err := h.service.FindAll()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// FindOne is a function that get the baan data by id
// @Summary Get the baan data by id
// @Description Return the baan dto if successfully
// @Param id path string true "id"
// @Tags baan
// @Accept json
// @Produce json
// @Success 200 {object} proto.Baan
// @Failure 400 {object} dto.ResponseBadRequestErr Invalid body request
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 404 {object} dto.ResponseNotfoundErr Not found baan
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Security     AuthToken
// @Router /baan/{id} [get]
func (h *Handler) FindOne(c IContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})

		return
	}

	result, errRes := h.service.FindOne(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUserBaan is a function that get the baan data by id
// @Summary Get the user's baan
// @Description Return the baan dto if successfully
// @Tags baan
// @Accept json
// @Produce json
// @Success 200 {object} proto.Baan
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 404 {object} dto.ResponseNotfoundErr Not found baan
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Security     AuthToken
// @Router /baan/user [get]
func (h *Handler) GetUserBaan(c IContext) {
	id := c.UserID()

	usr, err := h.userService.FindOne(id)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	result, err := h.service.FindOne(usr.BaanId)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
