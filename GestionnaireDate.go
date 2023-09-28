package main

import (
	"fmt"
	"os"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func NewDate(year, month, day int) Date {
	return Date{Year: year, Month: month, Day: day}
}

func (d Date) Conv() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d Date) Bissextile() bool {
	return (d.Year%4 == 0 && d.Year%100 != 0) || (d.Year%400 == 0)
}

func error() {

}

func Error(str string) {
	fmt.Println(str)
	os.Exit(0)
}

func ErrorToDate(MM int, DD int) {
	if MM < 0 || MM > 13 {
		Error("les mois ne sont pas bon")
	}

	if ((MM == 4 || MM == 6 || MM == 9 || MM == 11) && (DD > 30 || DD < 0)) || (MM == 2 && (DD > 29 || DD < 0)) {
		Error("les jours ne sont pas bon")
	}
}

func isValidDate(year, month, day int) bool {

	return true
}

func main() {

	Date := NewDate(2042, 12, 10)
	ErrorToDate(Date.Month, Date.Day)

	fmt.Println("Date :", Date.Conv())
	fmt.Println("Bissextile : ", Date.Bissextile())
}
