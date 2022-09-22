package main

import (
	"bou.ke/monkey"
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "Get", func(c *http.Client, url string) (*http.Response, error) {
		guard.Unpatch()
		defer guard.Restore()
		fmt.Print("hook")
		url = "https://www.baidu.com"
		return c.Get(url)
	})

	resp, err := http.Get("https://google.com")
	fmt.Println(resp.Status, err) // 200 OK <nil>
}
