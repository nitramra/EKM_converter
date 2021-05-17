package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

//ViisiArray global array
var ViisiArray ViisStruct

//RytmiArray global array
var RytmiArray RytmStruct

func main() {

	f, err := os.Create("output.csv")
	//Load CSV files into arrays
	LoadArrays()

	//RytmiMap rütmitüübi otsimiseks
	RytmiMap := make(map[string]string)

	for _, r := range RytmiArray {
		RytmiMap[r.Fkey] = r.Rytmiliik
	}
	//Create ABC encoding for each string
	for _, q := range ViisiArray {
		var viisPrimary string
		var viisAlterna string
		switch RytmiMap[q.FKey] {
		case "1":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + CN(q.P16, 2)
		case "5":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + "2 " + CN(q.P4, 1) + "2 | " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + "2 " + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + "2 " + CN(q.P12, 1) + "2 " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2 " + CN(q.P16, 1) + "2 "
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + "2 " + CN(q.P4, 2) + "2 | " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + "2 " + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + "2 " + CN(q.P12, 2) + "2 " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2 " + CN(q.P16, 2) + "2 "
		case "1/2A":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2 " + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2 " + CN(q.P16, 2) + "2"
		case "2/1":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + " H" + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + " H" + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + CN(q.P16, 2)
		case "2":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + " H" + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + " H" + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + " H" + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + " H" + CN(q.P16, 2) + "2"
		case "2/2A":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + " H" + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2 " + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + " H" + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2 " + CN(q.P16, 2) + "2"
		case "3A":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + " H" + CN(q.P4, 1) + "2 " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + " H" + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + " H" + CN(q.P12, 1) + "2 " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + " H" + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + " H" + CN(q.P4, 2) + "2 " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + " H" + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + " H" + CN(q.P12, 2) + "2 " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + " H" + CN(q.P16, 2) + "2"
		case "2A":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + "2 " + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2 " + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + "2 " + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2 " + CN(q.P16, 2) + "2"
		case "4":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + "2 " + CN(q.P3, 1) + CN(q.P4, 1) + "2 " + CN(q.P5, 1) + CN(q.P6, 1) + "2 " + CN(q.P7, 1) + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + "2 " + CN(q.P11, 1) + CN(q.P12, 1) + "2 " + CN(q.P13, 1) + CN(q.P14, 1) + "2 " + CN(q.P15, 1) + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + "2 " + CN(q.P3, 2) + CN(q.P4, 2) + "2 " + CN(q.P5, 2) + CN(q.P6, 2) + "2 " + CN(q.P7, 2) + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + "2 " + CN(q.P11, 2) + CN(q.P12, 2) + "2 " + CN(q.P13, 2) + CN(q.P14, 2) + "2 " + CN(q.P15, 2) + CN(q.P16, 2) + "2"
		case "9":
			viisPrimary = CN(q.P1, 1) + "2 " + CN(q.P2, 1) + " " + CN(q.P3, 1) + "2 " + CN(q.P4, 1) + " " + CN(q.P5, 1) + "2 " + CN(q.P6, 1) + " " + CN(q.P7, 1) + "2 " + CN(q.P8, 1) + " | " + CN(q.P9, 1) + "2 " + CN(q.P10, 1) + " " + CN(q.P11, 1) + "2 " + CN(q.P12, 1) + " " + CN(q.P13, 1) + "2 " + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2 " + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + "2 " + CN(q.P2, 2) + " " + CN(q.P3, 2) + "2 " + CN(q.P4, 2) + " " + CN(q.P5, 2) + "2 " + CN(q.P6, 2) + " " + CN(q.P7, 2) + "2 " + CN(q.P8, 2) + " | " + CN(q.P9, 2) + "2 " + CN(q.P10, 2) + " " + CN(q.P11, 2) + "2 " + CN(q.P12, 2) + " " + CN(q.P13, 2) + "2 " + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2 " + CN(q.P16, 2)
		case "13":
			viisPrimary = CN(q.P1, 1) + "/" + CN(q.P2, 1) + "/ " + CN(q.P3, 1) + "/" + CN(q.P4, 1) + "/ " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + "/" + CN(q.P10, 1) + "/ " + CN(q.P11, 1) + "/" + CN(q.P12, 1) + "/ " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + "/" + CN(q.P2, 2) + "/ " + CN(q.P3, 2) + "/" + CN(q.P4, 2) + "/ " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + "/" + CN(q.P10, 2) + "/ " + CN(q.P11, 2) + "/" + CN(q.P12, 2) + "/ " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + CN(q.P16, 2)
		case "1/2":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + " H" + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + " H" + CN(q.P16, 2) + "2"
		case "1B":
			viisPrimary = CN(q.P1, 1) + ">" + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + ">" + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + ">" + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + ">" + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + CN(q.P16, 2)
		case "1E/2B":
			viisPrimary = CN(q.P1, 1) + ">" + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + ">" + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + ">" + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + " " + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + ">" + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + ">" + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + ">" + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + " " + CN(q.P16, 2) + "2"
		case "9B/9F":
			viisPrimary = CN(q.P1, 1) + ">" + CN(q.P2, 1) + " " + CN(q.P3, 1) + ">" + CN(q.P4, 1) + " " + CN(q.P5, 1) + ">" + CN(q.P6, 1) + " " + CN(q.P7, 1) + ">" + CN(q.P8, 1) + " | " + CN(q.P9, 1) + ">" + CN(q.P10, 1) + " " + CN(q.P11, 1) + ">" + CN(q.P12, 1) + " " + CN(q.P13, 1) + ">" + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + ">" + CN(q.P2, 2) + " " + CN(q.P3, 2) + ">" + CN(q.P4, 2) + " " + CN(q.P5, 2) + ">" + CN(q.P6, 2) + " " + CN(q.P7, 2) + ">" + CN(q.P8, 2) + " | " + CN(q.P9, 2) + ">" + CN(q.P10, 2) + " " + CN(q.P11, 2) + ">" + CN(q.P12, 2) + " " + CN(q.P13, 2) + ">" + CN(q.P14, 2) + " " + CN(q.P15, 2) + CN(q.P16, 2)
		case "5B":
			viisPrimary = CN(q.P1, 1) + ">" + CN(q.P2, 1) + " " + CN(q.P3, 1) + "2 " + CN(q.P4, 1) + "2 | " + CN(q.P5, 1) + ">" + CN(q.P6, 1) + " " + CN(q.P7, 1) + "2 " + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + ">" + CN(q.P10, 1) + " " + CN(q.P11, 1) + "2 " + CN(q.P12, 1) + "2 | " + CN(q.P13, 1) + ">" + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2 " + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + ">" + CN(q.P2, 2) + " " + CN(q.P3, 2) + "2 " + CN(q.P4, 2) + "2 | " + CN(q.P5, 2) + ">" + CN(q.P6, 2) + " " + CN(q.P7, 2) + "2 " + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + ">" + CN(q.P10, 2) + " " + CN(q.P11, 2) + "2 " + CN(q.P12, 2) + "2 | " + CN(q.P13, 2) + ">" + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2 " + CN(q.P16, 2) + "2"
		case "1A":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2"
		case "1B/2":
			viisPrimary = CN(q.P1, 1) + ">" + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + ">" + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + CN(q.P16, 2)
		case "6":
			viisPrimary = CN(q.P1, 1) + "/" + CN(q.P2, 1) + "/ " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " (" + CN(q.P7, 1) + CN(q.P8, 1) + ") | " + CN(q.P9, 1) + "/" + CN(q.P10, 1) + "/ " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " (" + CN(q.P15, 1) + CN(q.P16, 1) + ")"
			viisAlterna = CN(q.P1, 2) + "/" + CN(q.P2, 2) + "/ " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " (" + CN(q.P7, 2) + CN(q.P8, 2) + ") | " + CN(q.P9, 2) + "/" + CN(q.P10, 2) + "/ " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " (" + CN(q.P15, 2) + CN(q.P16, 2) + ")"
		case "9A":
			viisPrimary = CN(q.P1, 1) + " | " + CN(q.P2, 1) + "2 " + CN(q.P3, 1) + " " + CN(q.P4, 1) + "2 " + CN(q.P5, 1) + " " + CN(q.P6, 1) + "2 " + CN(q.P7, 1) + " " + CN(q.P8, 1) + "2 | " + CN(q.P9, 1) + " | " + CN(q.P10, 1) + "2 " + CN(q.P11, 1) + " " + CN(q.P12, 1) + "2 " + CN(q.P13, 1) + " " + CN(q.P14, 1) + "2 " + CN(q.P15, 1) + " " + CN(q.P16, 1) + "2"
			viisAlterna = CN(q.P1, 2) + " | " + CN(q.P2, 2) + "2 " + CN(q.P3, 2) + " " + CN(q.P4, 2) + "2 " + CN(q.P5, 2) + " " + CN(q.P6, 2) + "2 " + CN(q.P7, 2) + " " + CN(q.P8, 2) + "2 | " + CN(q.P9, 2) + " | " + CN(q.P10, 2) + "2 " + CN(q.P11, 2) + " " + CN(q.P12, 2) + "2 " + CN(q.P13, 2) + " " + CN(q.P14, 2) + "2 " + CN(q.P15, 2) + " " + CN(q.P16, 2) + "2"
		case "1 ja 1A":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2"
		case "1 ja 9":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + CN(q.P16, 1)
			viisAlterna = CN(q.P1, 2) + "2 " + CN(q.P2, 2) + " " + CN(q.P3, 2) + "2 " + CN(q.P4, 2) + " " + CN(q.P5, 2) + "2 " + CN(q.P6, 2) + " " + CN(q.P7, 2) + "2 " + CN(q.P8, 2) + " | " + CN(q.P9, 2) + "2 " + CN(q.P10, 2) + " " + CN(q.P11, 2) + "2 " + CN(q.P12, 2) + " " + CN(q.P13, 2) + "2 " + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2 " + CN(q.P16, 2)
		case "1/1A":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " " + CN(q.P15, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " " + CN(q.P15, 2) + "2"
		case "1/3":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " " + CN(q.P7, 1) + CN(q.P8, 1) + " | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " H" + CN(q.P15, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " " + CN(q.P7, 2) + CN(q.P8, 2) + " | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " H" + CN(q.P15, 2) + "2"
		case "3":
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + " " + CN(q.P3, 1) + CN(q.P4, 1) + " " + CN(q.P5, 1) + CN(q.P6, 1) + " H" + CN(q.P7, 1) + "2 | " + CN(q.P9, 1) + CN(q.P10, 1) + " " + CN(q.P11, 1) + CN(q.P12, 1) + " " + CN(q.P13, 1) + CN(q.P14, 1) + " H" + CN(q.P15, 1) + "2"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + " " + CN(q.P3, 2) + CN(q.P4, 2) + " " + CN(q.P5, 2) + CN(q.P6, 2) + " H" + CN(q.P7, 2) + "2 | " + CN(q.P9, 2) + CN(q.P10, 2) + " " + CN(q.P11, 2) + CN(q.P12, 2) + " " + CN(q.P13, 2) + CN(q.P14, 2) + " H" + CN(q.P15, 2) + "2"

		default:
			viisPrimary = CN(q.P1, 1) + CN(q.P2, 1) + CN(q.P3, 1) + CN(q.P4, 1) + CN(q.P5, 1) + CN(q.P6, 1) + CN(q.P7, 1) + CN(q.P8, 1) + CN(q.P9, 1) + CN(q.P10, 1) + CN(q.P11, 1) + CN(q.P12, 1) + CN(q.P13, 1) + CN(q.P14, 1) + CN(q.P15, 1) + CN(q.P16, 1) + "; manual"
			viisAlterna = CN(q.P1, 2) + CN(q.P2, 2) + CN(q.P3, 2) + CN(q.P4, 2) + CN(q.P5, 2) + CN(q.P6, 2) + CN(q.P7, 2) + CN(q.P8, 2) + CN(q.P9, 2) + CN(q.P10, 2) + CN(q.P11, 2) + CN(q.P12, 2) + CN(q.P13, 2) + CN(q.P14, 2) + CN(q.P15, 2) + CN(q.P16, 2) + "; manual"

		}

		if viisPrimary == viisAlterna {

			_, err := f.WriteString(q.FKey + ";1;" + RytmiMap[q.FKey] + ";" + viisPrimary + "\n")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err := f.WriteString(q.FKey + ";1;" + RytmiMap[q.FKey] + ";" + viisPrimary + "\n" + q.FKey + ";2;" + RytmiMap[q.FKey] + ";" + viisAlterna + "\n")
			if err != nil {
				log.Fatal(err)
			}

		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

//CN on funktsioon mis konverteerib Accessi kirje ABC kujule vastavaks
func CN(noot string, index int) string {
	var nootClean string
	var tmpval string
	var vastus string
	var variatsioon bool
	//Kui väli on tühi siis saadame tühja vastuse
	if len(noot) == 0 {
		return "<empty>"
	}
	//Alternatiivide puhul võtame välja esimese ja teise alternatiivi
	if index == 1 && string(noot[len(noot)-1:]) == "]" {
		nootClean = (strings.Split(noot, "["))[0]
	} else if index == 2 && string(noot[len(noot)-1:]) == "]" {
		tmpval = (strings.Split(noot, "["))[1]
		nootClean = string(tmpval[:len(tmpval)-1])
		variatsioon = true
	} else {
		nootClean = noot
	}

	//Ehitame vastuse vastavalt sellele mis elemendid meile sattusid

	if strings.Contains(nootClean, "/") == true {
		//Kui on kiirem noot siis teeme nii
		vastus = ConvertNoot(strings.Split(nootClean, "/")[0]) + "/" + ConvertNoot(strings.Split(nootClean, "/")[1]) + "/"
	} else if strings.Contains(nootClean, "(") == true && strings.Contains(nootClean, ")") == true {
		//Kui on eellöök siis teeme nii
		vastus = "{" + ConvertNoot(strings.Split(strings.Split(nootClean, "(")[1], ")")[0]) + "}" + ConvertNoot(strings.Split(nootClean, ")")[1])
	} else {
		vastus = ConvertNoot(nootClean)
	}

	if variatsioon == true {
		vastus = "!mark!" + vastus
	}

	return vastus
}

//ConvertNoot on funktsioon mis annab helikõrguse vaste ABC kujul
func ConvertNoot(noot string) string {
	var diees bool
	var bekaar bool
	var nootClean string
	var nootOutput string

	if strings.Contains(noot, "#") == true {
		nootClean = strings.Trim(noot, "#")
		diees = true
	} else if strings.Contains(noot, "-bekaar") == true {
		nootClean = strings.Trim(noot, "-bekaar")
		bekaar = true
	} else {
		nootClean = noot
	}

	switch nootClean {
	case "a":
		nootOutput = "A"
	case "b":
		nootOutput = "B"
	case "c":
		nootOutput = "C"
	case "d":
		nootOutput = "D"
	case "e":
		nootOutput = "E"
	case "f":
		nootOutput = "F"
	case "g":
		nootOutput = "G"
	case "h":
		nootOutput = "B"
	case "2a":
		nootOutput = "a"
	case "2b":
		nootOutput = "b"
	case "2c":
		nootOutput = "c"
	case "2d":
		nootOutput = "d"
	case "2e":
		nootOutput = "e"
	case "2f":
		nootOutput = "f"
	case "2g":
		nootOutput = "g"
	case "2h":
		nootOutput = "b"
	case "1a":
		nootOutput = "A,"
	case "1b":
		nootOutput = "B,"
	case "1c":
		nootOutput = "C,"
	case "1d":
		nootOutput = "D,"
	case "1e":
		nootOutput = "E,"
	case "1f":
		nootOutput = "F,"
	case "1g":
		nootOutput = "G,"
	case "1h":
		nootOutput = "B,"
	case "ba":
		nootOutput = "_A"
	case "bh":
		nootOutput = "_B"
	case "bd":
		nootOutput = "_D"
	case "be":
		nootOutput = "_E"
	}
	if diees == true {
		return "^" + nootOutput
	}
	if bekaar == true {
		return "=" + nootOutput
	}
	if len(nootOutput) > 0 {
		return nootOutput
	}
	return "<" + noot + ">"
}

//LoadArrays function that Loads files into arrays
func LoadArrays() {
	//SetCSVReader params
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = ';'
		return r // Allows use dot as delimiter and use quotes in CSV
	})

	//Viis.csv
	ViisiFile, err := os.Open("Viis.csv")
	if err != nil {
		panic(err)
	}
	defer ViisiFile.Close()

	if err := gocsv.UnmarshalFile(ViisiFile, &ViisiArray); err != nil {
		panic(err)
	}
	//rytmityybid.csv
	RytmiFile, err := os.Open("rytmityybid.csv")
	if err != nil {
		panic(err)
	}
	defer RytmiFile.Close()

	if err := gocsv.UnmarshalFile(RytmiFile, &RytmiArray); err != nil {
		panic(err)
	}

}
