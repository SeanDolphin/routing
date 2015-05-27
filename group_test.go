package routing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestGrouping(t *testing.T) {
	for _, pathinfo := range [][]string{
		[]string{"/prefix", "/path"},
		[]string{"/long", "/thisisaverylongpathnamethatshouldnotbeusedinmostcases"},
		[]string{"/trailing", "/path/"},
		[]string{"/multi", "/one/two/three"},
		[]string{"/direct", "/"},
	} {
		testRoute(t, pathinfo[0], pathinfo[1])
	}
}

func testRoute(t *testing.T, prefix, path string) {
	router := mux.NewRouter()
	Group(router, prefix, func(r *mux.Router) {
		r.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {})
	})

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", prefix, path), nil)
	if err != nil {
		t.Fail()
	}

	var match mux.RouteMatch
	found := router.Match(req, &match)
	if !found {
		t.Fail()
	}
}
