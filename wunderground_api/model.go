package wunderground_api

type Request struct {
  Features []string
  Location string
}

type Response struct {
  CurrentConditions Conditions `json:"current_observation"`
}

type Conditions struct {
  TempC float32 `json:"temp_c"`
  TempF float32 `json:"temp_f"`

  FeelsLikeF float32 `json:"feelslike_c,string"`
  FeelsLikeC float32 `json:"feelslike_f,string"`

  PressureMb float32 `json:"pressure_mb,string"`
  RelativeHumidity string `json:"relative_humidity"`

  VisibilityKm float32 `json:"visibility_km,string"`

  Weather string `json:"weather"`

  Wind string `json:"wind_string"`
  WindKph float32 `json:"wind_kph,float"`
  WindGustKph float32 `json:"wind_gust_kph,string"`
  WindDegrees float32 `json:"wind_degrees,float"`
  WindDir string `json:"wind_dir"`
}
