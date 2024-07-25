package router

import (
	"encoding/json"
	"io"
)

type Decoder[D any] interface {
	Decode(r io.Reader) (D, error)
}

type Encoder[E any] interface {
	Encode(w io.Writer, value E) error
}

type None struct{}

func (None) Decode(_ io.Reader) (None, error) {
	return None{}, nil
}

func (None) Encode(_ io.Writer, _ None) error {
	return nil
}

type JSON[T any] struct{}

func (j JSON[T]) Decode(r io.Reader) (T, error) {
	var t T
	err := json.NewDecoder(r).Decode(&t)

	return t, err
}

func (j JSON[T]) Encode(w io.Writer, value T) error {
	return json.NewEncoder(w).Encode(value)
}
