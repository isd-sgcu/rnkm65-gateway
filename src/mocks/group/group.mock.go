package group

import (
	"context"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindByToken(token string) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Create(id string) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Update(in *dto.GroupDto, id string) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(in, id)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Join(token string, userId string, isLeader bool, members int) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(token, userId, isLeader, members)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) DeleteMember(userId string, leaderId string) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(userId, leaderId)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Leave(userId string) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(userId)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) FindByToken(_ context.Context, in *proto.FindByTokenGroupRequest, _ ...grpc.CallOption) (res *proto.FindByTokenGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindByTokenGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Create(_ context.Context, in *proto.CreateGroupRequest, _ ...grpc.CallOption) (res *proto.CreateGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Update(_ context.Context, in *proto.UpdateGroupRequest, _ ...grpc.CallOption) (res *proto.UpdateGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UpdateGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Join(_ context.Context, in *proto.JoinGroupRequest, _ ...grpc.CallOption) (res *proto.JoinGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.JoinGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) DeleteMember(_ context.Context, in *proto.DeleteMemberGroupRequest, _ ...grpc.CallOption) (res *proto.DeleteMemberGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteMemberGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Leave(_ context.Context, in *proto.LeaveGroupRequest, _ ...grpc.CallOption) (res *proto.LeaveGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.LeaveGroupResponse)
	}

	return res, args.Error(1)
}

type ContextMock struct {
	mock.Mock
	V interface{}
}

func (c *ContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	if args.Get(0) != nil {
		switch v.(type) {
		case *dto.JoinGroupRequest:
			*v.(*dto.JoinGroupRequest) = *args.Get(0).(*dto.JoinGroupRequest)
		case *dto.GroupDto:
			*v.(*dto.GroupDto) = *args.Get(0).(*dto.GroupDto)
		}
	}

	return args.Error(1)
}

func (c *ContextMock) ID() (string, error) {
	args := c.Called()

	return args.String(0), args.Error(1)
}

func (c *ContextMock) GroupID() string {
	args := c.Called()
	return args.String(0)
}

func (c *ContextMock) Param(string) (string, error) {
	args := c.Called()

	return args.String(0), args.Error(1)
}

func (c *ContextMock) UserID() string {
	args := c.Called()
	return args.String(0)
}
