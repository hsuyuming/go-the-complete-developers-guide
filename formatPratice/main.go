//參考文件1.https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter01/01.3.md
//參考文件2.https://gobyexample.com/string-formatting
package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)
	//result {1 2}

	fmt.Printf("%+v\n", p)
	//result {x:1 y:2}

	fmt.Printf("%#v\n", p)
	//result main.point{x:1, y:2}

	fmt.Printf("%T\n", p)
	// print type of value
	//result main.point

	fmt.Printf("%t\n", p)
	//result {%!t(int=1) %!t(int=2)}

	fmt.Printf("%t\n", true)
	//result true

	fmt.Printf("%d\n", 10)
	//result ??????

	fmt.Printf("%b\n", 2)
	//result 10

	fmt.Printf("%c\n", 33)
	// 相印unicode表示符號
	//result !

	fmt.Printf("%c\n", 0x4E2D)
	// 相印unicode表示符號
	//result 中

	fmt.Printf("%x\n", 456)
	// 16進制
	//result 1c8

	fmt.Printf("%s\n", "\"string\"")
	//result "string"

	fmt.Printf("%q\n", "\"string\"")
	//result "\"string\""

	fmt.Printf("%p\n", &p)
	//輸出指標的記憶體位置
	//result 0xc4200160a0

	/*
		fmt.Printf("", p)
		//result

		fmt.Printf("", p)
		//result

		fmt.Printf("", p)
		//result

		fmt.Printf("", p)
		//result

	*/

}
