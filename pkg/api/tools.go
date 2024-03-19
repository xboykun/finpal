package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/ggicci/httpin"
	"github.com/justinas/alice"
)

var (
	HandleChain = alice.New
)

func ParseInput[T any]() func(http.Handler) http.Handler {
	var d T
	return httpin.NewInput(d)
}

func TakeInput[T any](r *http.Request) *T {
	return r.Context().Value(httpin.Input).(*T)
}

const (
	_DefaultCode_Created       = http.StatusCreated
	_DefaultCode_Success       = http.StatusOK
	_DefaultCode_Error         = http.StatusBadRequest
	_DefaultCode_InternalError = http.StatusInternalServerError
)

type (
	Res[T any] struct {
		Status bool   `json:"status"`
		Code   string `json:"code"`
		Meta   *Meta  `json:"meta,omitempty"`
		Msg    string `json:"message"`
		Data   T      `json:"data"`
	}
)

func NewResponse[T any]() *Res[T] {
	return &Res[T]{}
}

func (r *Res[T]) SetStatus(status bool) *Res[T] {
	if r == nil {
		return nil
	}
	r.Status = status
	return r
}

func (r *Res[T]) SetMeta(meta *Meta) *Res[T] {
	if r == nil {
		return nil
	}
	r.Meta = meta
	return r
}

func (r *Res[T]) SetCode(code string) *Res[T] {
	if r == nil {
		return nil
	}
	r.Code = strings.TrimSpace(code)
	return r
}

func (r *Res[T]) SetCodeInt(code int) *Res[T] {
	if r == nil {
		return nil
	}
	r.Code = strconv.FormatInt(int64(code), 10)
	return r
}

func (r *Res[T]) SetMessage(msg string) *Res[T] {
	if r == nil {
		return nil
	}
	r.Msg = strings.TrimSpace(msg)
	return r
}

func (r *Res[T]) SetData(d T) *Res[T] {
	if r == nil {
		return nil
	}
	r.Data = d
	return r
}

func (r *Res[T]) Created() *Res[T] {
	if r == nil {
		return nil
	}
	r.SetStatus(true)
	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_DefaultCode_Created, 10))
	}
	return r
}

func (r *Res[T]) Success() *Res[T] {
	if r == nil {
		return nil
	}
	r.SetStatus(true)
	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_DefaultCode_Success, 10))
	}
	return r
}

func (r *Res[T]) Error() *Res[T] {
	if r == nil {
		return nil
	}
	r.SetStatus(false)
	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_DefaultCode_Error, 10))
	}
	return r
}

func (r *Res[T]) InternalError() *Res[T] {
	if r == nil {
		return nil
	}
	r.SetStatus(false)
	if len(r.Code) <= 0 {
		r.SetCode(strconv.FormatInt(_DefaultCode_InternalError, 10))
	}
	return r
}

func (r *Res[T]) JSON(w io.Writer) error {
	if r == nil {
		return errors.New("response is nil")
	}
	return json.NewEncoder(w).Encode(r)
}

type (
	Meta struct {
		Page         int64 `json:"page"`
		PerPage      int64 `json:"perPage"`
		TotalPages   int64 `json:"totalPages"`
		TotalRecords int64 `json:"totalRecords"`
	}
)

func NewMeta() *Meta {
	return &Meta{}
}

func (m *Meta) SetPage(n int) *Meta {
	m.Page = int64(n)
	return m
}

func (m *Meta) SetPerPage(n int) *Meta {
	m.PerPage = int64(n)
	return m
}

func (m *Meta) SetTotalPages(n int) *Meta {
	m.TotalPages = int64(n)
	return m
}

func (m *Meta) SetTotalRecord(n int64) *Meta {
	m.TotalRecords = n
	return m
}
