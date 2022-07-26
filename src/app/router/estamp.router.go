package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/handler/estamp"
)

func (r *FiberRouter) PostEstamp(path string, h func(ctx estamp.IContext)) {
	r.estamp.Post(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}
