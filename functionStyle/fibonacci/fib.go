package fib

import (
	"bufio"
	"fmt"
	"io"
)

// 1, 1, 2, 3, 5, 8 ...
//    a, b
//       a, b
// ...
func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b

		fmt.Println(a)
		return a
	}
}

type intGen func() int

func (intGen) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//func main() {
//	f := Fibonacci()
//	f()
//	f()
//	f()
//	f()
//	f()
//	f()
//	f()
//	f()
//	f()
//	f()
//}
