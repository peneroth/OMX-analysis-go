package main

import (
	"bufio"
	"os"
	"strings"
	"time"
)

func loadOmxData(filename string) []OMX_Data_t {
	var OMX_Data []OMX_Data_t

	// Determine number of rows and valid days in the dataset
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	nbrOfDays := 0
	row := 0
	for scanner.Scan() {
		row++
		// Determine if the first item on the row is a date, i.e. format "xxxx-xx-xx"
		line := strings.Split(scanner.Text(), ";")
		isDate := isValidDate(line[0])
		// Determine if it is a weekday
		isWeekday := false
		if isDate {
			date := getDate(line[0])
			weekday := date.Weekday() // Sunday = 0, Saturday = 6
			if int(weekday) >= 1 && int(weekday) <= 5 {
				isWeekday = true
			}
		}
		// Determine that OMX values are non-zero
		isNonzero := false
		if isDate && isWeekday {
			priceHigh, priceLow, priceClosure := getPrices(line[1], line[2], line[3])
			isNonzero = priceHigh > 0.1 && priceLow > 0.1 && priceClosure > 0.1
		}
		// Count valid rows
		if isDate && isWeekday && isNonzero {
			nbrOfDays++
		}
	}
	f.Close()
	// fmt.Println("nbrOfDays = ", nbrOfDays)

	OMX_Data = make([]OMX_Data_t, nbrOfDays)

	// Load data from file
	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner = bufio.NewScanner(f)
	row = 0
	position := nbrOfDays
	for scanner.Scan() {
		row++
		// Determine if the first item on the row is a date, i.e. format "xxxx-xx-xx"
		line := strings.Split(scanner.Text(), ";")
		isDate := isValidDate(line[0])
		// Determine if it is a weekday
		isWeekday := false
		var date time.Time
		if isDate {
			date = getDate(line[0])
			weekday := date.Weekday() // Sunday = 0, Saturday = 6
			if int(weekday) >= 1 && int(weekday) <= 5 {
				isWeekday = true
			}
		}
		// Determine that OMX values are non-zero
		isNonzero := false
		var priceHigh float64 = 0.0
		var priceLow float64 = 0.0
		var priceClosure float64 = 0.0
		if isDate && isWeekday {
			priceHigh, priceLow, priceClosure = getPrices(line[1], line[2], line[3])
			isNonzero = priceHigh > 0.1 && priceLow > 0.1 && priceClosure > 0.1
		}
		// If a valid row, add data to OMX_data
		if isDate && isWeekday && isNonzero {
			position--
			OMX_Data[position].date = date
			OMX_Data[position].weekday = date.Weekday() // Sunday = 0, Saturday = 6
			OMX_Data[position].priceHigh = priceHigh
			OMX_Data[position].priceLow = priceLow
			OMX_Data[position].priceClosure = priceClosure
		}
	}
	f.Close()
	// fmt.Println("position = ", position)
	// fmt.Println("nbrOfDays = ", nbrOfDays)

	return OMX_Data
}
