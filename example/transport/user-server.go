// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/seniorGolang/tg/example/interfaces"
	"github.com/seniorGolang/tg/example/interfaces/types"
)

type serverUser struct {
	svc            interfaces.User
	getUser        UserGetUser
	uploadFile     UserUploadFile
	customResponse UserCustomResponse
	customHandler  UserCustomHandler
}

type MiddlewareSetUser interface {
	Wrap(m MiddlewareUser)
	WrapGetUser(m MiddlewareUserGetUser)
	WrapUploadFile(m MiddlewareUserUploadFile)
	WrapCustomResponse(m MiddlewareUserCustomResponse)
	WrapCustomHandler(m MiddlewareUserCustomHandler)

	WithTrace()
	WithMetrics()
	WithLog(log zerolog.Logger)
}

func newServerUser(svc interfaces.User) *serverUser {
	return &serverUser{
		customHandler:  svc.CustomHandler,
		customResponse: svc.CustomResponse,
		getUser:        svc.GetUser,
		svc:            svc,
		uploadFile:     svc.UploadFile,
	}
}

func (srv *serverUser) Wrap(m MiddlewareUser) {
	srv.svc = m(srv.svc)
	srv.getUser = srv.svc.GetUser
	srv.uploadFile = srv.svc.UploadFile
	srv.customResponse = srv.svc.CustomResponse
	srv.customHandler = srv.svc.CustomHandler
}

func (srv *serverUser) GetUser(ctx context.Context, cookie string, userAgent string) (user *types.User, err error) {
	return srv.getUser(ctx, cookie, userAgent)
}

func (srv *serverUser) UploadFile(ctx context.Context, fileBytes []byte) (err error) {
	return srv.uploadFile(ctx, fileBytes)
}

func (srv *serverUser) CustomResponse(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (err error) {
	return srv.customResponse(ctx, arg0, arg1, opts...)
}

func (srv *serverUser) CustomHandler(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (err error) {
	return srv.customHandler(ctx, arg0, arg1, opts...)
}

func (srv *serverUser) WrapGetUser(m MiddlewareUserGetUser) {
	srv.getUser = m(srv.getUser)
}

func (srv *serverUser) WrapUploadFile(m MiddlewareUserUploadFile) {
	srv.uploadFile = m(srv.uploadFile)
}

func (srv *serverUser) WrapCustomResponse(m MiddlewareUserCustomResponse) {
	srv.customResponse = m(srv.customResponse)
}

func (srv *serverUser) WrapCustomHandler(m MiddlewareUserCustomHandler) {
	srv.customHandler = m(srv.customHandler)
}

func (srv *serverUser) WithTrace() {
	srv.Wrap(traceMiddlewareUser)
}

func (srv *serverUser) WithMetrics() {
	srv.Wrap(metricsMiddlewareUser)
}

func (srv *serverUser) WithLog(log zerolog.Logger) {
	srv.Wrap(loggerMiddlewareUser(log))
}
