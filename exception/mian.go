package main

import (
	"errors"
	"fmt"
)

func main() {
	// 错误拦截
	// painc1()
	// error嵌套
	// warp()
	// 异常抛出
	throw()

}

func throw() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	errFunc()
}
func errFunc() {
	panic(errors.New("hello err"))
}

func warp() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(errors.Unwrap(w))
}

func painc1() {
	fmt.Printf("hello world my name is %s, I'm %d\r\n", "songxingzhu", 26)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	myPainc()
	fmt.Printf("这里应该执行不到！")
}
func myPainc() {
	var x = 30
	var y = 0
	//panic("我就是一个大错误！")
	var c = x / y
	fmt.Println(c)
}
