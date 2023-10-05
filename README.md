# Weather API in Golang

This is a simple Golang weather API that acts as a proxy to the OpenWeatherMap API. It allows you to retrieve weather data for a specific city.

## Prerequisites

Before you begin, ensure you have the following prerequisites installed on your system:

- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/dl/) (only for development)

## Build and Run the Application

Follow these steps to build and run the application in a Docker container:

## Clone this repository:
```bash
git clone https://github.com/ar4mirez/workshop-zero-to-hero-devops.git
```

## Navigate to the project directory

```shell
cd weatherapp
```

## Build the Docker image
```shell
docker build -t weatherapp .
```

## Run the Docker container
```shell
docker run -p 8080:8080 -e OPENWEATHERMAP_API_KEY=your_api_key_here weatherapp
```

> Replace your_api_key_here with your actual OpenWeatherMap API key.
The weather API will be accessible at http://localhost:8080/weather?city=CityName.

