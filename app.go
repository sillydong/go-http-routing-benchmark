package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"
	"github.com/naoina/denco"
)

type ReqContext interface {
	context.Context

	GetRequest() *http.Request
	Get(string) string
}

type PathValue func(string) string

type reqContext struct {
	context.Context

	req *http.Request
	val PathValue
}

func (ctx *reqContext) GetRequest() *http.Request {
	return ctx.req
}

func (ctx *reqContext) Get(name string) string {
	return ctx.val(name)
}

func newReqContext(ctx context.Context, req *http.Request, val PathValue) (ReqContext, error) {
	rctx := &reqContext{
		Context: ctx,
		req:     req,
		val:     val,
	}
	return rctx, nil
}

type Controller func(ctx ReqContext) (interface{}, error)

func EmptyController(ctx ReqContext) (interface{}, error) {
	return "", nil
}

func UrlController(ctx ReqContext) (interface{}, error) {
	return ctx.GetRequest().RequestURI, nil
}

func NameController(ctx ReqContext) (interface{}, error) {
	name := ctx.Get("name")
	return name, nil
}

func Wrap(c Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rctx, err := newReqContext(r.Context(), r, func(name string) string { return r.PathValue(name) })
		if err != nil {
			w.WriteHeader(500)
			return
		}
		res, err := c(rctx)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(res.(string))
		w.WriteHeader(200)
	}
}

func WrapDenco(c Controller) denco.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p denco.Params) {
		rctx, err := newReqContext(r.Context(), r, func(s string) string { return p.Get(s) })
		if err != nil {
			w.WriteHeader(500)
			return
		}
		res, err := c(rctx)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(res.(string))
		w.WriteHeader(200)
	}
}

func WrapEcho(c Controller) echo.HandlerFunc {
	return func(cc echo.Context) error {
		rctx, err := newReqContext(cc.Request().Context(), cc.Request(), func(s string) string { return cc.Param(s) })
		if err != nil {
			cc.Response().WriteHeader(500)
			return nil
		}
		res, err := c(rctx)
		if err != nil {
			cc.Response().WriteHeader(500)
			return nil
		}
		cc.JSON(200, res.(string))
		return nil
	}
}

func WrapGin(c Controller) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rctx, err := newReqContext(ctx.Request.Context(), ctx.Request, func(s string) string { return ctx.Param(s) })
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}
		res, err := c(rctx)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}
		json.NewEncoder(ctx.Writer).Encode(res.(string))
		ctx.Writer.WriteHeader(200)
		// ctx.JSON(200, res.(string))
	}
}

func WrapGorrilaMux(c Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rctx, err := newReqContext(r.Context(), r, func(name string) string { return mux.Vars(r)[name] })
		if err != nil {
			w.WriteHeader(500)
			return
		}
		res, err := c(rctx)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(res.(string))
		w.WriteHeader(200)
	}
}

func WrapHttpRouter(c Controller) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		rctx, err := newReqContext(r.Context(), r, func(s string) string { return p.ByName(s) })
		if err != nil {
			w.WriteHeader(500)
			return
		}
		res, err := c(rctx)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(res.(string))
		w.WriteHeader(200)
	}
}

func WrapHttpTreeMux(c Controller) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, m map[string]string) {
		rctx, err := newReqContext(r.Context(), r, func(s string) string { return m[s] })
		if err != nil {
			w.WriteHeader(500)
			return
		}
		res, err := c(rctx)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(res.(string))
		w.WriteHeader(200)
	}
}
