package main

import (
	"fmt"
	"net/http"
)

func main() {

	//simple data for sample
	sampleUrls := []string{
		"https://www.tokopedia.com/",
		"https://www.traveloka.com/",
		"https://www.gojek.com/",
		"https://www.grab.com/",
		"https://www.bukalapak.com/",
	}
	channel := make(chan inputUrlType)
	for _, url := range sampleUrls {
		go urlCheck(url, channel)
	}

	result := make([]inputUrlType, len(sampleUrls))

	for i, _ := range result {

		result[i] = <-channel
		if result[i].status {
			fmt.Println(result[i].url, "is up.")
		} else {
			fmt.Println(result[i].url, "is down !")
		}

	}
}

//check and prints a message if a website is up or down
func urlCheck(url string, channel chan inputUrlType) {
	_, err := http.Get(url)
	if err != nil {
		channel <- inputUrlType{url, false}
	} else {
		//the website is up
		channel <- inputUrlType{url, true}
	}
}

type inputUrlType struct {
	url    string
	status bool
}
