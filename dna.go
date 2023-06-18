/*
    PROGRAM :  DNA
    AUTHOR  :  Laura Purcell-Gates
    EMAIL   :  laurapurcellgates@gmail.com

    DNA is a program that identifies a person based on their DNA.

    This is my conversion into Go of the Python program I wrote for Harvardx CS50P problem set 6, with added data validation and error handling. I also converted and adapted CS50's longest_match function. I reused the database and sequence files provided by CS50.

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
	"math"
	"os"
	"strconv"
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

	// create map for csv file
	csvFile, err := os.Open(csvFilename)

	// error checking - file unreadable
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer csvFile.Close()

	// read csv data into array of maps
	csvReader := csv.NewReader(csvFile)
	csvMaps := CSVToMap(csvFile, csvReader)
	
	// loop through maps, find longest match of database STRs in sequence
	matches := 0
	strMatches := 0
	personStrCount := 0
	
	for person := range csvMaps {
		csvPerson := csvMaps[person]

		for key := range csvPerson {
			if key != "name" {
				// count consecutive occurrences of str in sequence
				strMatches = longestMatch(string(sequence), string(key))
				personStrCount, err = strconv.Atoi(csvPerson[key])
				if err != nil {
					log.Fatal(err)
				}
				// check for match
				if strMatches == personStrCount {
					matches++
				} else {
					matches = 0
					break
				}
			}
		}
		if matches == len(csvPerson) - 1 {
			fmt.Println(csvPerson["name"])
			return
		} 
	}
	fmt.Println("No match")
}

// CSVToMap takes a reader and returns an array of maps, using the header row as the keys
// adapted from https://gist.github.com/drernie/5684f9def5bee832ebc50cabb46c377a
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

// conversion to Go & slight adaptation of function written by CS50 for problem set
func longestMatch (sequence string, key string) int {
	
	// initialize variables
	var longestRun float64 = 0
	strLength := len(key)
	sequenceLength := len(string(sequence))

	// check each character in sequence for most consecutive runs of subsequence
	for i := 0; i < sequenceLength; i++ {
		// initialize count of consecutive runs
		var count float64 = 0

		// check for str (key) match
		// if match, move substring to next potential match in sequence
		// continue until out of consecutive matches
		for {
			// adjust substring start and end
			start := i + int(count) * strLength
			if start + strLength > sequenceLength {
				break
			}
			end := start + strLength

			// if match in the substring
			if sequence[start:end] == key {
				count++
			} else {
				break
			}	
		}
		// update most consecutive matches found
		longestRun = math.Max(longestRun, count)
	}
	return int(longestRun)
}
