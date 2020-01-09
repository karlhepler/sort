package sort

func Merge(items []int) []int {
	numItems := len(items)
	if numItems <= 1 {
		return items
	}

	left := Merge(append([]int{}, items[:numItems/2]...))
	right := Merge(append([]int{}, items[numItems/2:]...))

	return merge(left, right, items)
}

func merge(left, right, items []int) []int {
	numLeft, numRight := len(left), len(right)
	i, l, r := 0, 0, 0
	for l < numLeft && r < numRight {
		if left[l] <= right[r] {
			items[i] = left[l]
			l++
		} else {
			items[i] = right[r]
			r++
		}
		i++
	}

	for l < numLeft {
		items[i] = left[l]
		i, l = i+1, l+1
	}
	for r < numRight {
		items[i] = right[r]
		i, r = i+1, r+1
	}

	return items
}
