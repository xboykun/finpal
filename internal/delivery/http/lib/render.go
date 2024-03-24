package lib

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"nami/pkg/response"
)

type (
	HttpRes[T any] interface {
		// Setup
		HttpSetup[T]

		// Builder :
		HttpBuilder[T]

		// Encoder
		HttpEncoder[T]
	}

	// Setup Response
	HttpSetup[T any] interface {
		SetMeta(meta *response.Meta) HttpRes[T]
		SetCode(code string) HttpRes[T]
		SetMessage(msg string) HttpRes[T]
		SetData(d T) HttpRes[T]
	}

	// Http Builder Response
	HttpBuilder[T any] interface {
		// 201
		Created() HttpRes[T]

		// 200
		Success() HttpRes[T]

		// 400
		Error() HttpRes[T]

		// 500
		InternalError() HttpRes[T]

		// HTTP Status Code
		Status(code int) HttpRes[T]
	}

	HttpEncoder[T any] interface {
		JSON(rw http.ResponseWriter) error
	}

	HttpResImp[T any] struct {
		*response.Res[T]
		httpStatusCode int `json:"-"`
	}
)

func NewHttpRes[T any]() HttpRes[T] {
	return &HttpResImp[T]{
		Res: new(response.Res[T]),
	}
}

func (r *HttpResImp[T]) SetMeta(meta *response.Meta) HttpRes[T] {
	r.Res.SetMeta(meta)
	return r
}

func (r *HttpResImp[T]) SetCode(code string) HttpRes[T] {
	r.Res.SetCode(code)
	return r
}

func (r *HttpResImp[T]) SetMessage(msg string) HttpRes[T] {
	r.Res.SetMessage(msg)
	return r
}

func (r *HttpResImp[T]) SetData(d T) HttpRes[T] {
	r.Res.SetData(d)
	return r
}

func (r *HttpResImp[T]) Status(code int) HttpRes[T] {
	r.httpStatusCode = code
	return r
}

func (r *HttpResImp[T]) Created() HttpRes[T] {
	if r == nil {
		return nil
	}

	if len(r.Msg) <= 0 {
		r.SetMessage(_CreatedMessage)
	}

	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_Created, 10))
	}

	if r.httpStatusCode <= 0 {
		r.Status(_Created)
	}

	return r
}

func (r *HttpResImp[T]) Success() HttpRes[T] {
	if r == nil {
		return nil
	}

	if len(r.Msg) <= 0 {
		r.SetMessage(_SuccessMessage)
	}

	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_Success, 10))
	}

	if r.httpStatusCode <= 0 {
		r.Status(_Success)
	}

	return r
}

func (r *HttpResImp[T]) Error() HttpRes[T] {
	if r == nil {
		return nil
	}

	if len(r.Msg) <= 0 {
		r.SetMessage(_ErrorMessage)
	}

	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_Error, 10))
	}

	if r.httpStatusCode <= 0 {
		r.Status(_Error)
	}

	return r
}

func (r *HttpResImp[T]) InternalError() HttpRes[T] {
	if r == nil {
		return nil
	}

	if len(r.Msg) <= 0 {
		r.SetMessage(_InternalErrorMessage)
	}

	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_InternalError, 10))
	}

	if r.httpStatusCode <= 0 {
		r.Status(_InternalError)
	}

	return r
}

func (r *HttpResImp[T]) JSON(rw http.ResponseWriter) error {
	if r == nil {
		var (
			code    = strconv.FormatInt(_InternalError, 10)
			message = _InternalErrorMessage
		)

		DefaultInternalError(rw, &code, &message)
		return errors.New("response is nil")
	}

	rw.Header().Set("Accept", "application/json")
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(r.httpStatusCode)

	if err := json.NewEncoder(rw).Encode(r); err != nil {
		DefaultInternalError(rw, &r.Code, &r.Msg)
		return err
	}

	return nil
}
