package norepeatsubstring

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func add(c, b int) int {
	return c + b
}

func LengthOfNonRepeatedSubString(s string) int {
	lastOccur := make(map[rune]int)
	length := 0
	start := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccur[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > length {
			length = i - start + 1
		}
		lastOccur[ch] = i
	}

	return length
}

func LengthOfNonRepeatedSubString2(s string) int {
	lastOccur := make([]int, 0xffff)
	length := 0
	start := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccur[ch]; lastI > start {
			start = lastI
		}
		if i-start+1 > length {
			length = i - start + 1
		}
		lastOccur[ch] = i + 1
	}

	return length
}

func testChineseChar() {
	s := "adb我爱中国"

	fmt.Println(utf8.RuneCountInString(s))
	// 获取字节
	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}

	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// 进行转换
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

	fmt.Println()

}

func runCurrentFile() {
	fmt.Println("hello")

	fmt.Println("new")
	fmt.Println(89)

	fmt.Println(add(1, 2))

	fmt.Println(LengthOfNonRepeatedSubString("a"))
	fmt.Println(LengthOfNonRepeatedSubString("sdfad"))
	fmt.Println(LengthOfNonRepeatedSubString("安慰法司法所"))
	fmt.Println(LengthOfNonRepeatedSubString("adb违法"))

	testChineseChar()

	fmt.Println(strings.Fields("sadf asdfa ddd aaa"))

	//tree.TestTree()
}
