package env

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router enables defining API routes and assinging handlers.
type Router struct {
	router *mux.Router
}

// NewRouter is a constructor for a Router.
func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}

// Get enables defining route and handler for GET method.
func (r *Router) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(path, f).Methods("GET")
}

// Post enables defining route and handler for POST method.
func (r *Router) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(path, f).Methods("POST")
}

// Put enables defining route and handler for PUT method.
func (r *Router) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(path, f).Methods("PUT")
}

// Delete enables defining route and handler for DELETE method.
func (r *Router) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(path, f).Methods("DELETE")
}

func (r *Router) InternalRouter() *mux.Router {
	return r.router
}
