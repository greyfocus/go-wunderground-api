package wunderground_api

import (
	"os"
	"path/filepath"
	"testing"
)

const TEST_CONDITIONS_PARIS_FILE = "test-conditions-paris.json"

func TestGeneralParsing(t *testing.T) {
	testfile := filepath.Join("testdata", TEST_CONDITIONS_PARIS_FILE)

	testData, err := os.Open(testfile)
	defer testData.Close()
	if err != nil {
		t.Fatal("Missing test data file: " + testfile)
	}

	response, err := parseWeatherResponse(testData)
	if err != nil {
		t.Fatal("Unable to parse the test file: ", err)
	}

	conditions := response.CurrentConditions

	if conditions == nil {
		t.Error("Missing current conditions")
	}

	assertFloat("invalid tempC", conditions.TempC, 11.0, t)
	assertFloat("invalid tempF", conditions.TempF, 52, t)
	assertFloat("invalid feelsLikeC", conditions.FeelsLikeC, 11, t)
	assertFloat("invalid feelsLikeF", conditions.FeelsLikeF, 52, t)
	assertFloat("invalid pressureMb", conditions.PressureMb, 1005, t)
	assertString("invalid relativeHumidity", conditions.RelativeHumidity, "62%", t)
	assertFloat("invalid visibility_km", conditions.VisibilityKm, 10, t)
	assertString("invalid weather string", conditions.Weather, "Mostly Cloudy", t)
	assertString("invalid wind string", conditions.Wind, "From the North at 10 MPH", t)
	assertFloat("invalid windKph", conditions.WindKph, 17, t)
	assertFloat("invalid windGustKph", conditions.WindGustKph, 0, t)
	assertFloat("invalid windDegrees", conditions.WindDegrees, 350, t)
	assertString("invalid windDir", conditions.WindDir, "North", t)
}

func assertFloat(m string, value float32, expected float32, t *testing.T) {
	if value != expected {
		t.Errorf("%s - expected: %2.2f, but was: %2.2f", m, expected, value)
	}
}

func assertString(m string, value string, expected string, t *testing.T) {
	if value != expected {
		t.Errorf("%s - expected: %s, but was: %s", m, expected, value)
	}
}
