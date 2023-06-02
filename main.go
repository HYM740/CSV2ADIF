package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type ADIFRecord struct {
	Call   string
	Band   string
	Mode   string
	Date   string
	Time   string
	Freq   string //非必须
	BandRx string //非必须
	FreqRx string //非必须
	Prop   string //非必须
	Sat    string //非必须
}

type FMRepeater struct {
	Sat      string `json:"SAT"`
	Mode     string `json:"MODE"`
	Uplink   string `json:"UPLINK"`
	Downlink string `json:"DOWNLINK"`
}

var FMRepeaters []FMRepeater

func LoadFMRepeaters() {
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

func (r ADIFRecord) String() string {
	var ADIFstr = fmt.Sprintf("<CALL:%d>%s\n   <BAND:%d>%s\n   <MODE:%d>%s\n   <QSO_DATE:%d>%s\n   <TIME_ON:%d>%s\n",
		len(r.Call), r.Call,
		len(r.Band), r.Band,
		len(r.Mode), r.Mode,
		len(r.Date), r.Date,
		len(r.Time), r.Time)
	if r.Freq != "" {
		ADIFstr = ADIFstr + fmt.Sprintf("   <FREQ:%d>%s\n",
			len(r.Freq), r.Freq)
	}
	if r.BandRx != "" {
		ADIFstr = ADIFstr + fmt.Sprintf("   <BAND_RX:%d>%s\n",
			len(r.BandRx), r.BandRx)
	}
	if r.FreqRx != "" {
		ADIFstr = ADIFstr + fmt.Sprintf("   <FREQ_RX:%d>%s\n",
			len(r.FreqRx), r.FreqRx)
	}
	if r.Prop != "" {
		ADIFstr = ADIFstr + fmt.Sprintf("   <PROP_MODE:%d>%s\n",
			len(r.Prop), r.Prop)
	}
	if r.Sat != "" {
		ADIFstr = ADIFstr + fmt.Sprintf("   <SAT_NAME:%d>%s\n",
			len(r.Sat), r.Sat)
	}
	ADIFstr = ADIFstr + "<EOR>"
	return ADIFstr
}

// 修正数据
func (r *ADIFRecord) FuckExcel() {
	dates := strings.Split(r.Date, "/")
	if len(dates[1]) < 2 {
		dates[1] = "0" + dates[1]
	}
	if len(dates[2]) < 2 {
		dates[2] = "0" + dates[2]
	}
	timestr := strings.Split(r.Time, ":")
	if len(timestr[0]) < 2 {
		timestr[0] = "0" + timestr[1]
	}
	if len(timestr[1]) < 2 {
		timestr[1] = "0" + timestr[1]
	}
	if len(timestr[2]) < 2 {
		timestr[2] = "0" + timestr[2]
	}
	r.Date = fmt.Sprintf("%s%s%s", dates[0], dates[1], dates[2])
	r.Time = fmt.Sprintf("%s%s%s", timestr[0], timestr[1], timestr[2])
}

func (r *ADIFRecord) CompleteSatInfo() {
	if r.Prop == "SAT" && r.Sat != "" {
		for _, v := range FMRepeaters {
			if r.Sat == v.Sat {
				r.Freq = v.Uplink
				r.FreqRx = v.Downlink
			}
		}
	}
}

func toADIF(records []string) ADIFRecord {
	record := ADIFRecord{
		Call:   records[0],
		Band:   records[1],
		Mode:   records[2],
		Date:   records[3],
		Time:   records[4],
		Freq:   records[5],
		BandRx: records[6],
		FreqRx: records[7],
		Prop:   records[8],
		Sat:    records[9],
	}
	record.FuckExcel()
	record.CompleteSatInfo()
	return record
}

func main() {
	LoadFMRepeaters()
	//filename := os.Args[1]
	//f, err := os.Open(filename)
	f, err := os.Open("qsl.csv")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		fmt.Println(toADIF(lines))
	}

}
