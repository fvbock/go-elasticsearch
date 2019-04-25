package esutil

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSONReader is an utility function which encodes v into JSON and returns it as a reader.
//
func JSONReader(v interface{}) io.Reader {
	return &jsonReader{val: v, buf: nil}
}

// JSONEncoder defines the interface for custom JSON encoders.
//
type JSONEncoder interface {
	EncodeJSON(io.Writer) error
}

type jsonReader struct {
	val interface{}
	buf io.ReadWriter
}

func (r *jsonReader) Read(p []byte) (int, error) {
	if r.buf == nil {
		r.buf = new(bytes.Buffer)
		if err := r.encode(); err != nil {
			return 0, err
		}
	}
	return r.buf.Read(p)
}

func (r *jsonReader) encode() error {
	var err error

	if e, ok := r.val.(JSONEncoder); ok {
		err = e.EncodeJSON(r.buf)
		if err != nil {
			return err
		}
		return nil
	}

	return json.NewEncoder(r.buf).Encode(r.val)
}
