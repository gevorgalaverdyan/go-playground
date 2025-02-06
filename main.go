package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// const apiKey=""
type University struct {
	Name    string   `json:"name"`
	// Country string   `json:"country"`
	// Domains []string `json:"domains"`
	// WebPages []string `json:"web_pages"`
}
//http://universities.hipolabs.com/search?country=Canada
func fetchUnis(country string) interface{} {
	var data []University
	
	url := fmt.Sprintf("http://universities.hipolabs.com/search?country=%s", country)
	res, err := http.Get(url)

	if err!=nil {
		fmt.Print("Error", err)
		return data
	}

	defer res.Body.Close()

		if err:=json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Print("Decode Error", err)
		return data
	}

	return data
}

func main(){
	unis := fetchUnis("Canada")
	
	for uni := range unis{
		
	}
}