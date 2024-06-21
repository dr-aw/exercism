/*
 In this exercise you're going to write some code to help you cook a brilliant lasagna from your favorite cooking book.
 You have four tasks, all related to the time spent cooking the lasagna.
*/

package main

import "fmt"

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return actualMinutesInOven + (numberOfLayers * 2)
}

func main() {
	fmt.Printf("%v\t%v\t%v\t", RemainingOvenTime(10), PreparationTime(2), ElapsedTime(2, 10))
}
