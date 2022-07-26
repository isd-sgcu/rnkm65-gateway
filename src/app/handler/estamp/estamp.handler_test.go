package estamp

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/validator"
	mock "github.com/isd-sgcu/rnkm65-gateway/src/mocks/estamp"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EstampHandlerTest struct {
	suite.Suite
	UId            string
	Event1         *proto.Event
	Event2         *proto.Event
	Event3         *proto.Event
	BadRequestErr  *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
	NotFoundErr    *dto.ResponseErr
	InternalErr    *dto.ResponseErr
	ForbiddenErr   *dto.ResponseErr
}

func TestEstampHandler(t *testing.T) {
	suite.Run(t, new(EstampHandlerTest))
}

func (t *EstampHandlerTest) SetupTest() {
	t.UId = faker.UUIDDigit()

	t.Event1 = &proto.Event{
		Id:            faker.UUIDDigit(),
		NameTH:        faker.Word(),
		DescriptionTH: faker.Word(),
		NameEN:        faker.Word(),
		DescriptionEN: faker.Word(),
		Code:          faker.Word(),
	}

	t.Event2 = &proto.Event{
		Id:            faker.UUIDDigit(),
		NameTH:        faker.Word(),
		DescriptionTH: faker.Word(),
		NameEN:        faker.Word(),
		DescriptionEN: faker.Word(),
		Code:          faker.Word(),
	}

	t.Event3 = &proto.Event{
		Id:            faker.UUIDDigit(),
		NameTH:        faker.Word(),
		DescriptionTH: faker.Word(),
		NameEN:        faker.Word(),
		DescriptionEN: faker.Word(),
		Code:          faker.Word(),
	}

	t.BadRequestErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid ID",
		Data:       nil,
	}

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	t.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found",
		Data:       nil,
	}

	t.InternalErr = &dto.ResponseErr{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Data:       nil,
	}

	t.ForbiddenErr = &dto.ResponseErr{
		StatusCode: http.StatusForbidden,
		Message:    "Forbidden resource",
		Data:       nil,
	}
}

func (t *EstampHandlerTest) TestFindByIdSuccess() {
	want := &proto.FindEventByIDResponse{
		Event: t.Event1,
	}

	s := &mock.ServiceMock{}
	s.On("FindEventByID", t.Event1.Id).Return(want, nil)

	v, _ := validator.NewValidator()

	hdr := NewHandler(s, v)

	cm := &mock.ContextMock{}
	cm.On("ID").Return(t.Event1.Id, nil)

	hdr.FindEventByID(cm)

	assert.Equal(t.T(), http.StatusOK, cm.Status)
	assert.Equal(t.T(), want, cm.V)
}

func (t *EstampHandlerTest) TestFindByIdBadRequest() {
	s := &mock.ServiceMock{}
	v, _ := validator.NewValidator()

	hdr := NewHandler(s, v)

	cm := &mock.ContextMock{}
	cm.On("ID").Return("", errors.New(""))

	hdr.FindEventByID(cm)

	assert.Equal(t.T(), http.StatusBadRequest, cm.Status)
}

func (t *EstampHandlerTest) TestFindByIdForbidden() {
	s := &mock.ServiceMock{}
	s.On("FindEventByID", t.Event1.Id).Return(nil, t.ForbiddenErr)
	v, _ := validator.NewValidator()

	hdr := NewHandler(s, v)

	cm := &mock.ContextMock{}
	cm.On("ID").Return(t.Event1.Id, nil)

	hdr.FindEventByID(cm)

	assert.Equal(t.T(), http.StatusForbidden, cm.Status)
}

func (t *EstampHandlerTest) TestFindByIdNotFound() {
	s := &mock.ServiceMock{}
	s.On("FindEventByID", t.Event1.Id).Return(nil, t.NotFoundErr)
	v, _ := validator.NewValidator()

	hdr := NewHandler(s, v)

	cm := &mock.ContextMock{}
	cm.On("ID").Return(t.Event1.Id, nil)

	hdr.FindEventByID(cm)

	assert.Equal(t.T(), http.StatusNotFound, cm.Status)
}

func (t *EstampHandlerTest) TestFindByIdInternal() {
	s := &mock.ServiceMock{}
	s.On("FindEventByID", t.Event1.Id).Return(nil, t.InternalErr)
	v, _ := validator.NewValidator()

	hdr := NewHandler(s, v)

	cm := &mock.ContextMock{}
	cm.On("ID").Return(t.Event1.Id, nil)

	hdr.FindEventByID(cm)

	assert.Equal(t.T(), http.StatusInternalServerError, cm.Status)
}

func (t *EstampHandlerTest) TestFindByIdUnavailable() {
	s := &mock.ServiceMock{}
	s.On("FindEventByID", t.Event1.Id).Return(nil, t.ServiceDownErr)
	v, _ := validator.NewValidator()

	hdr := NewHandler(s, v)

	cm := &mock.ContextMock{}
	cm.On("ID").Return(t.Event1.Id, nil)

	hdr.FindEventByID(cm)

	assert.Equal(t.T(), http.StatusServiceUnavailable, cm.Status)
}
