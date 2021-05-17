package main

//ViisStruct Viis.csv struktuur
type ViisStruct []struct {
	IDviis string `csv:"IDviis"`
	FKey   string `csv:"FKey"`
	VNr    string `csv:"VNr"`
	P1     string `csv:"P1"`
	P2     string `csv:"P2"`
	P3     string `csv:"P3"`
	P4     string `csv:"P4"`
	P5     string `csv:"P5"`
	P6     string `csv:"P6"`
	P7     string `csv:"P7"`
	P8     string `csv:"P8"`
	P9     string `csv:"P9"`
	P10    string `csv:"P10"`
	P11    string `csv:"P11"`
	P12    string `csv:"P12"`
	P13    string `csv:"P13"`
	P14    string `csv:"P14"`
	P15    string `csv:"P15"`
	P16    string `csv:"P16"`
}

//YldStruct Üld.csv struktuur
type YldStruct []struct {
	ID   string `csv:"ID"`
	Rytm string `csv:"Rytm"`
}

//RytmiStruct rytmityybid.csv struktuur
type RytmStruct []struct {
	Fkey      string `csv:"Fkey"`
	Nr        string `csv:"Nr"`
	Rytmiliik string `csv:"Rytmiliik"`
}
