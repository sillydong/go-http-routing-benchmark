package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	// load functions of all routers
	routers = []struct {
		name string
		load func(routes []route) http.Handler
	}{
		{"Chi", loadChi},
		{"Denco", loadDenco},
		{"Echo", loadEcho},
		{"Gin", loadGin},
		{"HttpRouter", loadHttpRouter},
		{"HttpTreeMux", loadHttpTreeMux},
		{"ServeMux", loadServeMux},
	}

	// all APIs
	apis = []struct {
		name   string
		routes []route
	}{
		{"GitHub", githubAPI},
		{"GPlus", gplusAPI},
		{"Parse", parseAPI},
		{"Static", staticRoutes},
		// {"Extra", extraAPI},
	}
)

func TestRouters(t *testing.T) {
	loadTestHandler = true

	for _, router := range routers {
		req, _ := http.NewRequest("GET", "/", nil)
		u := req.URL
		rq := u.RawQuery

		for _, api := range apis {
			r := router.load(api.routes)

			for _, route := range api.routes {
				w := httptest.NewRecorder()
				req.Method = route.method
				req.RequestURI = route.path
				u.Path = route.path
				u.RawQuery = rq
				r.ServeHTTP(w, req)
				if w.Code != 200 || strings.TrimSpace(w.Body.String()) != `"`+route.path+`"` {
					t.Errorf(
						"%s in API %s: %d - `%s`; expected `%s`\n",
						router.name, api.name, w.Code, strings.TrimSpace(w.Body.String()), route.path,
					)
				}
			}
		}
	}

	loadTestHandler = false
}
