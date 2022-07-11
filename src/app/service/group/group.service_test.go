package group

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/mocks/group"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"testing"
)

type GroupServiceTest struct {
	suite.Suite
	Group          *proto.Group
	GroupReq       *proto.Group
	GroupDto       *dto.GroupDto
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestGroupService(t *testing.T) {
	suite.Run(t, new(GroupServiceTest))
}

func (t *GroupServiceTest) SetupTest() {
	t.Group = &proto.Group{
		Id:       faker.UUIDDigit(),
		LeaderID: faker.Word(),
		Token:    faker.Word(),
	}

	t.GroupReq = &proto.Group{
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
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
}

func (t *GroupServiceTest) TestFindByTokenSuccess() {
	want := t.Group

	c := &group.ClientMock{}
	c.On("FindByToken", &proto.FindByTokenGroupRequest{Token: t.Group.Token}).Return(&proto.FindByTokenGroupResponse{Group: want}, nil)

	srv := NewService(c)
	actual, err := srv.FindByToken(t.Group.Token)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestFindByTokenNotFound() {
	want := t.NotFoundErr

	c := &group.ClientMock{}
	c.On("FindByToken", &proto.FindByTokenGroupRequest{Token: t.Group.Token}).Return(nil, status.Error(codes.NotFound, "Group not found"))

	srv := NewService(c)

	actual, err := srv.FindByToken(t.Group.Token)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *GroupServiceTest) TestFindByTokenGrpcErr() {
	want := t.ServiceDownErr

	c := &group.ClientMock{}
	c.On("FindByToken", &proto.FindByTokenGroupRequest{Token: t.Group.Token}).Return(nil, errors.New("Server is down"))

	srv := NewService(c)

	actual, err := srv.FindByToken(t.Group.Token)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *GroupServiceTest) TestCreateSuccess() {
	want := t.Group

	c := &group.ClientMock{}
	c.On("Create", t.GroupReq).Return(&proto.CreateGroupResponse{Group: want}, nil)

	srv := NewService(c)

	actual, err := srv.Create(t.GroupDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestCreateGrpcErr() {
	want := t.ServiceDownErr

	c := &group.ClientMock{}
	c.On("Create", t.GroupReq).Return(nil, errors.New("Service is down"))

	srv := NewService(c)

	actual, err := srv.Create(t.GroupDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *GroupServiceTest) TestUpdateSuccess() {
	want := t.Group

	c := &group.ClientMock{}
	c.On("Update", t.Group).Return(&proto.UpdateGroupResponse{Group: want}, nil)

	srv := NewService(c)

	actual, err := srv.Update(t.Group.Id, t.GroupDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestUpdateNotFound() {
	want := t.NotFoundErr

	c := &group.ClientMock{}
	c.On("Update", t.Group).Return(nil, status.Error(codes.NotFound, "Group not found"))

	srv := NewService(c)

	actual, err := srv.Update(t.Group.Id, t.GroupDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *GroupServiceTest) TestUpdateGrpcErr() {
	want := t.ServiceDownErr

	c := &group.ClientMock{}
	c.On("Update", t.Group).Return(nil, errors.New("Service is down"))

	srv := NewService(c)

	actual, err := srv.Update(t.Group.Id, t.GroupDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *GroupServiceTest) TestDeleteSuccess() {
	c := &group.ClientMock{}
	c.On("Delete", &proto.DeleteGroupRequest{Id: t.Group.Id}).Return(&proto.DeleteGroupResponse{Success: true}, nil)

	srv := NewService(c)

	actual, err := srv.Delete(t.Group.Id)

	assert.Nil(t.T(), err)
	assert.True(t.T(), actual)
}

func (t *GroupServiceTest) TestDeleteNotFound() {
	want := t.NotFoundErr

	c := &group.ClientMock{}
	c.On("Delete", &proto.DeleteGroupRequest{Id: t.Group.Id}).Return(nil, status.Error(codes.NotFound, "Group not found"))

	srv := NewService(c)

	actual, err := srv.Delete(t.Group.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *GroupServiceTest) TestDeleteGrpcErr() {
	want := t.ServiceDownErr

	c := &group.ClientMock{}
	c.On("Delete", &proto.DeleteGroupRequest{Id: t.Group.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)

	actual, err := srv.Delete(t.Group.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}
