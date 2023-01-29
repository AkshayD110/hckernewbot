package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type HackerNews struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

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
		var hackernews HackerNews
		json.Unmarshal([]byte(string(body)), &hackernews)
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.AddConfigPath(".")
		err = viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("couldn't read config file"))
		}
		topic_of_interest := viper.GetStringSlice("interestedtopics")

		for _, hitwords := range topic_of_interest {
			if strings.Contains(strings.ToLower(hackernews.Title), hitwords) {
				fmt.Println(hackernews.Title)
			}
		}

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
