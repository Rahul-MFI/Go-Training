package main

import (
	"fmt"
)

func main() {
	subjects := [4]string{"Math", "Science", "English", "History"}

	sortedSubjects := make([]string, len(subjects))
	copy(sortedSubjects, subjects[:])

	customSortFuncString(sortedSubjects)

	students := map[string][]int{
		"Rocky": {85, 90, 78, 88},
		"Rahul": {92, 81, 74, 95},
	}

	for name, marks := range students {
		fmt.Printf("Marks for %s:\n", name)
		total := 0
		for _, subj := range sortedSubjects {
			idx := indexOf(subjects[:], subj)
			fmt.Printf("%s: %d\n", subj, marks[idx])
			total += marks[idx]
		}
		fmt.Printf("Average: %.2f\n", float64(total)/float64(len(marks)))

		sortedMarks := make([]int, len(marks))
		copy(sortedMarks, marks)
		customSortFunc(sortedMarks)
		fmt.Printf("Sorted marks: %v\n\n", sortedMarks)
	}

	for _, subj := range sortedSubjects {
		highest := -1
		topStudent := ""
		for name, marks := range students {
			idx := indexOf(subjects[:], subj)
			if marks[idx] > highest {
				highest = marks[idx]
				topStudent = name
			}
		}
		fmt.Printf("Highest scorer in %s: %s (%d)\n", subj, topStudent, highest)
	}
}

func indexOf(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func customSortFuncString(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func customSortFunc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
