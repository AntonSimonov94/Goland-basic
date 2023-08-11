package main

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село"}, // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик"}, // города с населением 1000 тыс. человек и более
	 }
	 cityPopulation := map[string]int{
		"Город" : 50,
		"Село" : 100,
		"Миллионик" : 500,
	 }
	result := false
	for value1, _ := range cityPopulation {
		fmt.Print(value1)
		for _, value2 := range groupCity[100] {
			fmt.Print(value2)
			if value1 == value2 {
				result = true
				break 
			} else {result = false}

		}
		if result == false {
			delete(cityPopulation, value1)
		}
	}
	
	fmt.Print(cityPopulation)
	
}

