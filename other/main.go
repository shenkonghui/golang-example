package main

import (
	"fmt"
	"time"
)

type A struct {
	b *B
}
type B struct {
	c C
}

type C struct {
	v int
}

func main() {
	a := A{}
	fmt.Println(a.b.c.v)
	time.Sleep(time.Second * 10)
}
