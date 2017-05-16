package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func main() {
	urlPattern := "https://habrahabr.ru/post/%d/"

	var wg sync.WaitGroup

	for i := 3000; i < 3030; i++ {
		url := fmt.Sprintf(urlPattern, i)

		wg.Add(1)
		go func(url string) {
			getHtml(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func getHtml(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	if len(string(bytes)) == 1211 {
		fmt.Println(string(bytes))
		os.Exit(0)
	}

	fmt.Println(url, " Length:", len(string(bytes)))
}
