package main

import (
  "./wunderground_api"
  "fmt"
)

const API_KEY = "<yourkey>"

func main() {
  client := wunderground_api.JsonClient{ ApiKey: API_KEY }
  request := wunderground_api.Request{ Features: []string{ "conditions" }, Location: "Romania/Cluj-Napoca"}

  resp, err := client.Execute(&request)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("       Temperature: %2.1f C (feels like %2.1f C)\n", resp.CurrentConditions.TempC, resp.CurrentConditions.FeelsLikeC);
  fmt.Printf(" Relative Humidity: %s%\n", resp.CurrentConditions.RelativeHumidity)
  fmt.Printf("              Wind: %2.1f km/h from %2.1f %s\n",
    resp.CurrentConditions.WindKph, resp.CurrentConditions.WindDegrees, resp.CurrentConditions.WindDir)
  fmt.Printf("          WindGust: %2.1f km/h\n", resp.CurrentConditions.WindGustKph)
}
