package main

import "fmt"

func Powerset(array []int) [][]int {
	sets := [][]int{
		[]int{},
	}

	for i := 0; i < len(array); i++ {
		lenSets := len(sets)
		for j := 0; j < lenSets; j++ {

			// copy newSubSet value with the sets
			newSubSet := []int{}
			newSubSet = append(newSubSet, sets[j]...)

			// insert number to newSubSet
			newSubSet = append(newSubSet, array[i])

			// insert newSubSet to the sets
			sets = append(sets, newSubSet)

		}
	}

	return sets
}

func main() {
	fmt.Println("this is a main driver dummy")
}
