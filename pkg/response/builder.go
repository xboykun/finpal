package response

import (
	"strings"
)

type (
	Res[T any] struct {
		Code string `json:"code"`
		Meta *Meta  `json:"meta,omitempty"`
		Msg  string `json:"message"`
		Data T      `json:"data"`
	}

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
