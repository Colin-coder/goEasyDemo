package main

import (
	"net/http"
	"os"

	"goEasyDemo/errorhandler/filelistserver/filelisting"

	"github.com/gpmgo/gopm/modules/log"
)

// appHandler 只做业务处理，未处理任何错误的函数，而是将error返回出来
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 函数式编程
// 统一的错误处理机制，将业务代码包装，将返回的错误进行统一处理
// errWrapper 将函数包装，返回一个新的函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	// 将handler 包装一把，返回一个对handler 返回的error做出处理的函数
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() { // 解决 panic 错误
			if r := recover(); r != nil {
				log.Print(1, "panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil { // 对返回的错误进行处理
			log.Warn("nkError: %s", err.Error())

			// 解决 user error
			// 这里是能做类型转换的userError
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			// 解决 system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err): // os.ErrNotExist
				code = http.StatusNotFound
			case os.IsPermission(err): // os.ErrPermission
				code = http.StatusForbidden
			default: // 未知的系统错误
				code = http.StatusInternalServerError
			}
			// 写入错误信息，错误字符串，错误码
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {

	// 第二个参数为处理HTTP请求的函数
	http.HandleFunc("/goEasyDemo/",
		/*func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/goEasyDemo/"):]
			file, err := os.Open(path)
			if err != nil {
				//panic(err)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}

			writer.Write(all)
		},*/
		errWrapper(filelisting.HandlerFileList),
	)

	// 开服务器，监听端口
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
