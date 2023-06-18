/*
    PROGRAM :  DNA
    AUTHOR  :  Laura Purcell-Gates
    EMAIL   :  laurapurcellgates@gmail.com

    DNA is a program that identifies a person based on their DNA.

    This is my conversion into Go of the Python program I wrote for Harvardx CS50P problem set 6, with added data validation and error handling. I also replaced the CS50 longest_match function with *. I reused the database and sequence files provided by CS50.

		Link to my original Python version:
		https://github.com/laurapg1214/Harvard-CS50-lab-problemsets-Python/tree/main/CS50P-a-pset6-dna

		Link to the CS50 problem set:
		https://cs50.harvard.edu/x/2023/psets/6/dna/  */


package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// check for command-line usage
	if len(os.Args) != 3 {
		panic("Usage: dna.go database sequence")
	}

	// read database & sequence files into variables
	csvFilename := os.Args[1]
	txtFilename := os.Args[2]

	// open sequence file for reading, assing to variable
	sequence, err := os.ReadFile(txtFilename)

	// error checking - file unreadable
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	// initialize slice for names
	names := []string{}

	// create map for csv file
	csvFile, err := os.Open(csvFilename)

	// error checking - file unreadable
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer csvFile.Close()

	// CSVToMap takes a reader and returns an array of maps, using the header row as the keys
	// adapted from https://gist.github.com/drernie/5684f9def5bee832ebc50cabb46c377a

	// read csv data
	//csvReader := csv.NewReader(csvFile)
	//csvReader.FieldsPerRecord = -1 // allow variable number of fields
	//csvReader.Columns, err = csvReader.ReadHeader()
	//if err != nil {
		//fmt.Println(err)
		//return
	//}

	csvReader := csv.NewReader(csvFile)
	CSVToMap(csvFile, csvReader)

	//csvData, err := csvReader.ReadAll()

	// error checking
	//if err != nil {
		//panic(err)
	//}

	// TESTING print one record
	//fmt.Println(csvData[1]["name"])

	// TESTING print csv data
	//for _, row := range csvData {
		//for _, col := range row {
			//fmt.Printf("%s,", col)
		//}
		//fmt.Println()
	//}
	fmt.Print(string(sequence))
	//fmt.Println(names)

	// loop through csv file to append info to names[]
	//for _, row := range csvData {
		//for _, col := range row {
			//names = append(names, col)
		//}
	//}

	// initialize slice to store STR counts
	STRcounts := []int{}

	// TESTING print STRcounts
	fmt.Println(STRcounts)

	// initialize map for names
	// STRmap := map[string]int {
		
	// }

}

func CSVToMap(csvFile *os.File, csvReader *csv.Reader) []map[string]string {
	rows := []map[string]string{}
	var header []string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			csvMap := map[string]string{}
			for i := range header {
				csvMap[header[i]] = record[i]
			}
			rows = append(rows, csvMap)
		}
	}
	return rows
}
