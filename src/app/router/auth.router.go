package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/handler/auth"
)

func (r *FiberRouter) GetAuth(path string, h func(ctx auth.IContext)) {
	r.auth.Get(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PostAuth(path string, h func(ctx auth.IContext)) {
	r.auth.Post(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}
