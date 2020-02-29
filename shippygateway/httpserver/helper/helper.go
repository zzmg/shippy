package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var (
	AllErrors = make(map[int]string)
	ErrOK     = newErr(20000, "OK")
)

type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type Err struct {
	Code    int
	Message string
}

func newErr(code int, msg string) *Err {
	if _, ok := AllErrors[code]; ok {
		panic("Duplicated error code!!!")
	}
	AllErrors[code] = msg
	return &Err{code, msg}
}

func SuccessResponse(ctx echo.Context, data interface{}) error {
	var buffer bytes.Buffer
	m := json.NewEncoder(&buffer)
	if err := m.Encode(data); err != nil {
		return ErrorResponse(ctx, errors.New("Illegal JSON"))
	}
	return ctx.JSON(http.StatusOK, Payload(buffer.Bytes()))
}

func ErrorResponse(ctx echo.Context, err error) error {
	if err != nil {
		return ErrorResponseWithMessage(ctx, err.Error())
	}
	return ErrorResponseWithMessage(ctx, "")
}

func ErrorResponseWithMessage(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusOK, Payload([]byte("{}"), "50000", message))
}

func Payload(data json.RawMessage, fields ...string) *Response {
	res := &Response{
		Code:    ErrOK.Code,
		Message: ErrOK.Message,
		Data:    data,
	}

	length := len(fields)
	if length >= 1 {
		singleCode, _ := strconv.Atoi(fields[0])
		res.Code = int(singleCode)
	}
	if length >= 2 {
		res.Message = fields[1]
	}

	return res
}
