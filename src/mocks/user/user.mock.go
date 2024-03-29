package user

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

func (s *ServiceMock) FindOne(id string) (result *proto.User, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Create(in *dto.UserDto) (result *proto.User, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Verify(id string, verifyType string) (result bool, err *dto.ResponseErr) {
	args := s.Called(id, verifyType)

	if args.Get(0) != nil {
		result = args.Bool(0)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Update(id string, in *dto.UpdateUserDto) (result *proto.User, err *dto.ResponseErr) {
	args := s.Called(id, in)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) CreateOrUpdate(in *dto.UserDto) (result *proto.User, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		result = args.Get(0).(*proto.User)
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

func (s *ServiceMock) GetUserEstamp(id string) (res *proto.GetUserEstampResponse, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.GetUserEstampResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) ConfirmEstamp(uid, eid string) (res *proto.ConfirmEstampResponse, err *dto.ResponseErr) {
	args := s.Called(uid, eid)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.ConfirmEstampResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) Verify(_ context.Context, in *proto.VerifyUserRequest, _ ...grpc.CallOption) (res *proto.VerifyUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.VerifyUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) FindOne(_ context.Context, in *proto.FindOneUserRequest, _ ...grpc.CallOption) (res *proto.FindOneUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindOneUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) FindByStudentID(_ context.Context, in *proto.FindByStudentIDUserRequest, _ ...grpc.CallOption) (res *proto.FindByStudentIDUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindByStudentIDUserResponse)
	}

	return res, args.Error(1)

}

func (c *ClientMock) Create(_ context.Context, in *proto.CreateUserRequest, _ ...grpc.CallOption) (res *proto.CreateUserResponse, err error) {
	args := c.Called(in.User)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Update(_ context.Context, in *proto.UpdateUserRequest, _ ...grpc.CallOption) (res *proto.UpdateUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UpdateUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *proto.DeleteUserRequest, _ ...grpc.CallOption) (res *proto.DeleteUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) CreateOrUpdate(_ context.Context, in *proto.CreateOrUpdateUserRequest, _ ...grpc.CallOption) (res *proto.CreateOrUpdateUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateOrUpdateUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) ConfirmEstamp(_ context.Context, in *proto.ConfirmEstampRequest, _ ...grpc.CallOption) (res *proto.ConfirmEstampResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.ConfirmEstampResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) GetUserEstamp(_ context.Context, in *proto.GetUserEstampRequest, _ ...grpc.CallOption) (res *proto.GetUserEstampResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.GetUserEstampResponse)
	}

	return res, args.Error(1)
}

type ContextMock struct {
	mock.Mock
	V      interface{}
	Status int
}

func (c *ContextMock) JSON(status int, v interface{}) {
	c.V = v
	c.Status = status
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	if args.Get(0) != nil {
		switch v.(type) {
		case *dto.UserDto:
			*v.(*dto.UserDto) = *args.Get(0).(*dto.UserDto)
		case *dto.UpdateUserDto:
			*v.(*dto.UpdateUserDto) = *args.Get(0).(*dto.UpdateUserDto)
		case *dto.Verify:
			*v.(*dto.Verify) = *args.Get(0).(*dto.Verify)
		case *dto.ConfirmEstampRequest:
			*v.(*dto.ConfirmEstampRequest) = *args.Get(0).(*dto.ConfirmEstampRequest)
		}
	}

	return args.Error(1)
}

func (c *ContextMock) ID() (string, error) {
	args := c.Called()

	return args.String(0), args.Error(1)
}

func (c *ContextMock) Host() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) UserID() string {
	args := c.Called()
	return args.String(0)
}

func (c *ContextMock) Query(key string) string {
	args := c.Called(key)
	return args.String(0)
}
