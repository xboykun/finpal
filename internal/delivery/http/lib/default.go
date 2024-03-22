package lib

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/justinas/alice"
)

const (
	_Created       = http.StatusCreated
	_Success       = http.StatusOK
	_Error         = http.StatusBadRequest
	_InternalError = http.StatusInternalServerError

	_DecodeErrorMessage  = "error decode request"
	_DecodeErrorTemplate = `{"code":"%s","message":"%s", "data":null}`
	_EncodeErrorMessage  = "error encode response"
)

var (
	_CreatedMessage       string = http.StatusText(http.StatusCreated)
	_SuccessMessage       string = http.StatusText(http.StatusOK)
	_ErrorMessage         string = http.StatusText(http.StatusBadRequest)
	_InternalErrorMessage string = http.StatusText(http.StatusInternalServerError)
)

var (
	HttpMiddlewareChain = alice.New

	DefaultInternalError = func(rw http.ResponseWriter, code, msg *string) {
		if len(*code) <= 0 {
			*code = strconv.FormatInt(int64(_InternalError), 10)
		}

		if len(*msg) <= 0 {
			*msg = _EncodeErrorMessage
		}

		rw.Write([]byte(fmt.Sprintf(_DecodeErrorTemplate, *code, *msg)))
	}

	DefaultMiddlewareError = func(rw http.ResponseWriter, r *http.Request, err error) {
		NewHttpRes[any]().SetMessage(_DecodeErrorMessage).Error().JSON(rw)
	}
)
