package main

import (
	"fmt"
	"strings"
)

//This is for Capstone 3 (Temperature Tables)

type celsius float64
type fahrenheit float64

func fToC(f float64) float64 {
	return float64(fahrenheit(f).celsius())
}

func cToF(c float64) float64 {
	return float64(celsius(c).fahrenheit())
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * (5.0 / 9.0))
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func drawRow(left float64, right float64) {
	fmt.Printf("| %5.1f | %5.1f | \n", left, right)
}

func drawDivider(length int) {
	fmt.Printf("%s \n", strings.Repeat("=", length))
}

func drawHeader(left string, right string) {
	fmt.Printf("| %-5s | %-5s | \n", left, right)
}

func drawTable(leftHeader string, rightHeader string, conv func(x float64) float64) {

	fmt.Printf("\nConverting from %s to %s:\n", leftHeader, rightHeader)
	drawDivider(17)
	drawHeader(leftHeader, rightHeader)
	drawDivider(17)
	for i := 40.0; i <= 100; i += 5 {
		drawRow(i, conv(i))
	}
	drawDivider(17)

}

func runTemperature() {

	drawTable("C", "F", cToF)

	drawTable("F", "C", fToC)

}
