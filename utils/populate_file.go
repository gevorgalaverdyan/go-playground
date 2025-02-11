package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gevorgalaverdyan/go-playground/models"
)

type API_Response struct {
	State   string    `json:"state-province"`
	Country string    `json:"country"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date,omitempty"`
}

// Example country list
var countries = Countries

func PopulateFile() error {
	fmt.Println("Starting...")

	var wg sync.WaitGroup
	resultsChan := make(chan []models.University) // Channel to collect results

	// Start a goroutine for each country
	for _, country := range countries {
		wg.Add(1)

		go func(c string) {
			defer wg.Done()

			// Fetch universities from API
			apiUnis, err := fetchUniversities(c)
			if err != nil {
				fmt.Printf("Error fetching data for %s: %v\n", c, err)
				return
			}

			// Send the result to the channel
			resultsChan <- apiUnis
		}(country)
	}

	// **Goroutine to close the channel once all fetch operations are done**
	go func() {
		wg.Wait()       // Wait for all goroutines to finish
		close(resultsChan) // Close the channel to signal completion
	}()

	// Collect results from channel
	var unis []models.University
	for universities := range resultsChan {
		unis = append(unis, universities...) // Collect results safely
	}

	// Write results to file
	unisJson, err := json.MarshalIndent(unis, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling universities: %v", err)
	}

	err = os.WriteFile("data.json", unisJson, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Println("Data successfully written to data.json")
	return nil
}

// Fetch universities for a given country
func fetchUniversities(country string) ([]models.University, error) {
	baseURL := "http://universities.hipolabs.com/search?country="

	res, err := http.Get(baseURL + country)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var apiResponses []API_Response
	err = json.Unmarshal(body, &apiResponses)
	if err != nil {
		return nil, err
	}

	// Convert API responses to models.University
	var universities []models.University
	for _, apiUni := range apiResponses {
		universities = append(universities, models.University{
			Name:        apiUni.Name,
			Description: "Uni located in: " + apiUni.State + ", " + apiUni.Country,
			Location:    apiUni.State + ", " + apiUni.Country,
			DateTime:    time.Now(),
		})
	}

	return universities, nil
}
