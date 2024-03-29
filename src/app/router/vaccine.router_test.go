package router

import (
	"github.com/isd-sgcu/rnkm65-gateway/src/app/handler/vaccine"
	"github.com/isd-sgcu/rnkm65-gateway/src/config"
	mock "github.com/isd-sgcu/rnkm65-gateway/src/mocks/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type VaccineRouterTest struct {
	suite.Suite
}

func TestVaccineRouter(t *testing.T) {
	suite.Run(t, new(VaccineRouterTest))
}

func (t *VaccineRouterTest) TestPostVaccineRouter() {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "POST /vaccine status 200",
			route:        "/vaccine/verify",
			expectedCode: http.StatusOK,
		},
		{
			description:  "POST HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: http.StatusNotFound,
		},
	}

	g := mock.GuardMock{}
	conf := config.App{
		Port:        3000,
		Debug:       true,
		MaxFileSize: 1000000,
	}

	r := NewFiberRouter(&g, conf)

	r.PostVaccine("/verify", func(ctx vaccine.IContext) {
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
