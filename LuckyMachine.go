package main

import (
	"fmt"
	"math/rand"
)

// choose picks a random element from the slice
func choose(nums []int) int {
	return nums[rand.Intn(len(nums))]
}

func main() {
	// start with {1, 2, 3}
	choices := []int{1, 2, 3}

	// now choose from the remaining set
	fmt.Println("Random choice:", choose(choices))
}
