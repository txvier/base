package common

import (
	bt "github.com/txvier/base/time"
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

func NilParseJSONTime(arg interface{}) (j bt.JSONTime) {
	if arg == nil {
		return j
	}
	return bt.JSONTime(time.Unix(0, arg.(int64)))
}
