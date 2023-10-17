package main

import (
	"fmt"
	"net/http"
)

/*
协程泄露的场景
http request body未关闭

resp.Body.Close() 未被调用时，goroutine不会退出
*/

func requestWithNoClose() {
	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("error occurred while fetching page, error: %s\n", err.Error())
	}
}

func requestWithClose() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("error occurred while fetching page, error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()
}

func block4() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestWithClose()
		}()
	}
}

//var wg = sync.WaitGroup{}

func main() {
	block4()
	wg.Wait()
}
