// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	// If you add new routers please:
	// - Keep the benchmark functions etc. alphabetically sorted
	// - Make a pull request (without benchmark results) at
	//   https://github.com/julienschmidt/go-http-routing-benchmark

	"github.com/dimfeld/httptreemux"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"
	"github.com/naoina/denco"
	_ "github.com/naoina/kocha-urlrouter/doublearray"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

var nullLogger *log.Logger

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initGin()
}

// Common
// func httpHandlerFunc(_ http.ResponseWriter, _ *http.Request) {}

// func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, r.RequestURI)
// }

// new servmux
// func serveMuxHandlerWrite(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, r.PathValue("name"))
// }

func loadServeMux(routes []route) http.Handler {
	h := Wrap(EmptyController)
	if loadTestHandler {
		h = Wrap(UrlController)
	}

	re := regexp.MustCompile(":([^/]*)")

	mux := http.NewServeMux()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "{$1}")

		mux.HandleFunc(fmt.Sprintf("%s %s", route.method, path), h)
	}

	return mux
}

func loadServeMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc(fmt.Sprintf("%s %s", method, path), handler)
	return router
}

// chi
// func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, chi.URLParam(r, "name"))
// }

func loadChi(routes []route) http.Handler {
	h := Wrap(EmptyController)
	if loadTestHandler {
		h = Wrap(UrlController)
	}

	re := regexp.MustCompile(":([^/]*)")

	mux := chi.NewRouter()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "{$1}")

		switch route.method {
		case "GET":
			mux.Get(path, h)
		case "POST":
			mux.Post(path, h)
		case "PUT":
			mux.Put(path, h)
		case "PATCH":
			mux.Patch(path, h)
		case "DELETE":
			mux.Delete(path, h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return mux
}

func loadChiSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := chi.NewRouter()
	switch method {
	case "GET":
		mux.Get(path, handler)
	case "POST":
		mux.Post(path, handler)
	case "PUT":
		mux.Put(path, handler)
	case "PATCH":
		mux.Patch(path, handler)
	case "DELETE":
		mux.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return mux
}

// Denco
// func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {}

// func dencoHandlerWrite(w http.ResponseWriter, r *http.Request, params denco.Params) {
// 	io.WriteString(w, params.Get("name"))
// }

// func dencoHandlerTest(w http.ResponseWriter, r *http.Request, params denco.Params) {
// 	io.WriteString(w, r.RequestURI)
// }

func loadDenco(routes []route) http.Handler {
	h := WrapDenco(EmptyController)
	if loadTestHandler {
		h = WrapDenco(UrlController)
	}

	mux := denco.NewMux()
	handlers := make([]denco.Handler, 0, len(routes))
	for _, route := range routes {
		handler := mux.Handler(route.method, route.path, h)
		handlers = append(handlers, handler)
	}
	handler, err := mux.Build(handlers)
	if err != nil {
		panic(err)
	}
	return handler
}

func loadDencoSingle(method, path string, h denco.HandlerFunc) http.Handler {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{mux.Handler(method, path, h)})
	if err != nil {
		panic(err)
	}
	return handler
}

// Echo
// func echoHandler(c echo.Context) error {
// 	return nil
// }

// func echoHandlerWrite(c echo.Context) error {
// 	io.WriteString(c.Response(), c.Param("name"))
// 	return nil
// }

// func echoHandlerTest(c echo.Context) error {
// 	io.WriteString(c.Response(), c.Request().RequestURI)
// 	return nil
// }

func loadEcho(routes []route) http.Handler {
	var h echo.HandlerFunc = WrapEcho(EmptyController)
	if loadTestHandler {
		h = WrapEcho(UrlController)
	}

	e := echo.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			e.GET(r.path, h)
		case "POST":
			e.POST(r.path, h)
		case "PUT":
			e.PUT(r.path, h)
		case "PATCH":
			e.PATCH(r.path, h)
		case "DELETE":
			e.DELETE(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return e
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	case "POST":
		e.POST(path, h)
	case "PUT":
		e.PUT(path, h)
	case "PATCH":
		e.PATCH(path, h)
	case "DELETE":
		e.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return e
}

// Gin
// func ginHandle(_ *gin.Context) {}

// func ginHandleWrite(c *gin.Context) {
// 	io.WriteString(c.Writer, c.Params.ByName("name"))
// }

// func ginHandleTest(c *gin.Context) {
// 	io.WriteString(c.Writer, c.Request.RequestURI)
// }

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []route) http.Handler {
	h := WrapGin(EmptyController)
	if loadTestHandler {
		h = WrapGin(UrlController)
	}

	router := gin.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, handle)
	return router
}

// gorilla/mux
// func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	io.WriteString(w, params["name"])
// }

func loadGorillaMux(routes []route) http.Handler {
	h := Wrap(EmptyController)
	if loadTestHandler {
		h = Wrap(UrlController)
	}

	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.path, "{$1}"),
			h,
		).Methods(route.method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// HttpRouter
// func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

// func httpRouterHandleWrite(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
// 	io.WriteString(w, ps.ByName("name"))
// }

// func httpRouterHandleTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	io.WriteString(w, r.RequestURI)
// }

func loadHttpRouter(routes []route) http.Handler {
	h := WrapHttpRouter(EmptyController)
	if loadTestHandler {
		h = WrapHttpRouter(UrlController)
	}

	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

// httpTreeMux
// func httpTreeMuxHandler(_ http.ResponseWriter, _ *http.Request, _ map[string]string) {}

// func httpTreeMuxHandlerWrite(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
// 	io.WriteString(w, vars["name"])
// }

// func httpTreeMuxHandlerTest(w http.ResponseWriter, r *http.Request, _ map[string]string) {
// 	io.WriteString(w, r.RequestURI)
// }

func loadHttpTreeMux(routes []route) http.Handler {
	h := WrapHttpTreeMux(EmptyController)
	if loadTestHandler {
		h = WrapHttpTreeMux(UrlController)
	}

	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpTreeMuxSingle(method, path string, handler httptreemux.HandlerFunc) http.Handler {
	router := httptreemux.New()
	router.Handle(method, path, handler)
	return router
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
