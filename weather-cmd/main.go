package main

import (
	"github.com/greyfocus/go-wunderground-api"
	"flag"
	"fmt"
)

const API_KEY = "<yourkey>"

func main() {

	location := flag.String("l", "France/Paris", "the location for which to query the weather")
	api_key := flag.String("api_key", API_KEY, "the API key to be used")
	flag.Parse()

	client := wunderground_api.JsonClient{*api_key}
	request := wunderground_api.Request{Features: []string{"conditions"}, Location: *location}

	resp, err := client.Execute(&request)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.CurrentConditions == nil {
		fmt.Println("The current conditions were not returned. Is the API key correct?")
		return
	}

	fmt.Printf("       Temperature: %6.1f C (feels like %2.1f C)\n", resp.CurrentConditions.TempC, resp.CurrentConditions.FeelsLikeC)
	fmt.Printf("          Pressure: %6.1f mb\n", resp.CurrentConditions.PressureMb)
	fmt.Printf(" Relative Humidity: %7s\n", resp.CurrentConditions.RelativeHumidity)
	fmt.Printf("              Wind: %6.1f km/h from %2.1f %s\n",
		resp.CurrentConditions.WindKph, resp.CurrentConditions.WindDegrees, resp.CurrentConditions.WindDir)
	fmt.Printf("          WindGust: %6.1f km/h\n", resp.CurrentConditions.WindGustKph)
	fmt.Printf("  Observation Time: %s\n", resp.CurrentConditions.ObservationTime)
}
