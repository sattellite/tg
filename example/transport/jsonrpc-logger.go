// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/seniorGolang/dumper/viewer"

	"github.com/seniorGolang/tg/example/interfaces"
)

type loggerJsonRPC struct {
	next interfaces.JsonRPC
	log  zerolog.Logger
}

func loggerMiddlewareJsonRPC(log zerolog.Logger) MiddlewareJsonRPC {
	return func(next interfaces.JsonRPC) interfaces.JsonRPC {
		return &loggerJsonRPC{
			log:  log,
			next: next,
		}
	}
}

func (m loggerJsonRPC) Test(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (ret1 int, ret2 string, err error) {
	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method": "test",
			"request": viewer.Sprintf("%+v", requestJsonRPCTest{
				Arg0: arg0,
				Arg1: arg1,
				Opts: opts,
			}),
			"response": viewer.Sprintf("%+v", responseJsonRPCTest{
				Ret1: ret1,
				Ret2: ret2,
			}),
			"service": "JsonRPC",
			"took":    time.Since(begin),
		}
		if ctx.Value(headerRequestID) != nil {
			fields["requestID"] = ctx.Value(headerRequestID)
		}
		if err != nil {
			m.log.Error().Err(err).Fields(fields).Msg("call test")
			return
		}
		m.log.Info().Fields(fields).Msg("call test")
	}(time.Now())
	return m.next.Test(ctx, arg0, arg1, opts...)
}
