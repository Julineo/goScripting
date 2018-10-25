// $ go run main.go > output.csv

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
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

	lines, err := csv.NewReader(fi).ReadAll()
	if err != nil {
		panic(err)
	}

	// Output to Stdout
	w := csv.NewWriter(os.Stdout)
	// Windows
	w.UseCRLF = true

	// If email is here pass it
	for _, record := range lines {
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
