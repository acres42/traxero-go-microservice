package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type VINResponse struct {
	Results []struct {
		Make      string `json:"Make"`
		Model     string `json:"Model"`
		ModelYear string `json:"ModelYear"`
	} `json:"Results"`
}

type Vehicle struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  string `json:"year"`
}

func apihandler(w http.ResponseWriter, r *http.Request) {
	tokenCheck := r.Header.Get("Authorization")
	token := BearerAuthHeader(tokenCheck)

	if token == "ShYj6GJ37rXBFRiTew3GkdKVMpcp52" {
		if r.Method == "GET" {
			// Get the VIN from the request
			vin := r.URL.Query().Get("vin")
			if len(vin) == 0 {
				http.Error(w, "Invalid request: missing VIN", http.StatusBadRequest)
				return
			}
			safeVin := url.QueryEscape(vin)
			url := fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/decodevinvalues/%s?format=json", safeVin)

			// Create a new HTTP request to the VIN decoder API
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal("NewRequest: ", err)
				return
			}
			// For control over HTTP client headers,
			// redirect policy, and other settings,
			// create a Client
			// A Client is an HTTP client
			client := &http.Client{}

			// Send the request via a client
			// Do sends an HTTP request and
			// returns an HTTP response
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal("Do: ", err)
				return
			}

			// Callers should close resp.Body
			// when done reading from it
			// Defer the closing of the body
			defer resp.Body.Close()

			// Fill the record with the data from the JSON
			var record VINResponse

			// Use json.Decode for reading streams of JSON data
			if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
				log.Println(err)
			}

			vehicle := Vehicle{
				Make:  strings.ToLower(record.Results[0].Make),
				Model: strings.ToLower(record.Results[0].Model),
				Year:  record.Results[0].ModelYear,
			}

			json_bytes, err := json.Marshal(vehicle)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(json_bytes)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["message"] = "Bad Request"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Unauthorized"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

// BearerAuthHeader validates incoming `r.Header.Get("Authorization")` header
// and returns token otherwise an empty string.
func BearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}

func main() {

	//start the server
	log.Println("Listening for requests at http://localhost:8000/")
	http.ListenAndServe(":8000", http.HandlerFunc(apihandler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
