package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func errPanic(writer http.ResponseWriter, req *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, req *http.Request) error {
	return testingUserError("user error")
}
func errNotFound(writer http.ResponseWriter, req *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, req *http.Request) error {
	return os.ErrPermission
}

func errUnKnown(writer http.ResponseWriter, req *http.Request) error {
	return errors.New("unknown error")
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	// 函数，返回错误code，错误信息字符串
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNoPermission, 403, "Forbidden"},
	{errNotFound, 404, "Not Found"},
	{errUnKnown, 500, "Internal Server Error"},
}

// 未开服务器，只是测试了一下代码
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {

		// 需要测试的函数 errWrapper
		f := errWrapper(tt.h) // 先包装一下

		// 假的返回值
		resp := httptest.NewRecorder()

		// 假的请求，，get 请求
		req := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com", nil)

		f(resp, req)

		verifyResponse(resp.Result(), tt.code, tt.message, t)
	}
}

// 真实启动了一个服务器，进行测试
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))

		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

// 验证结果是否正确的函数
func verifyResponse(resp *http.Response,
	expectedCode int,
	expectedMsg string,
	t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d %s); "+"got (%d %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
