package sort

func Insertion(items []int) []int {
	numItems := len(items)
	if numItems <= 1 {
		return items
	}

	for i := 1; i < numItems; i++ {
		j := i - 1
		tmp := items[i]
		for ; j >= 0; j-- {
			if items[j] <= tmp {
				break
			}
			items[j+1] = items[j]
		}
		items[j+1] = tmp
	}

	return items
}
