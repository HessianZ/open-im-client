package openim

import (
	"fmt"
)

type OpenIMError interface {
	GetErrCode() int
	GetErrMsg() string
	GetErrDetail() string
	Error() string
}
type openIMErrorSt struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	ErrDlt  string `json:"errDlt"`
}

type OpenIMResponse[T any] struct {
	openIMErrorSt
	Data T `json:"data"`
}

func (r *OpenIMResponse[T]) GetErrCode() int {
	return r.openIMErrorSt.ErrCode
}
func (r *OpenIMResponse[T]) GetErrMsg() string {
	return r.openIMErrorSt.ErrMsg
}
func (r *OpenIMResponse[T]) GetErrDetail() string {
	return r.openIMErrorSt.ErrDlt
}

func (e *OpenIMResponse[T]) Error() string {
	return fmt.Sprintf("openim error[%d] %s - %s", e.ErrCode, e.ErrMsg, e.ErrDlt)
}
