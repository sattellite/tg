// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	maxParallelBatch = 100
	// Version defines the version of the JSON RPC implementation
	Version = "2.0"
	// contentTypeJson defines the content type to be served
	contentTypeJson = "application/json"
	// ParseError defines invalid JSON was received by the server
	// An error occurred on the server while parsing the JSON text
	parseError = -32700
	// InvalidRequestError defines the JSON sent is not a valid Request object
	invalidRequestError = -32600
	// MethodNotFoundError defines the method does not exist / is not available
	methodNotFoundError = -32601
	// InvalidParamsError defines invalid method parameter(s)
	invalidParamsError = -32602
	// InternalError defines a server error
	internalError = -32603
)

type idJsonRPC = json.RawMessage

type baseJsonRPC struct {
	ID      idJsonRPC       `json:"id"`
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method,omitempty"`
	Error   *errorJsonRPC   `json:"error,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type errorJsonRPC struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err errorJsonRPC) Error() string {
	return err.Message
}

type jsonrpcResponses []baseJsonRPC

func (responses *jsonrpcResponses) append(response *baseJsonRPC) {
	if response == nil {
		return
	}
	if response.ID != nil {
		*responses = append(*responses, *response)
	}
}

type methodJsonRPC func(span opentracing.Span, ctx *fiber.Ctx, requestBase baseJsonRPC) (responseBase *baseJsonRPC)

func (srv *Server) serveBatch(ctx *fiber.Ctx) (err error) {
	batchSpan := extractSpan(srv.log, fmt.Sprintf("jsonRPC:%s", ctx.Path()), ctx)
	defer injectSpan(srv.log, batchSpan, ctx)
	defer batchSpan.Finish()
	methodHTTP := ctx.Method()
	if methodHTTP != fiber.MethodPost {
		ext.Error.Set(batchSpan, true)
		batchSpan.SetTag("msg", "only POST method supported")
		ctx.Response().SetStatusCode(fiber.StatusMethodNotAllowed)
		if _, err = ctx.WriteString("only POST method supported"); err != nil {
			return
		}
		return
	}
	if value := ctx.Context().Value(CtxCancelRequest); value != nil {
		return
	}
	ctx.Response().Header.SetContentType(contentTypeJson)
	var requests []baseJsonRPC
	if err = json.Unmarshal(ctx.Body(), &requests); err != nil {
		ext.Error.Set(batchSpan, true)
		batchSpan.SetTag("msg", "request body could not be decoded: "+err.Error())
		sendResponse(srv.log, ctx, makeErrorResponseJsonRPC([]byte("\"0\""), parseError, "request body could not be decoded: "+err.Error(), nil))
		return
	}
	responses := make(jsonrpcResponses, 0, len(requests))
	var n int
	var wg sync.WaitGroup
	for _, request := range requests {
		methodNameOrigin := request.Method
		method := strings.ToLower(request.Method)
		switch method {

		case "jsonrpc.test":
			wg.Add(1)
			go func(request baseJsonRPC) {

				span := opentracing.StartSpan(request.Method, opentracing.ChildOf(batchSpan.Context()))
				span.SetTag("batch", true)
				defer span.Finish()
				if request.ID != nil {
					responses.append(srv.httpJsonRPC.test(span, ctx, request))
					wg.Done()
					return
				}
				srv.httpJsonRPC.test(span, ctx, request)
				wg.Done()
			}(request)
		default:
			span := opentracing.StartSpan(request.Method, opentracing.ChildOf(batchSpan.Context()))
			span.SetTag("batch", true)
			ext.Error.Set(span, true)
			span.SetTag("msg", "invalid method '"+methodNameOrigin+"'")
			responses.append(makeErrorResponseJsonRPC(request.ID, methodNotFoundError, "invalid method '"+methodNameOrigin+"'", nil))
			span.Finish()
		}
		if n > maxParallelBatch {
			n = 0
			wg.Wait()
		}
		n++
	}
	wg.Wait()
	sendResponse(srv.log, ctx, responses)
	return
}

func makeErrorResponseJsonRPC(id idJsonRPC, code int, msg string, data interface{}) *baseJsonRPC {

	if id == nil {
		return nil
	}

	return &baseJsonRPC{
		Error: &errorJsonRPC{
			Code:    code,
			Data:    data,
			Message: msg,
		},
		ID:      id,
		Version: Version,
	}
}
