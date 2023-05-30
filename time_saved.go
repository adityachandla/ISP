package main

import "fmt"

var minutesSaved = [...]uint{5, 10, 15, 20, 30, 40}
var numPeople = [...]uint{10, 25, 50, 70, 100}

const SALARY_EUROS_YEARLY = 50_000
const WORKER_HOURS_PER_WEEK = 40
const WORKING_DAYS_PER_WEEK = 7
const FEES_YEARLY = 36_000

type ValueGetter func(uint, uint) float32

// This function returns the number of hours saved per week
func computeTotalTime(minutes uint, people uint) float32 {
	totalMinSavedDaily := (float32)(minutes * people)
	totalHoursSavedDaily := totalMinSavedDaily / 60
	return totalHoursSavedDaily * WORKING_DAYS_PER_WEEK
}

func computeWorkersSaved(minutes uint, people uint) float32 {
	hrsPerWeek := computeTotalTime(minutes, people)
	return hrsPerWeek / WORKER_HOURS_PER_WEEK
}

// Returns money saved in thousand euros
func computeMoneySaved(minutes uint, people uint) float32 {
	workersSaved := computeWorkersSaved(minutes, people)
	return (workersSaved * SALARY_EUROS_YEARLY) / 1000
}

func computeRoi(minutes uint, people uint) float32 {
	moneySaved := computeMoneySaved(minutes, people) * 1000
	return ((moneySaved - FEES_YEARLY) / FEES_YEARLY) * 100
}

func main() {
	printTable(computeTotalTime, "Hours saved per week")
	printTable(computeWorkersSaved, "Workers Saved")
	printTable(computeMoneySaved, "Money Saved (Thousand euros)")
	printTable(computeRoi, "Return on Investment")
}

func printTable(valueGetter ValueGetter, title string) {
	fmt.Println(title)
	printMinutes()
	for _, p := range numPeople {
		fmt.Printf("%8d ", p)
		for _, m := range minutesSaved {
			val := valueGetter(p, m)
			fmt.Printf("%7.1f ", val)
		}
		fmt.Println()
	}
}

func printMinutes() {
	fmt.Printf("%-8s", "")
	for _, minute := range minutesSaved {
		fmt.Printf("%8d", minute)
	}
	fmt.Println()
}
