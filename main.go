package main

import (
	"CSV2ADIF/Record"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toADIF(records []string) Record.ADIFRecord {
	record := Record.ADIFRecord{
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
