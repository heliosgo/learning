package errx

import "fmt"

type CodeError struct {
	errCode uint32
	errMsg  string
}

func (err *CodeError) GetCode() uint32 {
	return err.errCode
}

func (err *CodeError) GetMsg() string {
	return err.errMsg
}

func (err *CodeError) Error() string {
	return fmt.Sprintf("ErrCode: %d, ErrMsg: %s", err.errCode, err.errMsg)
}

func NewWithCode(code uint32) *CodeError {
	return &CodeError{errCode: code, errMsg: mapErrMsg(code)}
}

func NewWithMsg(msg string) *CodeError {
	return &CodeError{errCode: ServerCommonError, errMsg: msg}
}
