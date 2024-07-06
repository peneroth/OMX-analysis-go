//
// Analyze OMX data
// https://www.nasdaqomxnordic.com/index/historiska_kurser?Instrument=SE0000337842
//

package main

import (
	"fmt"
	"os"
	"time"
)

type OMX_Data_t struct {
	date         time.Time
	weekday      time.Weekday
	priceHigh    float64
	priceLow     float64
	priceClosure float64
}

func calcValueWeekday(OMX_Data []OMX_Data_t, startDate_s string, endDate_s string) []float64 {
	var accOMXweekday [5]float64

	// Find start index
	var startDate time.Time
	if isValidDate(startDate_s) {
		startDate = getDate(startDate_s)
	} else {
		fmt.Println("invalid start date")
		os.Exit(0)
	}
	index := 0
	// fmt.Println(len(OMX_Data))
	for (startDate.Compare(OMX_Data[index].date) == 1) && (index < len(OMX_Data)-1) {
		index++
	}
	fmt.Println(OMX_Data[index].date)
	if index > len(OMX_Data)-60 {
		fmt.Println("Start date after or to close to last date in dataset")
		os.Exit(0)
	}

	// Determine end date
	var endDate time.Time
	if isValidDate(endDate_s) {
		endDate = getDate(endDate_s)
	} else {
		fmt.Println("invalid end date")
		os.Exit(0)
	}

	// find first Monday
	for OMX_Data[index].date.Weekday() != 1 {
		index++
		// fmt.Println(OMX_Data[index].date.Weekday())
	}
	fmt.Println(OMX_Data[index].date)
	fmt.Println(OMX_Data[index].date.Weekday())

	// Loop over all weeks
	for index < len(OMX_Data)-5 && endDate.Compare(OMX_Data[index].date) == 1 {
		// Loop over one week
		for i := 0; i < 5; i++ {
			accOMXweekday[i] += 1000.0 / OMX_Data[index].priceClosure
			// if OMX is closed on the actual day, buy next open day =>
			// step to next day if needed (handle closed days)
			if int(OMX_Data[index].date.Weekday()) == i+1 {
				index++
			} // else {
			// fmt.Println("non-tradeing day", OMX_Data[index].date)
			// }

		}

	}
	return accOMXweekday[:]
}

func calcValueDay(OMX_Data []OMX_Data_t, startDate_s string, endDate_s string) []float64 {
	var accOMXday [31]float64

	// Find start index
	var startDate time.Time
	if isValidDate(startDate_s) {
		startDate = getDate(startDate_s)
	} else {
		fmt.Println("invalid start date")
		panic(0)
	}
	index := 0
	for startDate.Compare(OMX_Data[index].date) == 1 {
		index++
	}
	// fmt.Println(index)
	// fmt.Println(OMX_Data[index].date)

	// find first trading day in the month
	// == the day before was in previous month
	// One exeption, if index == 0, use this day
	if index > 0 {
		for OMX_Data[index].date.Month() == OMX_Data[index-1].date.Month() {
			index++
			// fmt.Println(OMX_Data[index].date.Weekday())
		}
	}
	fmt.Println(index)
	fmt.Println(OMX_Data[index].date)

	// Determine end date
	var endDate time.Time
	if isValidDate(endDate_s) {
		endDate = getDate(endDate_s)
	} else {
		fmt.Println("invalid end date")
		os.Exit(0)
	}

	// Loop over all months till end date
	for index < len(OMX_Data)-50 && endDate.Compare(OMX_Data[index].date) == 1 {
		// Loop over all days in the month
		for i := 0; i < 31; i++ {
			accOMXday[i] += 1000.0 / OMX_Data[index].priceClosure
			// if OMX is closed on the actual day, buy next open day =>
			// step to next day if needed (handle closed days)
			if int(OMX_Data[index].date.Day()) == i+1 {
				index++
			} // else {
			// fmt.Println("non-tradeing day", OMX_Data[index].date)
			// }

		}

	}
	fmt.Println("Last used date, or the next trading day", OMX_Data[index].date)
	return accOMXday[:]
}

func main() {

	// filename := "OmxData/_SE0000337842_2023-12-25.csv"
	filename := "OmxData/OMX_20000103-20231222.csv"
	// filename := "OmxData/OMX_20000103-20091230.csv"
	// filename := "OmxData/OMX_20100104-20191230.csv"
	// filename := "OmxData/OMX_20200102-20231222.csv"

	OMX_Data := loadOmxData(filename)

	accOMXweekday := calcValueWeekday(OMX_Data, "2010-01-01", "2030-01-01")
	fmt.Println(accOMXweekday)
	accOMXweekday = calcValueWeekday(OMX_Data, "2010-04-01", "2015-01-01")
	fmt.Println(accOMXweekday)

	accOMXday := calcValueDay(OMX_Data, "2000-01-05", "2015-01-01")
	fmt.Println(accOMXday)
	accOMXday = calcValueDay(OMX_Data, "2011-01-01", "2015-01-01")
	fmt.Println(accOMXday)
}
