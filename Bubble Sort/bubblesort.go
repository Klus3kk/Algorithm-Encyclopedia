func bubblesort(array[] int)[]int {
	bool swapped = false;
	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - 1 - i; j++ {
			if (array[j] > array[j + 1]) {
				array[j], array[j + 1] = array[j + 1], array[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}