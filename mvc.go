package main

import (
	"net/http"
)

// 组合设计模式
// 简易的MVC框架

// Filter, 可以添加到特定的Handler中，
// 只对该handler做处理
type Filter interface {
	BeforeServeHTTP(http.ResponseWriter, *http.Request) bool
	AfterServeHTTP(http.ResponseWriter, *http.Request) bool
}

// 基础的Http Handler，
// 加入对Filter的支持
type BaseHandler struct {
	Filters []Filter
	Handle  func(w http.ResponseWriter, r *http.Request) error
}

// http处理入口
func (b BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// handle all the error
		err_ := recover()
		if err_ == nil {
			return
		}

		var stack string
		var buf bytes.Buffer
		buf.Write(debug.Stack())
		stack = buf.String()

		log.Error(fmt.Sprintf("%v", err_), "\n", stack)
		return
	}()

	if b.Filters != nil {
		for _, filter := range b.Filters {
			if filter.BeforeServeHTTP(w, r) == false {
				return
			}
		}
	}

	b.Handle(w, r)

	if b.Filters != nil {
		for _, filter := range b.Filters {
			if filter.AfterServeHTTP(w, r) == false {
				return
			}
		}
	}
}

// 好吧，总结下来还是不好
// 毛大师的Repositories github.com/tiancaiamao/middleware 好好看看吧
