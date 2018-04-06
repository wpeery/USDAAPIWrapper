package main

// Food Search Types
type FoodSearch struct {
	List Info `json:"list"`
}

type Info struct {
	Q     string           `json: "q"`
	Sr    string           `json: "sr"`
	Ds    string           `json: "ds"`
	Start int              `json: "start"`
	End   int              `json: "end"`
	Total int              `json: "total"`
	Group string           `json: "group"`
	Sort  string           `json: "sort"`
	Item  []FoodSearchItem `json: "item"`
}

type FoodSearchItem struct {
	Offset     int    `json:"offset"`
	Group      string `json:"group"`
	Name       string `json:"name"`
	NDBNO      string `json:"ndbno"`
	DataSource string `json:"ds"`
}

// Food Report Types
type FoodReport struct {
}

// fmt.Println("Q: ", search.List.Q)
// fmt.Println("SR: ", search.List.Sr)
// fmt.Println("DS: ", search.List.Ds)
// fmt.Println("Start: ", search.List.Start)
// fmt.Println("End: ", search.List.End)
// fmt.Println("Total: ", search.List.Total)
// fmt.Println("Group: ", search.List.Group)
// fmt.Println("Sort: ", search.List.Sort)
// fmt.Println("item[0].Offset: ", search.List.Item[0].Offset)
// fmt.Println("item[0].Group: ", search.List.Item[0].Group)
// fmt.Println("item[0].Name: ", search.List.Item[0].Name)
// fmt.Println("item[0].NDBNO: ", search.List.Item[0].NDBNO)
// fmt.Println("item[0].DataSource: ", search.List.Item[0].DataSource)
