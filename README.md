# Weather Web Application

This is a simple Go-based web application that fetches and displays real-time weather data for cities using the OpenWeatherMap API. The user can search for a city and choose the temperature unit (Celsius or Fahrenheit) to get the current weather conditions.

## Features

- Fetches real-time weather data including temperature, humidity, wind speed, and weather conditions.
- Supports both metric (Celsius) and imperial (Fahrenheit) units for temperature display.
- Provides a simple web interface where users can input a city and select units for the weather information.

## Technologies Used

1. **Go (Golang)**: The main programming language for the application.
2. **OpenWeatherMap API**: Provides the real-time weather data for the requested city.
3. **HTML Templates**: Used to render the weather information on the web page.
4. **net/http**: Go's built-in package to handle HTTP requests.

## How to Run

### 1. Clone the Repository

First, clone this repository to your local machine:

```bash
git clone https://github.com/your-username/weather-app.git
cd weather-app
```

### 2. Install Dependencies
Make sure you have Go installed. If you don't have Go installed, download and install it from here.

### 3. Set up the OpenWeatherMap API Key
Create an API key from OpenWeatherMap. Once you have the API key, create a .apiConfig file in the root directory and add the following content:

```bash
{
    "OpenWeatherMapApiKey": "your-api-key-here"
}
```
