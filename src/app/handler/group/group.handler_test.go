package group

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/validator"
	mock "github.com/isd-sgcu/rnkm65-gateway/src/mocks/group"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type GroupHandlerTest struct {
	suite.Suite
	Group          *proto.Group
	GroupDto       *dto.GroupDto
	BindErr        *dto.ResponseErr
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestGroupHandler(t *testing.T) {
	suite.Run(t, new(GroupHandlerTest))
}

func (t *GroupHandlerTest) SetupTest() {
	t.Group = &proto.Group{
		Id:       faker.UUIDDigit(),
		LeaderID: faker.Word(),
		Token:    faker.Word(),
	}

	t.GroupDto = &dto.GroupDto{
		ID:       t.Group.Id,
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
	}

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	t.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Group not found",
		Data:       nil,
	}

	t.BindErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid ID",
	}
}

func (t *GroupHandlerTest) TestFindByTokenSuccess() {
	want := t.Group

	srv := new(mock.ServiceMock)
	srv.On("FindByToken", t.Group.Token).Return(t.Group, nil)

	c := &mock.ContextMock{
		Group:    t.Group,
		GroupDto: t.GroupDto,
	}
	c.On("Param").Return(t.Group.Token, nil)

	v, _ := validator.NewValidator()

	h := NewHandler(srv, v)
	h.FindByToken(c)

	assert.Equal(t.T(), want, c.V)
}

func (t *GroupHandlerTest) TestFindByTokenFoundErr() {
	want := t.NotFoundErr

	srv := new(mock.ServiceMock)
	srv.On("FindByToken", t.Group.Token).Return(nil, t.NotFoundErr)

	c := &mock.ContextMock{
		Group:    t.Group,
		GroupDto: t.GroupDto,
	}
	c.On("Param").Return(t.Group.Token, nil)

	v, _ := validator.NewValidator()

	h := NewHandler(srv, v)
	h.FindByToken(c)

	assert.Equal(t.T(), want, c.V)
}

func (t *GroupHandlerTest) TestFindByTokenInternalErr() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusInternalServerError,
		Message:    "Invalid Token",
		Data:       nil,
	}

	srv := new(mock.ServiceMock)
	srv.On("FindByToken", t.Group.Token).Return(nil, t.ServiceDownErr)

	c := &mock.ContextMock{
		Group:    t.Group,
		GroupDto: t.GroupDto,
	}
	c.On("Param").Return("", errors.New("Cannot parse token"))

	v, _ := validator.NewValidator()

	h := NewHandler(srv, v)
	h.FindByToken(c)

	assert.Equal(t.T(), want, c.V)
}

func (t *GroupHandlerTest) TestFindByTokenGrpcErr() {
	want := t.ServiceDownErr

	srv := new(mock.ServiceMock)
	srv.On("FindByToken", t.Group.Token).Return(nil, t.ServiceDownErr)

	c := &mock.ContextMock{
		Group:    t.Group,
		GroupDto: t.GroupDto,
	}
	c.On("Param").Return(t.Group.Token, nil)

	v, _ := validator.NewValidator()

	h := NewHandler(srv, v)
	h.FindByToken(c)

	assert.Equal(t.T(), want, c.V)
}

func (t *GroupHandlerTest) TestCreateSuccess() {
	want := t.Group

	srv := new(mock.ServiceMock)
	srv.On("Create", t.GroupDto).Return(want, nil)

	c := &mock.ContextMock{
		Group:    t.Group,
		GroupDto: t.GroupDto,
	}
	c.On("Bind", &dto.GroupDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := NewHandler(srv, v)
	h.Create(c)

	assert.Equal(t.T(), want, c.V)
}

func (t *GroupHandlerTest) TestCreateGrpcErr() {
	want := t.ServiceDownErr

	srv := new(mock.ServiceMock)
	srv.On("Create", t.GroupDto).Return(nil, t.ServiceDownErr)

	c := &mock.ContextMock{
		Group:    t.Group,
		GroupDto: t.GroupDto,
	}
	c.On("Bind", &dto.GroupDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := NewHandler(srv, v)
	h.Create(c)

	assert.Equal(t.T(), want, c.V)
}

//func (t *GroupHandlerTest) TestUpdateSuccess() {
//	want := t.Group
//
//	srv := new(mock.ServiceMock)
//	srv.On("Update", t.Group.Id, t.GroupDto).Return(want, nil)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//
//	c.On("ID").Return(t.Group.Id, nil)
//	c.On("GroupID").Return(t.Group.Id, nil)
//	c.On("Bind", &dto.GroupDto{}).Return(nil)
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Update(c)
//
//	assert.Equal(t.T(), want, c.V)
//}
//
//func (t *GroupHandlerTest) TestUpdateForbidden() {
//	want := &dto.ResponseErr{
//		StatusCode: http.StatusForbidden,
//		Message:    "Insufficiency permission to update group",
//	}
//
//	srv := new(mock.ServiceMock)
//	srv.On("Update", t.Group.Id, t.GroupDto).Return(nil, t.NotFoundErr)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return(faker.UUIDDigit(), nil)
//	c.On("GroupID").Return(t.Group.Id, nil)
//	c.On("Bind", &dto.GroupDto{}).Return(nil)
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Update(c)
//
//	assert.Equal(t.T(), want, c.V)
//}
//
//func (t *GroupHandlerTest) TestUpdateNotFound() {
//	want := t.NotFoundErr
//
//	srv := new(mock.ServiceMock)
//	srv.On("Update", t.Group.Id, t.GroupDto).Return(nil, t.NotFoundErr)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return(t.Group.Id, nil)
//	c.On("GroupID").Return(t.Group.Id, nil)
//	c.On("Bind", &dto.GroupDto{}).Return(nil)
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Update(c)
//
//	assert.Equal(t.T(), want, c.V)
//}
//
//func (t *GroupHandlerTest) TestUpdateInvalidID() {
//	want := &dto.ResponseErr{
//		StatusCode: http.StatusBadRequest,
//		Message:    "ID must be the uuid",
//	}
//
//	srv := new(mock.ServiceMock)
//	srv.On("Update", t.Group.Id).Return(nil, t.NotFoundErr)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return("", errors.New(want.Message))
//	c.On("GroupID").Return(t.Group.Id, nil)
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Update(c)
//
//	assert.Equal(t.T(), want, c.V)
//}
//
//func (t *GroupHandlerTest) TestUpdateGrpcErr() {
//	want := t.ServiceDownErr
//
//	srv := new(mock.ServiceMock)
//	srv.On("Update", t.Group.Id, t.GroupDto).Return(nil, t.ServiceDownErr)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return(t.Group.Id, nil)
//	c.On("GroupID").Return(t.Group.Id, nil)
//	c.On("Bind", &dto.GroupDto{}).Return(nil)
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Update(c)
//
//	assert.Equal(t.T(), want, c.V)
//}

//func (t *GroupHandlerTest) TestDeleteSuccess() {
//	srv := new(mock.ServiceMock)
//	srv.On("Delete", t.Group.Id).Return(true, nil)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return(t.Group.Id, nil)
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Delete(c)
//
//	assert.True(t.T(), c.V.(bool))
//}
//
//func (t *GroupHandlerTest) TestDeleteNotFound() {
//	want := t.NotFoundErr
//
//	srv := new(mock.ServiceMock)
//	srv.On("Delete", t.Group.Id).Return(false, t.NotFoundErr)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return(t.Group.Id, nil)
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Delete(c)
//
//	assert.Equal(t.T(), want, c.V)
//}
//
//func (t *GroupHandlerTest) TestDeleteInvalidID() {
//	want := &dto.ResponseErr{
//		StatusCode: http.StatusBadRequest,
//		Message:    "ID must be the uuid",
//	}
//
//	srv := new(mock.ServiceMock)
//	srv.On("Delete", t.Group.Id).Return(false, t.NotFoundErr)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return("", errors.New(want.Message))
//
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Delete(c)
//
//	assert.Equal(t.T(), want, c.V)
//}
//
//func (t *GroupHandlerTest) TestDeleteGrpcErr() {
//	want := t.ServiceDownErr
//
//	srv := new(mock.ServiceMock)
//	srv.On("Delete", t.Group.Id).Return(false, t.ServiceDownErr)
//
//	c := &mock.ContextMock{
//		Group:    t.Group,
//		GroupDto: t.GroupDto,
//	}
//	c.On("ID").Return(t.Group.Id, nil)
//	v, _ := validator.NewValidator()
//
//	h := NewHandler(srv, v)
//	h.Delete(c)
//
//	assert.Equal(t.T(), want, c.V)
//}
