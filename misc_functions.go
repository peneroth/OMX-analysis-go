package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func isValidDate(s string) bool {
	line := strings.Split(s, "-")
	year, err := strconv.Atoi(line[0])
	if err != nil {
		return false
	}
	if year < 1900 || year > 2050 {
		return false
	}
	month, err := strconv.Atoi(line[1])
	if err != nil {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	day, err := strconv.Atoi(line[2])
	if err != nil {
		return false
	}
	if day < 1 || month > 31 {
		return false
	}
	return true
}

// Check if a valid date before using getDate()
func getDate(s string) time.Time {
	line := strings.Split(s, "-")
	year, err := strconv.Atoi(line[0])
	if err != nil {
		fmt.Println("strconv.Atoi error")
		panic(err)
	}
	month, err := strconv.Atoi(line[1])
	if err != nil {
		fmt.Println("strconv.Atoi error")
		panic(err)
	}
	day, err := strconv.Atoi(line[2])
	if err != nil {
		fmt.Println("strconv.Atoi error")
		panic(err)
	}
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date
}

func getPrices(sHigh string, sLow string, sClosure string) (priceHigh float64, priceLow float64, priceClosure float64) {
	sHigh = strings.Replace(sHigh, ",", ".", -1)
	sLow = strings.Replace(sLow, ",", ".", -1)
	sClosure = strings.Replace(sClosure, ",", ".", -1)
	priceHigh, err := strconv.ParseFloat(sHigh, 64)
	if err != nil {
		priceHigh = 0.0
	}
	priceLow, err = strconv.ParseFloat(sLow, 64)
	if err != nil {
		priceLow = 0.0
	}
	priceClosure, err = strconv.ParseFloat(sClosure, 64)
	if err != nil {
		priceClosure = 0.0
	}
	return
}
