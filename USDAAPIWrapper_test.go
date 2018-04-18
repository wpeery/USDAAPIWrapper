package USDAAPIWrapper

import (
	"fmt"
	"testing"
)

func TestBuildRequest(t *testing.T) {
	url := "www.test.com"
	req_want := "&{GET " + url + " HTTP/1.1 1 1 map[] <nil> <nil> 0 [] false  map[] map[] <nil> map[]   <nil> <nil> <nil> <nil>}"
	req_got := fmt.Sprint(buildRequest(url))
	if req_want != req_got {
		t.Errorf("Error in buildRequest, want: %s,\n\t\t\t\t\t   got: %s", req_want, req_got)
	}
}

func TestBuildReportURL(t *testing.T) {
	ndbno := "111"
	reportType := "f"
	url := "https://api.nal.usda.gov/ndb/V2/reports?ndbno=111&type=f&format=json&api_key=" + Config.API_KEY
	if url != buildReportURL(ndbno, reportType) {
		t.Errorf("Error in buildReportURL want: %s, got: %s", url, buildReportURL(ndbno, reportType))
	}
}

func TestBuildSearchURL(t *testing.T) {
	url_want := "http://api.nal.usda.gov/ndb/search/?format=json&q=111&ds=sr&fg=1&sort=r&max=50&offset=25&api_key=" + Config.API_KEY
	url_got := buildSearchURL("111", "sr", "1", "r", "50", "25")
	if url_want != url_got {
		t.Errorf("buildSearchURL is not correct want: %s,\n\t\t\t\t\t\t  got: %s", url_want, url_got)
	}
}

func TestFormatResponse(t *testing.T) {
	url := "https://api.nal.usda.gov/ndb/V2/reports?ndbno=01009&type=s&format=json&api_key=9ZH2gqcl0QxJfd7aqlhaeIdckvl0ha3qFWnHtVUH"
	// resp := doRequest(buildRequest(url))
	fmt.Println(url)
	t.Errorf("STUB!")
}

func TestInitConfig(t *testing.T) {
	t.Errorf("STUB!")
}

func TestInit(t *testing.T) {
	t.Errorf("STUB!")
}

func TestSearchFood(t *testing.T) {
	// foodSearch := SearchFood("cheese", "Standard%20Reference", "", "r", "25", "0")
	// fmt.Printf("%+v\n", foodSearch)
	// fmt.Println(foodSearch.List.Query)
	t.Errorf("STUB!")
}

func TestGetFoodReport(t *testing.T) {
	// foodReport := GetFoodReport("01009", "s")
	// fmt.Println(foodReport)
	t.Errorf("STUB!")
}

func TestDoRequest(t *testing.T) {
	t.Errorf("STUB!")
}
