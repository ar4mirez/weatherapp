package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// WeatherResponse represents the structure of the weather data returned by the API.
type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

type WeatherLiveResponse struct {
	Message string `json:"message"`
}

func main() {
	// Get your OpenWeatherMap API key from an environment variable.
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set your OPENWEATHERMAP_API_KEY environment variable.")
		return
	}

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Get your OpenWeatherMap API key from an environment variable.
		apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
		if apiKey == "" {
			http.Error(w, "Please set your OPENWEATHERMAP_API_KEY environment variable.", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&WeatherLiveResponse{
			Message: "Yes the application is ready",
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&WeatherLiveResponse{
			Message: "Yes it's alive",
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Define a handler function to get weather data from OpenWeatherMap.
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		// Build the API URL.
		apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

		// Make a GET request to OpenWeatherMap.
		resp, err := http.Get(apiUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response body.
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse the JSON response.
		var weather WeatherResponse
		if err := json.Unmarshal(body, &weather); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the weather data as JSON response.
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(weather); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server.
	fmt.Println("Weather API server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
