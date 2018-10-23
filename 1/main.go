package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	// Map for emails to exclude
	m := map[string]struct{}{}

	// Open files
	fi, err := os.Open("bsd-yo7Fqs.csv")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fx, err := os.Open("cnstid.csv")
	if err != nil {
		panic(err)
	}
	defer fx.Close()

	// Read file into a variable
	linesEx, err := csv.NewReader(fx).ReadAll()
	if err != nil {
		panic(err)
	}

	// Loop through lines & put emails in map
	for _, line := range linesEx {
		m[line[0]] = struct{}{}
	}

/*
	r := csv.NewReader(fi)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		_, ok := m[record[4]]
		if !ok {
			fmt.Println(record)
		}
	}
*/

	lines, err := csv.NewReader(fi).ReadAll()
	if err != nil {
		panic(err)
	}

	// Output to Stdout
	w := csv.NewWriter(os.Stdout)
	// Windows
	w.UseCRLF = true

	for _, record := range lines {
		fmt.Println(record)
		_, ok := m[record[4]]
		if ok {
			continue
		}

		if err := w.Write(record); err != nil {
			fmt.Println("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
