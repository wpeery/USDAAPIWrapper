package USDAAPIWrapper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"tokenbucket"
)

type RateLimiter interface {
	Consume(int)
}

var Config struct {
	API_KEY          string
	CLIENT_TIME_OUT  int
	API_RATE_LIMITER RateLimiter
}

func SearchFood(query, dataSource,
	foodGroupID, sort, max, offset string) *FoodSearch {
	response := doRequest(buildRequest(buildSearchURL(query, dataSource,
		foodGroupID, sort, max, offset)))
	search := new(FoodSearch)
	formatResponse(response, search)
	return search
}

func GetFoodReport(query, reportType string) *FoodReport {
	response := doRequest(buildRequest(buildReportURL(query, reportType)))
	report := new(FoodReport)
	formatResponse(response, report)
	return report
}

func buildReportURL(ndbno, reportType string) string {
	url := fmt.Sprintf("https://api.nal.usda.gov/ndb/V2/reports?ndbno=%s&type=%s&format=json&api_key=%s",
		ndbno, reportType, Config.API_KEY)
	return url
}

func buildSearchURL(query, dataSource, // escape the arguments
	foodGroupID, sort, max, offset string) string {
	safeQuery := url.QueryEscape(query)
	url := fmt.Sprintf("http://api.nal.usda.gov/ndb/search/?format=json&q=%s&ds=%s&fg=%s&sort=%s&max=%s&offset=%s&api_key=%s",
		safeQuery, dataSource, foodGroupID, sort, max, offset, Config.API_KEY)
	return url
}

func buildRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	return req
}

func doRequest(req *http.Request) *http.Response {
	client := &http.Client{Timeout: time.Second * time.Duration(Config.CLIENT_TIME_OUT)}
	Config.API_RATE_LIMITER.Consume(1) // This statement blocks if there are no more available requests
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	return resp
}

func formatResponse(resp *http.Response, record interface{}) {
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
	Config.API_RATE_LIMITER = tokenbucket.TokenBucket(1000, 3600)
}
