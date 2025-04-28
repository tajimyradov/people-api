package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"people-api/config"
)

type Enricher struct {
	cfg *config.Config
}

func NewEnricher(cfg *config.Config) *Enricher {
	return &Enricher{cfg: cfg}
}

type EnrichedData struct {
	Age         int
	Gender      string
	Nationality string
}

func (e *Enricher) Enrich(name string) (EnrichedData, error) {
	var data EnrichedData

	// Get Age
	resp, err := http.Get(fmt.Sprintf("%s?name=%s", e.cfg.AgifyURL, name))
	if err == nil {
		defer resp.Body.Close()
		var result struct{ Age int }
		json.NewDecoder(resp.Body).Decode(&result)
		data.Age = result.Age
	}

	// Get Gender
	resp, err = http.Get(fmt.Sprintf("%s?name=%s", e.cfg.GenderizeURL, name))
	if err == nil {
		defer resp.Body.Close()
		var result struct{ Gender string }
		json.NewDecoder(resp.Body).Decode(&result)
		data.Gender = result.Gender
	}

	// Get Nationality
	resp, err = http.Get(fmt.Sprintf("%s?name=%s", e.cfg.NationalizeURL, name))
	if err == nil {
		defer resp.Body.Close()
		var result struct {
			Country []struct {
				CountryID string `json:"country_id"`
			} `json:"country"`
		}
		json.NewDecoder(resp.Body).Decode(&result)
		if len(result.Country) > 0 {
			data.Nationality = result.Country[0].CountryID
		}
	}

	return data, nil
}
