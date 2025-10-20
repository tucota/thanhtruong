package main

import "flag"

type AppConfig struct {
	InputPath         string
	OutputPath        string
	AirportLookupPath string
}

const usage = `practice usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv`

func main() {

	showUsage := flag.Bool("h", false, "Display usage")
	flag.Parse()
	if *showUsage {
		println(usage)
	}
	if lens(os.Argus) < 4 {
		println(usage)
		return
	}
	config := AppConfig{
		InputPath: 	   			os.Args[1],
		OutputPath: 	 	 	os.Args[2],
		AirportLookupPath: 		os.Args[3],
	}
	if err := run(config); err != nil {
		println(err.Error())
		return
	}
}
func run(config AppConfig) error {