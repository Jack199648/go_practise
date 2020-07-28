package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func netHttp(url string) (err error) {
	fmt.Println()
	resp, err := http.Get(url)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("status:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body:%v\n", string(bodyContent))
	return
}
