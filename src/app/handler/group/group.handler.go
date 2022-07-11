package group

import (
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	validate "github.com/isd-sgcu/rnkm65-gateway/src/app/validator"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"net/http"
)

type Handler struct {
	service  IService
	validate *validate.DtoValidator
}

func NewHandler(service IService, validate *validate.DtoValidator) *Handler {
	return &Handler{
		service:  service,
		validate: validate,
	}
}

type IContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID() (string, error)
	Param(string) (string, error)
}

type IService interface {
	FindByToken(string) (*proto.Group, *dto.ResponseErr)
	Create(groupDto *dto.GroupDto) (*proto.Group, *dto.ResponseErr)
	//Update(string, *dto.GroupDto) (*proto.Group, *dto.ResponseErr)
	//Delete(string) (bool, *dto.ResponseErr)
}

// FindByToken is a function that get the group data by token
// @Summary Get the group data by token
// @Description Return the group dto if successfully
// @Param id path string true "id"
// @Tags group
// @Accept json
// @Produce json
// @Success 200 {object} proto.Group
// @Failure 400 {object} dto.ResponseBadRequestErr Invalid body request
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 404 {object} dto.ResponseNotfoundErr Not found group
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Router /group/{token} [get]
func (h *Handler) FindByToken(ctx IContext) {
	token, err := ctx.Param("token")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Token",
			Data:       nil,
		})
		return
	}

	group, errRes := h.service.FindByToken(token)
	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		return
	}

	ctx.JSON(http.StatusOK, group)
	return
}

// Create is a function that create new group
// @Summary Create new group
// @Description Return the group dto if successfully
// @Param group body dto.GroupDto true "Group DTO"
// @Tags group
// @Accept json
// @Produce json
// @Success 201 {object} proto.Group
// @Failure 400 {object} dto.ResponseBadRequestErr Invalid request body
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 403 {object} dto.ResponseForbiddenErr Insufficiency permission to create group
// @Failure 404 {object} dto.ResponseNotfoundErr Not found group
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Security     AuthToken
// @Router /group [post]
func (h *Handler) Create(ctx IContext) {
	grpDto := dto.GroupDto{}

	err := ctx.Bind(&grpDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	group, errRes := h.service.Create(&grpDto)
	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		return
	}

	ctx.JSON(http.StatusOK, group)
	return
}

// Update is a function that update the group
// @Summary Update the existing group
// @Description Return the group dto if successfully
// @Param id path string true "id"
// @Param group body dto.GroupDto true "group dto"
// @Tags group
// @Accept json
// @Produce json
// @Success 200 {object} proto.Group
// @Failure 400 {object} dto.ResponseBadRequestErr Invalid ID
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 403 {object} dto.ResponseForbiddenErr Insufficiency permission to update user
// @Failure 404 {object} dto.ResponseNotfoundErr Not found group
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Security     AuthToken
// @Router /group/{id} [put]
//func (h *Handler) Update(ctx IContext) {
//	id, err := ctx.ID()
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, &dto.ResponseErr{
//			StatusCode: http.StatusBadRequest,
//			Message:    err.Error(),
//		})
//		return
//	}
//
//	grpId := ctx.GroupID()
//	if grpId != id {
//		ctx.JSON(http.StatusForbidden, &dto.ResponseErr{
//			StatusCode: http.StatusForbidden,
//			Message:    "Insufficiency permission to update group",
//		})
//		return
//	}
//
//	grpDto := dto.GroupDto{}
//
//	err = ctx.Bind(&grpDto)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, err)
//		return
//	}
//
//	group, errRes := h.service.Update(id, &grpDto)
//	if errRes != nil {
//		ctx.JSON(errRes.StatusCode, errRes)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, group)
//	return
//}

// Delete is a function that delete the group
// @Summary Delete the group
// @Description Return the group dto if successfully
// @Param id path string true "id"
// @Tags group
// @Accept json
// @Produce json
// @Success 200 {bool} true
// @Failure 400 {object} dto.ResponseBadRequestErr Invalid ID
// @Failure 401 {object} dto.ResponseUnauthorizedErr Unauthorized
// @Failure 403 {object} dto.ResponseForbiddenErr Insufficiency permission to delete group
// @Failure 404 {object} dto.ResponseNotfoundErr Not found group
// @Failure 503 {object} dto.ResponseServiceDownErr Service is down
// @Security     AuthToken
// @Router /group/members/{id} [delete]
//func (h *Handler) Delete(ctx IContext) {
//	id, err := ctx.ID()
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, &dto.ResponseErr{
//			StatusCode: http.StatusBadRequest,
//			Message:    err.Error(),
//			Data:       nil,
//		})
//		return
//	}
//
//	group, errRes := h.service.Delete(id)
//	if errRes != nil {
//		ctx.JSON(errRes.StatusCode, errRes)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, group)
//	return
//}
