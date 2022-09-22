package main

import (
	"errors"
	"fmt"
)

func main() {
	a := do()
	fmt.Println(a)

}

func do() string{

	defer func() string{
		if err := recover(); err != nil {
		}
		return "painc"
	}()
	panic(errors.New("aaa"))
	return "abc"
}