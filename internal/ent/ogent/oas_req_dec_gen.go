// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeCreateUserRequest(r *http.Request, span trace.Span) (req CreateUserReq, err error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request CreateUserReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode CreateUser:application/json request")
		}
		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func decodeUpdateUserRequest(r *http.Request, span trace.Span) (req UpdateUserReq, err error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request UpdateUserReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode UpdateUser:application/json request")
		}
		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}
