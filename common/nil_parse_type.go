package common

import (
	"github.com/txvier/base/txtime"
	"time"
)

func NilParseString(arg interface{}) (s string) {
	if arg == nil {
		return s
	}
	return arg.(string)
}

func NilParseInt(arg interface{}) (i int) {
	if arg == nil {
		return i
	}
	return arg.(int)
}

func NilParseInt64(arg interface{}) (i int64) {
	if arg == nil {
		return i
	}
	return arg.(int64)
}

func NilParseJSONTime(arg interface{}) (j txtime.JSONTime) {
	if arg == nil {
		return j
	}
	return txtime.JSONTime(time.Unix(0, arg.(int64)))
}
