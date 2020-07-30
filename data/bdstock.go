package data

import "encoding/json"

type BDStock struct {
	Meta struct {
		FileName        string `json:"file_name"`
		FileDescription string `json:"file_description"`
		Package         string `json:"package"`
		License         string `json:"license"`
		Author          string `json:"author"`
		Encoding        string `json:"encoding"`
	} `json:"meta"`
	Data struct {
		Companies []struct {
			Ticker string   `json:"ticker"`
			Sector string   `json:"sector"`
			Names  []string `json:"names"`
		} `json:"companies"`
	} `json:"data"`
}

func LoadBDStock() (*BDStock, error) {
	binData, err := Asset("data/bdstock.json")
	if err != nil {
		return nil, err
	}
	return LoadJSON(binData)
}

func LoadJSON(b []byte) (*BDStock, error) {
	d := &BDStock{}
	err := json.Unmarshal(b, d)
	return d, err
}
