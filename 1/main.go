package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fileInput := "bsd-yo7Fqs.csv"
	fileExclude := "cnstid.csv"

	// Map for emails to exclude
	m := map[string]struct{}{}

	// Open files
	fi, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}

	fx, err := os.Open(fileExclude)
	if err != nil {
		panic(err)
	}
	defer fx.Close()

	// Read file into a Variable
	linesEx, err := csv.NewReader(fx).ReadAll()
	if err != nil {
		panic(err)
	}

	// Loop through lines & put emails in map
	for _, line := range linesEx {
		m[line[8]] = struct{}{}
	}


	lines, err := csv.NewReader(fi).ReadAll()
	if err != nil {
		panic(err)
	}

	// Create output file
	fo, err := os.Create("noEISid.csv")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	writer := csv.NewWriter(fo)
	defer writer.Flush()

	for _, line := range lines {
		line = append(line, "\r\n")
		err := writer.Write(line)
		if err != nil {
			fmt.Println(err)
		}
	}

}
