package FMRepeater

import (
	"encoding/json"
	"io"
	"os"
)

type FMRepeater struct {
	Sat      string `json:"SAT"`
	Mode     string `json:"MODE"`
	Uplink   string `json:"UPLINK"`
	Downlink string `json:"DOWNLINK"`
}

var FMRepeaters []FMRepeater

func init() {
	f, err := os.Open("Repeaters.json")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	if err != nil {

	}
	bytes, err := io.ReadAll(f)
	if err != nil {

	}
	err = json.Unmarshal(bytes, &FMRepeaters)
	if err != nil {

	}
}
