package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Router[D, E any] interface {
		http.Handler
		Mount(base string, sub http.Handler)

		GET(u string, h func(EC[E]))
		POST(u string, h func(C[D, E]))
		PUT(u string, h func(C[D, E]))
		PATCH(u string, h func(C[D, E]))
		DELETE(u string, h func(EC[E]))
	}
)

type router[D Decoder[DT], E Encoder[ET], DT, ET any] struct {
	*chi.Mux
}

func NewRouter[D Decoder[DT], E Encoder[ET], DT, ET any]() Router[DT, ET] {
	return &router[D, E, DT, ET]{
		Mux: chi.NewRouter(),
	}
}

func (r *router[D, E, DT, ET]) GET(u string, h func(EC[ET])) {
	r.Get(u, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewContext[D, E](w, r))
	}))
}

func (r *router[D, E, DT, ET]) POST(u string, h func(C[DT, ET])) {
	r.Post(u, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewContext[D, E](w, r))
	}))
}

func (r *router[D, E, DT, ET]) PUT(u string, h func(C[DT, ET])) {
	r.Put(u, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewContext[D, E](w, r))
	}))
}

func (r *router[D, E, DT, ET]) PATCH(u string, h func(C[DT, ET])) {
	r.Patch(u, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewContext[D, E](w, r))
	}))
}

func (r *router[D, E, DT, ET]) DELETE(u string, h func(EC[ET])) {
	r.Delete(u, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewContext[D, E](w, r))
	}))
}
