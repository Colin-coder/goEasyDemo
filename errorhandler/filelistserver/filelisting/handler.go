package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var prefix = "/goEasyDemo/"

// 有错就返回出去
func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return errors.New("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		//return
		return err
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
