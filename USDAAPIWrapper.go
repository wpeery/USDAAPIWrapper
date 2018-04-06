package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	// "tokenbucket" limit requests to 1,000 requests per hour
)

var Config struct {
	API_KEY  string
	TIME_OUT int
}

func SearchFood(query string) *FoodSearch {
	response := doRequest(buildSearchRequest(query))
	search := new(FoodSearch)
	formatSearchResponse(response, search)
	return search
}

func GetFoodReport(query string) *FoodReport {
	response := doRequest(buildSearchRequest(query))
	report := new(FoodReport)
	formatSearchResponse(response, report)
	return report
}

func buildReportRequest(query string) *http.Request {
	safeQuery := url.QueryEscape(query)
	url := "https://api.nal.usda.gov/ndb/V2/reports?ndbn0=" +
		safeQuery +
		"&type=f&format=json&api_key=" +
		Config.API_KEY
	req, err := http.NewRequest("GET", url, nil)
	fmt.Println("req: ", req)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	return req
}

func buildSearchRequest(query string) *http.Request { // Multiple requests
	safeQuery := url.QueryEscape(query)
	url := "http://api.nal.usda.gov/ndb/search/?format=json&q=" +
		safeQuery +
		"&sort=r&max=25&offset=0&api_key=" +
		Config.API_KEY
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	return req
}

func doRequest(req *http.Request) *http.Response {
	client := &http.Client{Timeout: time.Second * time.Duration(Config.TIME_OUT)}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	return resp
}

func formatSearchResponse(resp *http.Response, record interface{}) {
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(record); err != nil {
		panic(err)
	}
}

func initConfig(configFilename string) {
	file, err := os.Open(configFilename)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
	file.Close()
}

func main() {
	initConfig("./config.json")
	search := SearchFood("butter")
	fmt.Println(search)
}
