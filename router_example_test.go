package router_test

import (
	"net/http"

	"github.com/robinknaapen/router"
)

func ExampleNewRouter() {
	r := router.NewRouter[router.JSON[string], router.JSON[string]]()
	r.GET(`/`, func(c router.EncodeContext[string]) {
		c.Response().WriteHeader(http.StatusOK)
	})

	r.Mount(`/admin`, admin())
	http.ListenAndServe(`:8080`, r)
}

type Admin struct {
	IsAdmin bool `json:"is_admin"`
}

func admin() http.Handler {
	r := router.NewRouter[router.JSON[Admin], router.None]()
	r.GET(`/`, func(c router.EncodeContext[router.None]) {
		c.Response().WriteHeader(http.StatusOK)
	})
	r.POST(`/`, func(c router.Context[Admin, router.None]) {
		c.Response().WriteHeader(http.StatusOK)
	})

	r.Mount(`/user`, adminUsers())
	return r
}

func adminUsers() http.Handler {
	r := router.NewRouter[router.None, router.None]()
	r.GET(`/`, func(c router.EncodeContext[router.None]) {
		c.Response().WriteHeader(http.StatusOK)
	})

	return r
}
