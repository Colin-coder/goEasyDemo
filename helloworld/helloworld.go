package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"runtime"
)

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb)
}

func readFile() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abd")
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling %s with %d, %d \n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(), a, b)
	return op(a, b)
}

func sums(numbers ...int) int {
	ret := 0
	for i := range numbers {
		ret += numbers[i]
	}
	return ret
}

// go语言函数入参只有值传递，即只是拷贝一份传入
func swap(a, b *int) {
	*b, *a = *a, *b
}

func arraysTest() {
	var arr1 = [...]int{1, 2, 5, 6}
	for i, v := range arr1 {
		fmt.Println(i, v)
	}
}

func updateSlice(slicein []int) {
	slicein[0] = 100
}
func sliceTest() {
	sliceav := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//updateSlice(sliceav[1:])

	s1 := sliceav[2:6]
	fmt.Println(s1)
	fmt.Println("After slice1 update:")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(sliceav)

	s2 := sliceav[5:7]
	fmt.Println("After slice2 update:")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(sliceav)
}

func resliceTest() {
	resliceTotal := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(resliceTotal)

	s1 := resliceTotal[3:5]
	fmt.Println(s1)
	s2 := s1[2:4]
	fmt.Println(s2)

	s3 := append(s1, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println(resliceTotal)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))
}

func sliceOps() {
	var slc []int
	for i := 0; i < 100; i++ {
		printSlice(slc)
		slc = append(slc, 2*i+1)
	}
	fmt.Println(slc)
}

func testMap() {
	m1 := make(map[string]int)
	fmt.Println(m1)

	m2 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	fmt.Println(m2)
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	for k := range m2 {
		fmt.Println(k)
	}
	for _, v := range m2 {
		fmt.Println(v)
	}

	// 如果key不存在，则输出初始值
	if vTest, ok := m2["ddaf"]; ok {
		fmt.Println(vTest)
	} else {
		fmt.Println("not exist")
	}
	fmt.Println()
}

func main() {
	enums()
	//readFile()

	printFile("abc.txt")
	//forever()

	q, r := div(9, 3)
	fmt.Println(q, r)
	fmt.Println(apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i), float64(i2)))
	}, 3, 4))

	fmt.Println(sums(1, 2, 3))

	a, b := 5, 6
	swap(&a, &b)
	fmt.Println(a, b)

	arraysTest()

	sliceTest()

	resliceTest()

	sliceOps()

	testMap()
}
