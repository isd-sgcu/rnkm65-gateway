package router

import (
	"bytes"
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	guard "github.com/isd-sgcu/rnkm65-gateway/src/app/middleware/auth"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/utils"
	"github.com/isd-sgcu/rnkm65-gateway/src/config"
	"github.com/pkg/errors"
	"io"
)

type FiberRouter struct {
	*fiber.App
	user    fiber.Router
	auth    fiber.Router
	file    fiber.Router
	group   fiber.Router
	vaccine fiber.Router
	baan    fiber.Router
	qr      fiber.Router
	estamp  fiber.Router
}

type IGuard interface {
	Validate(ctx guard.IContext)
	CheckConfig(ctx guard.IContext)
}

func NewFiberRouter(authGuard IGuard, conf config.App) *FiberRouter {
	r := fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "RNKM65 API",
		BodyLimit:     conf.MaxFileSize * 1024 * 1024,
	})

	r.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
		}),
	)

	if conf.Debug {
		r.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/"
		}}))
		r.Get("/docs/*", swagger.HandlerDefault)
	}

	user := NewGroupRouteWithAuthMiddleware(r, "/user", authGuard.Validate, authGuard.CheckConfig)
	auth := NewGroupRouteWithAuthMiddleware(r, "/auth", authGuard.Validate, authGuard.CheckConfig)
	file := NewGroupRouteWithAuthMiddleware(r, "/file", authGuard.Validate, authGuard.CheckConfig)
	vaccine := NewGroupRouteWithAuthMiddleware(r, "/vaccine", authGuard.Validate, authGuard.CheckConfig)
	baan := NewGroupRouteWithAuthMiddleware(r, "/baan", authGuard.CheckConfig)
	group := NewGroupRouteWithAuthMiddleware(r, "/group", authGuard.Validate, authGuard.CheckConfig)
	qr := NewGroupRouteWithAuthMiddleware(r, "/qr", authGuard.Validate, authGuard.CheckConfig)
	estamp := NewGroupRouteWithAuthMiddleware(r, "/estamp", authGuard.Validate, authGuard.CheckConfig)

	return &FiberRouter{r, user, auth, file, group, vaccine, baan, qr, estamp}
}

func NewGroupRouteWithAuthMiddleware(r *fiber.App, path string, middlewares ...func(ctx guard.IContext)) fiber.Router {
	mList := createMiddlewaresList(middlewares)
	return r.Group(path, mList...)
}

func createMiddlewaresList(middlewares []func(ctx guard.IContext)) []func(ctx *fiber.Ctx) error {
	var result []func(ctx *fiber.Ctx) error
	for _, middleware := range middlewares {
		m := func(c *fiber.Ctx) error {
			middleware(NewFiberCtx(c))
			return nil
		}

		result = append(result, m)
	}

	return result
}

type FiberCtx struct {
	*fiber.Ctx
}

func (c *FiberCtx) UserID() string {
	return c.Ctx.Locals("UserId").(string)
}

func NewFiberCtx(c *fiber.Ctx) *FiberCtx {
	return &FiberCtx{c}
}

func (c *FiberCtx) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberCtx) JSON(statusCode int, v interface{}) {
	c.Ctx.Status(statusCode).JSON(v)
}

func (c *FiberCtx) ID() (id string, err error) {
	id = c.Params("id")

	_, err = uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *FiberCtx) Param(key string) (value string, err error) {
	value = c.Params(key)

	if key == "id" {
		_, err = uuid.Parse(value)
		if err != nil {
			return "", err
		}
	}
	return value, nil
}

func (c *FiberCtx) Token() string {
	return c.Ctx.Get(fiber.HeaderAuthorization, "")
}

func (c *FiberCtx) Method() string {
	return c.Ctx.Method()
}

func (c *FiberCtx) Path() string {
	return c.Ctx.Path()
}

func (c *FiberCtx) StoreValue(k string, v string) {
	c.Locals(k, v)
}

func (c *FiberCtx) Next() error {
	return c.Ctx.Next()
}

func (c *FiberCtx) Query(key string) string {
	return c.Ctx.Query(key)
}

func (c *FiberCtx) File(key string, allowContent map[string]struct{}, maxSize int64) (*dto.DecomposedFile, error) {
	file, err := c.Ctx.FormFile(key)
	if err != nil {
		return nil, err
	}

	if !utils.IsExisted(allowContent, file.Header["Content-Type"][0]) {
		return nil, errors.New("Not allow content")
	}

	if file.Size > maxSize {
		return nil, errors.New(fmt.Sprintf("Max file size is %v", maxSize))
	}
	content, err := file.Open()
	if err != nil {
		return nil, errors.New("Cannot read file")
	}

	defer content.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, content); err != nil {
		return nil, err
	}

	return &dto.DecomposedFile{
		Filename: file.Filename,
		Data:     buf.Bytes(),
	}, nil
}

func (c *FiberCtx) GetFormData(key string) string {
	return c.Ctx.FormValue(key)
}

func (c *FiberCtx) Host() string {
	return c.Ctx.Hostname()
}

func (c *FiberCtx) GetCTX() *fiber.Ctx {
	return c.Ctx
}
