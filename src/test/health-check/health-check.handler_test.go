package health_check

import (
	"github.com/samithiwat/rnkm65-gateway/src/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HealthCheckHandlerTest struct {
	suite.Suite
}

func TestHealthCheckHandler(t *testing.T) {
	suite.Run(t, new(HealthCheckHandlerTest))
}

func (t *HealthCheckHandlerTest) TestCallHealthCheck() {
	want := map[string]interface{}{
		"Health": "OK!",
	}

	c := &ContextMock{}
	h := handler.NewHealthCheckHandler()

	h.HealthCheck(c)

	assert.Equal(t.T(), want, c.V)
}
