package router

import "net/http"

type (
	EC[E any] interface {
		Request() *http.Request
		Response() http.ResponseWriter

		Encode(E) error
	}

	C[D, E any] interface {
		Request() *http.Request
		Response() http.ResponseWriter

		Decode() (D, error)
		Encode(E) error
	}
)

type context[D Decoder[DT], E Encoder[ET], DT, ET any] struct {
	w http.ResponseWriter
	r *http.Request

	d D
	e E
}

func NewContext[D Decoder[DT], E Encoder[ET], DT, ET any](w http.ResponseWriter, r *http.Request) C[DT, ET] {
	return &context[D, E, DT, ET]{
		w: w,
		r: r,
	}
}

func (c *context[_, _, _, _]) Request() *http.Request {
	return c.r
}

func (c *context[_, _, _, _]) Response() http.ResponseWriter {
	return c.w
}

func (c *context[_, _, D, _]) Decode() (D, error) {
	return c.d.Decode(c.r.Body)
}

func (c *context[_, _, _, E]) Encode(e E) error {
	return c.e.Encode(c.w, e)
}
