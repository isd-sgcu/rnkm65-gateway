package auth

import (
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/handler/auth"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/utils"
	"github.com/isd-sgcu/rnkm65-gateway/src/config"
	"net/http"
	"strings"
)

type Guard struct {
	service     auth.IService
	excludes    map[string]struct{}
	allowPhases map[string][]string
	conf        config.App
	isValidate  bool
}

type IContext interface {
	Token() string
	Method() string
	Path() string
	StoreValue(string, string)
	JSON(int, interface{})
	Next() error
}

func NewAuthGuard(s auth.IService, e map[string]struct{}, allowPhase map[string][]string, conf config.App) Guard {
	return Guard{
		service:     s,
		excludes:    e,
		allowPhases: allowPhase,
		conf:        conf,
		isValidate:  true,
	}
}

func (m *Guard) Validate(ctx IContext) {
	method := ctx.Method()
	path := ctx.Path()

	ids := utils.FindIDFromPath(path)

	path = utils.FormatPath(method, path, ids)
	if utils.IsExisted(m.excludes, path) {
		ctx.Next()
		return
	}

	token := ctx.Token()
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, &dto.ResponseErr{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid token",
		})
		m.isValidate = false
		return
	}

	payload, errRes := m.service.Validate(token)
	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		m.isValidate = false
		return
	}

	ctx.StoreValue("UserId", payload.UserId)
	ctx.Next()
}

func (m *Guard) CheckConfig(ctx IContext) {
	method := ctx.Method()
	path := ctx.Path()

	//check whether there is a token in path
	//if token exist, replace token with ":token"
	pathSlice := strings.Split(path, "/")
	//paths which can have a token is "/group/token"
	if pathSlice[1] == "group" {
		if len(pathSlice) > 2 && pathSlice[2] != "members" && pathSlice[2] != "leave" {
			token := pathSlice[2]
			path = strings.Replace(path, token, ":token", 1)
		}
	}

	ids := utils.FindIDFromPath(path)

	path = utils.FormatPath(method, path, ids)

	if utils.IsExisted(m.excludes, path) {
		ctx.Next()
		return
	}

	phses, ok := m.allowPhases[path]
	if !ok {
		ctx.Next()
		return
	}

	currentPhase := m.conf.Phase
	for _, phs := range phses {
		if phs == currentPhase {
			ctx.Next()
			return
		}
	}

	ctx.JSON(http.StatusForbidden, &dto.ResponseErr{
		StatusCode: http.StatusForbidden,
		Message:    "Forbidden Resource",
		Data:       nil,
	})
	m.isValidate = false
}
