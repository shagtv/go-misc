package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urlPattern := "https://habrahabr.ru/post/%d/"

	for i := 3000; i < 3050; i++ {
		url := fmt.Sprintf(urlPattern, i)
		getHtml(url)
	}
}

func getHtml(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(url, " Length:", len(string(bytes)))
}
