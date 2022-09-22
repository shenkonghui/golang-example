package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	clt := http.Client{}
	url := "http://www.google.com"
	resp, err := clt.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
