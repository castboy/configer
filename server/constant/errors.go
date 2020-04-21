package constant

import (
	"fmt"
	"runtime"
)

type MyErr struct {
	code     ErrCode
	funcName string
	v        interface{}
}

type ErrCode uint

const (
	SystemErr = 0
	TradeErr  = 1
	DBErr     = 2
	ArgsErr   = 3
	AssertErr = 4
	UnknowErr = 5
)

var ErrMsg = map[ErrCode]string{
	SystemErr: "System Error",
	TradeErr:  "Trade Error",
	DBErr:     "DB Error",
	ArgsErr:   "Args Err",
}

func NewErr(code ErrCode, v interface{}) *MyErr {
	my := &MyErr{
		code: code,
		v:    v,
	}
	my.FuncName()
	return my
}

func (me *MyErr) Error() string {
	return fmt.Sprintf("funcName:%v,ErrorCode%s, Interface %v", me.funcName, ErrMsg[me.code], me.v)
}

func (me *MyErr) FuncName() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	me.funcName = runtime.FuncForPC(pc[0]).Name()
}
func IsDBErr(err error) bool {
	v, ok := err.(*MyErr)
	if ok && v.v == DBErr {
		return true
	}

	return false
}

func IsArgsErr(err error) bool {
	v, ok := err.(*MyErr)
	if ok && v.v == ArgsErr {
		return true
	}

	return false
}
