package USDAAPIWrapper

import "encoding/json"

// Food Search Types

type FoodSearch struct {
	List Info `json:"list"`
}

type Info struct {
	Query           string           `json:"q"`
	StandardRelease string           `json:"sr"`
	DatabaseSource  string           `json:"ds"`
	Start           int              `json:"start"`
	End             int              `json:"end"`
	Total           int              `json:"total"`
	Group           string           `json:"group"`
	Sort            string           `json:"sort"`
	Item            []FoodSearchItem `json:"item"`
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
	Foods    []Food  `json:"foods"`
	Count    int     `json:"count"`
	NotFound int     `json:"notfound"`
	API      float32 `json:"api"`
}

type Food struct {
	Food Data `json:"food"`
}

type Data struct {
	StandardRelease string     `json:"sr"`
	Type            string     `json:"type"`
	Description     FoodDesc   `json:"desc"`
	Nutrients       []Nutrient `json:"nutrients"`
}

type FoodDesc struct {
	NDBNO              string `json:ndbno`
	Name               string `json:"name"`
	DataSource         string `json:"ds"`
	ShortDescription   string `json:"sd"`
	FoodGroup          string `json:"fg"`
	ScientificName     string `json:"sn"`
	CommercialName     string `json:"cn"`
	Manufacturer       string `json:"manu"`
	NitrogenFactor     int    `json:"nf"`
	CarbohydrateFactor int    `json:"cf"`
	FatFactor          int    `json:"ff"`
	ProteinFactor      int    `json:"pf"`
	RefusePercent      string `json:"r"`
	RefuseDescription  string `json:"rd"`
	DatabaseSource     string `json:"ds"`
	ReportingUnit      string `json:"ru"`
}

type Nutrient struct {
	NutrientID    string      `json:"nutrient_id"`
	Name          string      `json:"name"`
	Group         string      `json:"group"`
	Unit          string      `json:"unit"`
	Value         string      `json:"value"`
	Derivation    string      `json:"derivation"`
	DataPoints    json.Number `json:"dp"`
	StandardError string      `json:"se"`
}

type Measure struct {
	Label      string  `json:"label"`
	Equivalent float32 `json:"eqv"`
	Eunit      string  `json:"eunit"`
	Quantity   int     `json:"qty"`
	Value      string  `json:"value"`
}
