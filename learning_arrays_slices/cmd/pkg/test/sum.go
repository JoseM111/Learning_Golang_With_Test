package test

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */

func SumArray(numArray [5]int) int {
	// ™ ____________
	var sum int
	
	/* range lets you iterate over an array. On each iteration,
	 * range returns two values - the index and the value. We are
	 * choosing to ignore the index value by using _ blank identifier. */
	for _, num := range numArray {
		sum += num
	}
	
	// Same as above
	// for i := 0; i < len(numArray); i++ {
	//     sum += numArray[i]
	// }
	
	return sum
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func SumSlice(numSlice []int) int {
	// ™ ____________
	var sum int
	
	/* range lets you iterate over an array. On each iteration,
	 * range returns two values - the index and the value. We are
	 * choosing to ignore the index value by using _ blank identifier. */
	for _, num := range numSlice {
		sum += num
	}
	
	// Same as above
	// for i := 0; i < len(numSlice); i++ {
	//     sum += numSlice[i]
	// }
	
	return sum
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func SumAllSlice(numberSliceSum ...[]int) []int {
	// ™ ____________
	var sums []int
	
	for _, nums := range numberSliceSum {
		sums = append(sums, SumSlice(nums))
	}
	
	return sums
	// lengthOfNums := len(numberSliceSum)
	// sums := make([]int, lengthOfNums)
	//
	// for i, num := range numberSliceSum {
	// 	sums[i] = SumSlice(num)
	// }
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func SumAllTailSlices(numberSliceSum ...[]int) []int {
	// ™ ____________
	// declare a new slice of ints which defaults to = []int{} || nil
	var sums []int
	
	for _, nums := range numberSliceSum {
		// #..........
		if len(nums) != 0 {
			tail := nums[1:]
			// will only some at index 1 until the end of the slice
			sums = append(sums, SumSlice(tail))
			// ™ ____________
		} else { sums = append(sums, 0) }
	}
	
	return sums
}
/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */

















