package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func get_page(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("request body error")
	}
	return resp
}

func build_url(request_ids string) {

	request_ids = strings.ReplaceAll(request_ids, "[", "")
	request_ids = strings.ReplaceAll(request_ids, "]", "")
	request_ids_slice := strings.Split(request_ids, ",")
	for _, each_request_id := range request_ids_slice {
		each_request_id = strings.TrimSpace(each_request_id)
		request_url := "https://hacker-news.firebaseio.com/v0/item/myrequestid.json?print=pretty"
		request_url = strings.ReplaceAll(request_url, "myrequestid", each_request_id)
		each_request_body := get_page(request_url)
		defer each_request_body.Body.Close()

		body, err := ioutil.ReadAll(each_request_body.Body)
		if err != nil {
			fmt.Println("error in getting response body")
		}
		fmt.Println(string(body))

	}
	fmt.Println(len(request_ids_slice))

}

func main() {
	fmt.Println("checking")
	url_requestId := "https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty"
	resp := get_page(url_requestId)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in getting response body")
	}
	// fmt.Println(string(body))
	build_url(string(body))
}
