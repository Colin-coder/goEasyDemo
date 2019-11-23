package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com@abc.com
dadff sadfaf gfaf@aa.com
fawfe agrg kkk@ll.xom
dsfaio iwef ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)

	for _, m := range match {
		fmt.Println(m)
	}
}
