package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/isd-sgcu/rnkm65-gateway/src/app/handler/user"
	"github.com/isd-sgcu/rnkm65-gateway/src/config"
	mock "github.com/isd-sgcu/rnkm65-gateway/src/mocks/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QrRouterTest struct {
	suite.Suite
}

func TestQrRouter(t *testing.T) {
	suite.Run(t, new(QrRouterTest))
}

func (t *QrRouterTest) TestPostQrRouter() {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "POST /qr/checkin status 200",
			route:        "/qr/checkin",
			expectedCode: http.StatusOK,
		},
	}

	g := mock.GuardMock{}
	conf := config.App{
		Port:        3000,
		Debug:       true,
		MaxFileSize: 1000000,
	}

	r := NewFiberRouter(&g, conf)

	r.DeleteUser("/", func(ctx user.IContext) {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "Hello World",
		})
	})

	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, nil)
		res, _ := r.Test(req, 1)

		assert.Equal(t.T(), test.expectedCode, res.StatusCode, test.description)
	}
}
