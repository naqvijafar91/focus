package focus

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeMarshal(t *testing.T) {
	loc, _ := time.LoadLocation("")
	time1 := time.Date(2020, 5, 2, 0, 0, 0, 0, loc)
	time2 := time.Date(2020, 7, 25, 0, 0, 0, 0, loc)
	time3 := time.Date(2020, 12, 28, 0, 0, 0, 0, loc)
	sample1 := Time{time1}
	sample2 := Time{time2}
	sample3 := Time{time3}
	ser, err := sample1.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	var strTime string
	err = json.Unmarshal(ser, &strTime)
	if strTime != "02-05-2020" {
		t.Error("Failed to parse")
	}
	ser, err = sample2.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	err = json.Unmarshal(ser, &strTime)
	if strTime != "25-07-2020" {
		t.Error("Failed to parse")
	}
	ser, err = sample3.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	err = json.Unmarshal(ser, &strTime)
	if strTime != "28-12-2020" {
		t.Error("Failed to parse")
	}
}

func TestTimeUnMarshal(t *testing.T) {
	str1 := "25-06-2020"
	sample1 := &Time{}
	err := sample1.UnmarshalJSON([]byte(str1))
	if err != nil {
		t.Error(err)
		return
	}
	if sample1.Year() != 2020 || sample1.Month() != 6 || sample1.Day() != 25 {
		t.Error("Failed to unmarshal properly")
	}
	str2 := "05-12-2000"
	sample2 := &Time{}
	err = sample2.UnmarshalJSON([]byte(str2))
	if err != nil {
		t.Error(err)
		return
	}
	if sample2.Year() != 2000 || sample2.Month() != 12 || sample2.Day() != 5 {
		t.Error("Failed to unmarshal properly")
	}
}

func TestSpecialUnmarshal(t *testing.T) {
	sample3 := &Time{}
	// ""05-05-2020"" this adds extra " at beginning and end
	err := sample3.UnmarshalJSON([]byte{34, 48, 53, 45, 48, 53, 45, 50, 48, 50, 48, 34})
	if err != nil {
		t.Error(err)
		return
	}
	if sample3.Year() != 2020 || sample3.Month() != 5 || sample3.Day() != 5 {
		t.Error("Failed to unmarshal properly")
	}
}

func TestTimeUnmarshalNilOrEmpty(t *testing.T) {
	sample1 := &Time{}
	err := sample1.UnmarshalJSON([]byte(nil))
	err = sample1.UnmarshalJSON(nil)
	err = sample1.UnmarshalJSON([]byte(""))
	if err != nil {
		t.Error("Should not throw error", err)
	}
}
