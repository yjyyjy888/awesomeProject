package main

import (
	"fmt"
)

func main() {
	m1 := map[string]int{
		"sss":   22,
		"ffdfd": 4,
	}
	//noteFrequency := map[string]float32 {
	//	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	//	"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	s1 := make([]int, 5)
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1, 2, 3, 5)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(m1)
}
