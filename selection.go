package sort

func Selection(items []int) []int {
	numItems := len(items)

	if numItems <= 1 {
		return items
	}

	for i := 0; i < numItems; i++ {
		minIndex := i
		for j := i + 1; j < numItems; j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}
		tmp := items[i]
		items[i] = items[minIndex]
		items[minIndex] = tmp
	}

	return items
}
