// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Extra
var extraAPI = []route{
	{"GET", "/1/users/me"},
	{"GET", "/1/users/:objectId"},
	// {"PUT", "/1/users/:objectId"},
}

var (
	extraDenco      http.Handler
	extraEcho       http.Handler
	extraGin        http.Handler
	extraGorillaMux http.Handler
	// extraHttpRouter  http.Handler
	extraHttpTreeMux http.Handler
	extraServeMux    http.Handler
)

func init() {
	println("#ExtraAPI Routes:", len(extraAPI))

	calcMem("Denco", func() {
		extraDenco = loadDenco(extraAPI)
	})
	calcMem("Echo", func() {
		extraEcho = loadEcho(extraAPI)
	})
	calcMem("Gin", func() {
		extraGin = loadGin(extraAPI)
	})
	calcMem("GorillaMux", func() {
		extraGorillaMux = loadGorillaMux(extraAPI)
	})
	// calcMem("HttpRouter", func() {
	// 	extraHttpRouter = loadHttpRouter(extraAPI)
	// })
	calcMem("HttpTreeMux", func() {
		extraHttpTreeMux = loadHttpTreeMux(extraAPI)
	})
	calcMem("ServeMux", func() {
		extraServeMux = loadServeMux(extraAPI)
	})

	println()
}

// Static
func BenchmarkDenco_ExtraStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, extraDenco, req)
}
func BenchmarkEcho_ExtraStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, extraEcho, req)
}
func BenchmarkGin_ExtraStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, extraGin, req)
}
func BenchmarkGorillaMux_ExtraStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, extraGorillaMux, req)
}

//	func BenchmarkHttpRouter_ExtraStatic(b *testing.B) {
//		req, _ := http.NewRequest("GET", "/1/users", nil)
//		benchRequest(b, extraHttpRouter, req)
//	}
func BenchmarkHttpTreeMux_ExtraStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, extraHttpTreeMux, req)
}
func BenchmarkServeMux_ExtraStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, extraServeMux, req)
}

// One Param
func BenchmarkDenco_ExtraParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, extraDenco, req)
}
func BenchmarkEcho_ExtraParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, extraEcho, req)
}
func BenchmarkGin_ExtraParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, extraGin, req)
}
func BenchmarkGorillaMux_ExtraParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, extraGorillaMux, req)
}

//	func BenchmarkHttpRouter_ExtraParam(b *testing.B) {
//		req, _ := http.NewRequest("GET", "/1/classes/go", nil)
//		benchRequest(b, extraHttpRouter, req)
//	}
func BenchmarkHttpTreeMux_ExtraParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, extraHttpTreeMux, req)
}
func BenchmarkServeMux_ExtraParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, extraServeMux, req)
}

// Two Params
func BenchmarkDenco_Extra2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, extraDenco, req)
}
func BenchmarkEcho_Extra2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, extraEcho, req)
}
func BenchmarkGin_Extra2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, extraGin, req)
}
func BenchmarkGorillaMux_Extra2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, extraGorillaMux, req)
}

//	func BenchmarkHttpRouter_Extra2Params(b *testing.B) {
//		req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
//		benchRequest(b, extraHttpRouter, req)
//	}
func BenchmarkHttpTreeMux_Extra2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, extraHttpTreeMux, req)
}
func BenchmarkServeMux_Extra2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, extraServeMux, req)
}

// All routes
func BenchmarkDenco_ExtraAll(b *testing.B) {
	benchRoutes(b, extraDenco, extraAPI)
}
func BenchmarkEcho_ExtraAll(b *testing.B) {
	benchRoutes(b, extraEcho, extraAPI)
}
func BenchmarkGin_ExtraAll(b *testing.B) {
	benchRoutes(b, extraGin, extraAPI)
}

//	func BenchmarkHttpRouter_ExtraAll(b *testing.B) {
//		benchRoutes(b, extraHttpRouter, extraAPI)
//	}
func BenchmarkHttpTreeMux_ExtraAll(b *testing.B) {
	benchRoutes(b, extraHttpTreeMux, extraAPI)
}
func BenchmarkServeMux_ExtraAll(b *testing.B) {
	benchRoutes(b, extraServeMux, extraAPI)
}
