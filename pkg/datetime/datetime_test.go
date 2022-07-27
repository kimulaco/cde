package datetime

import (
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	resultTime := time.Date(2022, time.January, 1, 12, 30, 10, 0, time.UTC)
	result := TimeToString(resultTime)
	if result != "2022-01-01 12:30:10" {
		t.Error("Recieve: " + result)
	}
}

func TestStringToTime(t *testing.T) {
	result, err := StringToTime("2022-01-01 12:30:10")

	if err != nil {
		t.Error(err)
	}

	if result.Year() != 2022 ||
		int(result.Month()) != 1 ||
		result.Day() != 1 ||
		result.Hour() != 12 ||
		result.Minute() != 30 ||
		result.Second() != 10 {
		t.Error("time.Time is not the correct value")
	}
}
