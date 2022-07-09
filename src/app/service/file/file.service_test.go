package file

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/constant"
	mock "github.com/isd-sgcu/rnkm65-gateway/src/mocks/file"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type FileServiceTest struct {
	suite.Suite
	url            string
	userId         string
	fileDecomposed *dto.DecomposedFile
	ServiceDownErr *dto.ResponseErr
}

func TestFileService(t *testing.T) {
	suite.Run(t, new(FileServiceTest))
}

func (t *FileServiceTest) SetupTest() {
	t.url = faker.URL()
	t.userId = faker.UUIDDigit()

	t.fileDecomposed = &dto.DecomposedFile{
		Filename: faker.Word(),
		Data:     []byte("Hello"),
	}

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

func (t *FileServiceTest) TestUploadImageSuccess() {
	want := t.url

	c := mock.ClientMock{}
	c.On("UploadImage", &proto.UploadImageRequest{
		Filename: t.fileDecomposed.Filename, Data: t.fileDecomposed.Data, Tag: 1, UserId: t.userId}).Return(&proto.UploadImageResponse{Url: t.url}, nil)

	srv := NewService(&c)

	actual, err := srv.UploadImage(t.fileDecomposed, t.userId, constant.Profile)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *FileServiceTest) TestUploadImageFailed() {
	want := t.ServiceDownErr

	c := mock.ClientMock{}
	c.On("UploadImage", &proto.UploadImageRequest{
		Filename: t.fileDecomposed.Filename, Data: t.fileDecomposed.Data, Tag: 1, UserId: t.userId}).Return(nil, errors.New("Cannot connect to service"))

	srv := NewService(&c)

	actual, err := srv.UploadImage(t.fileDecomposed, t.userId, constant.Profile)

	assert.Equal(t.T(), "", actual)
	assert.Equal(t.T(), want, err)
}
