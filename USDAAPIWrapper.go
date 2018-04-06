package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	// "tokenbucket" limit requests to 1,000 requests per hour
)

/*
	SEARCH
	* list of up to 20 nutrients per food
	* Nutrients have weight and stuff

	FOOD REPORT(v2)
		v2:
			* list of food nutrients or food groups
			* good for paged browse
			* Can get report for multiple foods (up to 50)
*/
var Config struct {
	API_KEY  string
	TIME_OUT int
}

func SearchFood(query string) *FoodSearch {
	return formatSearchResponse(doRequest(buildSearchRequest(query)))
}

func GetFoodReport(query string) {
	formatReportResponse(doRequest(buildReportRequest(query)))
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
	client := &http.Client{
		Timeout: time.Second * time.Duration(Config.TIME_OUT),
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	return resp
}

func formatSearchResponse(resp *http.Response) *FoodSearch {
	defer resp.Body.Close()
	record := new(FoodSearch)
	if err := json.NewDecoder(resp.Body).Decode(record); err != nil {
		panic(err)
	}
	return record
}

func formatReportResponse(resp *http.Response) *FoodReport { // Multiple formatted responses
	defer resp.Body.Close()
	record := new(FoodReport)
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	return record
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
