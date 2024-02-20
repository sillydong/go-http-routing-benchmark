// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var benchRe *regexp.Regexp

func isTested(name string) bool {
	if benchRe == nil {
		// Get -test.bench flag value (not accessible via flag package)
		bench := ""
		for _, arg := range os.Args {
			if strings.HasPrefix(arg, "-test.bench=") {
				// ignore the benchmark name after an underscore
				bench = strings.SplitN(arg[12:], "_", 2)[0]
				break
			}
		}

		// Compile RegExp to match Benchmark names
		var err error
		benchRe, err = regexp.Compile(bench)
		if err != nil {
			panic(err.Error())
		}
	}
	return benchRe.MatchString(name)
}

func calcMem(name string, load func()) {
	if !isTested(name) {
		return
	}

	m := new(runtime.MemStats)

	// before
	// force GC multiple times, since Go is using a generational GC
	// TODO: find a better approach
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	println("   "+name+":", after-before, "Bytes")
}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func benchRoutes(b *testing.B, router http.Handler, routes []route) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.method
			r.RequestURI = route.path
			u.Path = route.path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}

// Micro Benchmarks

// Route with Param (no write)
func BenchmarkChi_Param(b *testing.B) {
	router := loadChiSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param(b *testing.B) {
	router := loadDencoSingle("GET", "/user/:name", dencoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param(b *testing.B) {
	router := loadEchoSingle("GET", "/user/:name", echoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param(b *testing.B) {
	router := loadGocraftWebSingle("GET", "/user/:name", gocraftWebHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param(b *testing.B) {
	router := loadGojiSingle("GET", "/user/:name", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_Param(b *testing.B) {
	router := loadGojiv2Single("GET", "/user/:name", gojiv2Handler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", "/user/:name", httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkServeMux_Param(b *testing.B) {
	router := loadServeMuxSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// Route with 5 Params (no write)
const fiveColon = "/:a/:b/:c/:d/:e"
const fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
const fiveRoute = "/test/test/test/test/test"

func BenchmarkChi_Param5(b *testing.B) {
	router := loadChiSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param5(b *testing.B) {
	router := loadDencoSingle("GET", fiveColon, dencoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param5(b *testing.B) {
	router := loadEchoSingle("GET", fiveColon, echoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param5(b *testing.B) {
	router := loadGinSingle("GET", fiveColon, ginHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param5(b *testing.B) {
	router := loadGocraftWebSingle("GET", fiveColon, gocraftWebHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param5(b *testing.B) {
	router := loadGojiSingle("GET", fiveColon, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_Param5(b *testing.B) {
	router := loadGojiv2Single("GET", fiveColon, gojiv2Handler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param5(b *testing.B) {
	router := loadGorillaMuxSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param5(b *testing.B) {
	router := loadHttpRouterSingle("GET", fiveColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param5(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", fiveColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkServeMux_Param5(b *testing.B) {
	router := loadServeMuxSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}

// Route with 20 Params (no write)
const twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
const twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
const twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

func BenchmarkChi_Param20(b *testing.B) {
	router := loadChiSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param20(b *testing.B) {
	router := loadDencoSingle("GET", twentyColon, dencoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param20(b *testing.B) {
	router := loadEchoSingle("GET", twentyColon, echoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param20(b *testing.B) {
	router := loadGinSingle("GET", twentyColon, ginHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param20(b *testing.B) {
	router := loadGocraftWebSingle("GET", twentyColon, gocraftWebHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param20(b *testing.B) {
	router := loadGojiSingle("GET", twentyColon, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_Param20(b *testing.B) {
	router := loadGojiv2Single("GET", twentyColon, gojiv2Handler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param20(b *testing.B) {
	router := loadGorillaMuxSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param20(b *testing.B) {
	router := loadHttpRouterSingle("GET", twentyColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param20(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", twentyColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkServeMux_Param20(b *testing.B) {
	router := loadServeMuxSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}

// Route with Param and write
func BenchmarkChi_ParamWrite(b *testing.B) {
	router := loadChiSingle("GET", "/user/{name}", chiHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_ParamWrite(b *testing.B) {
	router := loadDencoSingle("GET", "/user/:name", dencoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_ParamWrite(b *testing.B) {
	router := loadEchoSingle("GET", "/user/:name", echoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_ParamWrite(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_ParamWrite(b *testing.B) {
	router := loadGocraftWebSingle("GET", "/user/:name", gocraftWebHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_ParamWrite(b *testing.B) {
	router := loadGojiSingle("GET", "/user/:name", gojiFuncWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_ParamWrite(b *testing.B) {
	router := loadGojiv2Single("GET", "/user/:name", gojiv2HandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_ParamWrite(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", gorillaHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_ParamWrite(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_ParamWrite(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", "/user/:name", httpTreeMuxHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkServeMux_ParamWrite(b *testing.B) {
	router := loadServeMuxSingle("GET", "/user/{name}", serveMuxHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
