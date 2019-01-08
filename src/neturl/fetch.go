package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	
	concurFetch()
	/*for _, url := range os.Args[1:] {

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
			b, err := getBodyByIOUtil(resp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("content of url: %s = \n%s", url, b)
		
	}*/
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
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	fmt.Println("final url = " + url)
	return url
}

func concurFetch() {
	start := time.Now()
	ch := make(chan string)
	for _,url :=range os.Args[1:]{
		go fetch(url,ch)	
	}
	
	for index,_ := range os.Args[1:]{
		fmt.Printf("current index: %d, time: %d |  ", index,time.Now().Nanosecond());
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string){
	start := time.Now()
	resp,err := http.Get(url)
	
	if err != nil{
		ch <- fmt.Sprint(err)
	}
	nbytes,err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <-fmt.Sprintf("while reading %s:%v", url,err);
		return 
	}
	secs := time.Since(start).Seconds()
	
	ch<-fmt.Sprintf("%.2fs    %7d    %s", secs,nbytes, url)
}
