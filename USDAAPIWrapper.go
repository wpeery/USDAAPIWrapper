package USDAAPIWrapper

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

func SearchFood(query, dataSource,
	foodGroupID, sort, max, offset string) *FoodSearch {
	response := doRequest(buildSearchRequest(query, dataSource,
		foodGroupID, sort, max, offset))
	search := new(FoodSearch)
	formatSearchResponse(response, search)
	return search
}

func GetFoodReport(query, reportType string) *FoodReport {
	response := doRequest(buildReportRequest(query, reportType))
	report := new(FoodReport)
	formatSearchResponse(response, report)
	return report
}

func buildReportRequest(ndbno, reportType string) *http.Request {
	url := fmt.Sprintf("https://api.nal.usda.gov/ndb/V2/reports?ndbno=%s&type=%s&format=json&api_key=%s",
		ndbno, reportType, Config.API_KEY)
	return buildRequest(url)
}

func buildSearchRequest(query, dataSource,
	foodGroupID, sort, max, offset string) *http.Request {
	safeQuery := url.QueryEscape(query)
	url := fmt.Sprintf("http://api.nal.usda.gov/ndb/search/?format=json&q=%s&ds=%s&fg=%s&sort=%s&max=%s&offset=%s&api_key=%s",
		safeQuery, dataSource, foodGroupID, sort, max, offset, Config.API_KEY)
	fmt.Println(url)
	return buildRequest(url)
}

func buildRequest(url string) *http.Request {
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

func init() {
	initConfig("./config.json")
}
