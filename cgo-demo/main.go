package main

/*
   #include <stdio.h>
   int SayHello() {
    puts("Hello World");
       return 0;
   }
*/

//#cgo CFLAGS: -I./func
//#cgo LDFLAGS: -L${SRCDIR}/func -ladd
//#include "add.h"
import "C"
import (
	"fmt"
)

func main() {
	// ret := C.SayHello()
	// fmt.Println(ret)

	fmt.Println(C.add(1, 2))
}
