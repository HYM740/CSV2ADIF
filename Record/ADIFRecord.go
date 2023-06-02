package Record

import (
	"CSV2ADIF/FMRepeater"
	"fmt"
	"strconv"
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
	if len(timestr) < 3 {
		timestr = append(timestr, "00")
	}
	if len(timestr[2]) < 2 {
		timestr[2] = "0" + timestr[2]
	}
	r.Date = fmt.Sprintf("%s%s%s", dates[0], dates[1], dates[2])
	r.Time = fmt.Sprintf("%s%s%s", timestr[0], timestr[1], timestr[2])
}

func (r *ADIFRecord) CompleteSatInfo() {
	if r.Prop == "SAT" && r.Sat != "" {
		for _, v := range FMRepeater.FMRepeaters {
			if r.Sat == v.Sat {
				r.Freq = v.Uplink
				r.FreqRx = v.Downlink
			}
		}
	}
	r.CompleteBandInfo()
}

func (r *ADIFRecord) CompleteBandInfo() {
	freq, err := strconv.ParseFloat(r.Freq, 64)
	if err != nil {

	}
	if freq2band(freq) != "" {
		r.Band = freq2band(freq)
	}
	freqrx, err := strconv.ParseFloat(r.Freq, 64)
	if err != nil {

	}
	if freq2band(freq) != "" {
		r.BandRx = freq2band(freqrx)
	}
}
func freq2band(freq float64) string {
	if freq > 0.135 && freq < 0.138 {
		return "2190M"

	} else if freq > 0.472 && freq < 0.479 {
		return "630M"

	} else if freq > 1.8 && freq < 2 {
		return "160M"

	} else if freq > 3.5 && freq < 4 {
		return "80M"

	} else if freq > 5.25 && freq < 5.45 {
		return "60M"

	} else if freq > 7 && freq < 7.3 {
		return "40M"

	} else if freq > 10.1 && freq < 10.15 {
		return "30M"

	} else if freq > 14 && freq < 14.35 {
		return "20M"

	} else if freq > 18.068 && freq < 18.168 {
		return "17M"

	} else if freq > 21 && freq < 21.45 {
		return "15M"

	} else if freq > 24.89 && freq < 24.99 {
		return "12M"

	} else if freq > 28 && freq < 29.7 {
		return "10M"

	} else if freq > 50 && freq < 54 {
		return "6M"

	} else if freq > 70 && freq < 71 {
		return "4M"

	} else if freq > 144 && freq < 148 {
		return "2M"

	} else if freq > 220 && freq < 225 {
		return "1.25M"

	} else if freq > 420 && freq < 450 {
		return "70CM"

	} else if freq > 902 && freq < 928 {
		return "33CM"

	} else if freq > 1240 && freq < 1300 {
		return "23CM"

	} else if freq > 2300 && freq < 2450 {
		return "13CM"

	} else if freq > 3300 && freq < 3500 {
		return "9CM"

	} else if freq > 5650 && freq < 5925 {
		return "6CM"

	} else if freq > 10000 && freq < 10500 {
		return "3CM"

	} else if freq > 24000 && freq < 24250 {
		return "1.25CM"

	} else if freq > 47000 && freq < 47200 {
		return "6MM"

	} else if freq > 75500 && freq < 81000 {
		return "4MM"

	} else if freq > 122250 && freq < 123000 {
		return "2.5MM"

	} else if freq > 142000 && freq < 149000 {
		return "2MM"

	} else if freq > 241000 && freq < 250000 {
		return "1MM"

	} else if freq > 300000 && freq < 306000 {
		return "SUBMM"
	}
	return ""
}
