package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const baseurl = "https://api.openweathermap.org/data/2.5/weather"

type WeatherStruct struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func apiCall(c *gin.Context) {
	//?q={city name}&appid={API key}
	city := "London"
	res, err := http.Get(baseurl + "?q=" + city + "&appid=2b64dd976322d684ff5c17f2f86e058f")
	fmt.Println("URL ", baseurl+"?q="+city+"&appid=2b64dd976322d684ff5c17f2f86e058f")
	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	var weather WeatherStruct

	if err := json.Unmarshal(bytes, &weather); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(weather)
}

func main() {
	router := gin.Default()

	router.GET("/weather", apiCall)

	router.Run()
}
