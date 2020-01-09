package sort

func Bubble(items []int) []int {
	numItems := len(items)
	if numItems <= 1 {
		return items
	}

	for i := 0; i < numItems-1; i++ {
		for j := 1; j < numItems-i; j++ {
			if items[j] < items[j-1] {
				tmp := items[j]
				items[j] = items[j-1]
				items[j-1] = tmp
			}
		}
	}

	return items
}
