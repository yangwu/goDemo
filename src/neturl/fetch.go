package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		rturl := geturl(url)
		resp, err := http.Get(rturl)
		
		fmt.Println("get url status:" + resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		count, err := getBodyByIOCopy(resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("length of body:%d", count)
		/*
			b, err := getBodyByIOUtil(resp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("content of url: %s = \n%s", url, b)
		*/
	}
}

func getBodyByIOUtil(resp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return b, err
}

func getBodyByIOCopy(resp *http.Response) (int64, error) {
	count, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	fmt.Fprintf(os.Stdout, "content length: %d", count)
	return count, err
}
func geturl(url string) string {
	fmt.Println("init url = " + url)
	if (!strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://")) {
		url = "http://" + url
	}
	fmt.Println("final url = " + url)
	return url
}
