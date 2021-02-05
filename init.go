package gojieba

import (
	"github.com/xiafei114/gojieba/deps/cppjieba"
	"github.com/xiafei114/gojieba/deps/limonp"
	"github.com/xiafei114/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
