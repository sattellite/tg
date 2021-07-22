// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/seniorGolang/tg/example/implement"
	"github.com/seniorGolang/tg/example/interfaces"
)

type httpUser struct {
	log          zerolog.Logger
	errorHandler ErrorHandler
	svc          *serverUser
	base         interfaces.User
}

func NewUser(log zerolog.Logger, svcUser interfaces.User) (srv *httpUser) {

	srv = &httpUser{
		base: svcUser,
		log:  log,
		svc:  newServerUser(svcUser),
	}
	return
}

func (http httpUser) Service() MiddlewareSetUser {
	return http.svc
}

func (http *httpUser) WithLog(log zerolog.Logger) *httpUser {
	http.svc.WithLog(log)
	return http
}

func (http *httpUser) WithTrace() *httpUser {
	http.svc.WithTrace()
	return http
}

func (http *httpUser) WithMetrics() *httpUser {
	http.svc.WithMetrics()
	return http
}

func (http *httpUser) WithErrorHandler(handler ErrorHandler) *httpUser {
	http.errorHandler = handler
	return http
}

func (http *httpUser) SetRoutes(route *fiber.App) {
	route.Get("/api/v2/user/info", http.serveGetUser)
	route.Post("/api/v2/user/file", http.serveUploadFile)
	route.Patch("/api/v2/user/custom/response", http.serveCustomResponse)
	route.Delete("/api/v2/user/custom", func(ctx *fiber.Ctx) (err error) {
		return implement.CustomHandler(ctx, http.base)
	})
}
