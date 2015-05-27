# routing

[![GoDoc](https://godoc.org/github.com/SeanDolphin/routing?status.png)](http://godoc.org/github.com/SeanDolphin/routing)  
 [![Build Status](https://travis-ci.org/SeanDolphin/routing.svg?branch=master)](https://travis-ci.org/SeanDolphin/routing)  
[![Coverage Status](https://coveralls.io/repos/SeanDolphin/routing/badge.svg?branch=master)](https://coveralls.io/r/SeanDolphin/routing?branch=master)  
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

routing is a convenience wrapper for gorrila mux routing to make grouping routes easier.  It makes it act more like martini.

## Installation

The import path for the package is *github.com/SeanDolphin/routing*.

To install it, run:

    go get github.com/SeanDolphin/routing

## Usage

~~~ go
package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routing.Group(router, "/meta", func(r *mux.Router) {
		r.Methods("GET").Path("/").HandlerFunc(sendOK)
		routing.Group(r, "/watch", func(r *mux.Router) {
			r.Methods("POST").Path("/doc/").Handler(sendOK)
			r.Methods("POST").Path("/student/").Handler(sendOK)
			r.Methods("POST").Path("/product/").Handler(sendOK)
		})

		routing.Group(r, "/update", func(r *mux.Router) {
			r.Methods("POST").Path("/doc/").Handler(sendOK)
		})

		routing.Group(r, "/tree", func(r *mux.Router) {
			r.Path("/{key}/takeownership/{userKey}/{newKey}/").Handler(sendOK)
		})

		routing.Group(r, "/graph/{key}", func(r *mux.Router) {
			r.Methods("POST").Path("/tokens/").Handler(sendOK)
			r.Methods("POST").Path("/rename/").Handler(sendOK)
			r.Methods("POST").Path("/tokens/cache/").Handler(sendOK)
		})
	})

	n := negroni.New()
	n.UseHandler(router)
	http.Handle("/", n)
}


func sendOK(writer http.ResponseWriter, req *http.Request) {
	json.NewEncoder(writer).Encode(map[string]int{"status": http.StatusOK})
}
~~~