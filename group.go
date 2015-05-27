package routing

import "github.com/gorilla/mux"

//Group takes a mux.Router and creates a sub route
func Group(r *mux.Router, path string, f func(*mux.Router)) {
	f(r.PathPrefix(path).Subrouter())
}
