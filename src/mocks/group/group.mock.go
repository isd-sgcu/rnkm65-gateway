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

func (s *ServiceMock) FindOne(id string) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
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

func (s *ServiceMock) Create(in *dto.GroupDto) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Update(id string, in *dto.GroupDto) (result *proto.Group, err *dto.ResponseErr) {
	args := s.Called(id, in)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.Group)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Delete(id string) (result bool, err *dto.ResponseErr) {
	args := s.Called(id)

	result = args.Bool(0)

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) FindOne(_ context.Context, in *proto.FindOneGroupRequest, _ ...grpc.CallOption) (res *proto.FindOneGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindOneGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) FindByToken(_ context.Context, in *proto.FindByTokenGroupRequest, _ ...grpc.CallOption) (res *proto.FindByTokenGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindByTokenGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Create(_ context.Context, in *proto.CreateGroupRequest, _ ...grpc.CallOption) (res *proto.CreateGroupResponse, err error) {
	args := c.Called(in.Group)

	if args.Get(0) != nil {
		res = args.Get((0)).(*proto.CreateGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Update(_ context.Context, in *proto.UpdateGroupRequest, _ ...grpc.CallOption) (res *proto.UpdateGroupResponse, err error) {
	args := c.Called(in.Group)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UpdateGroupResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *proto.DeleteGroupRequest, _ ...grpc.CallOption) (res *proto.DeleteGroupResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteGroupResponse)
	}

	return res, args.Error(1)
}

type ContextMock struct {
	mock.Mock
	V        interface{}
	Group    *proto.Group
	GroupDto *dto.GroupDto
}

func (c *ContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	*v.(*dto.GroupDto) = *c.GroupDto

	return args.Error(0)
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
