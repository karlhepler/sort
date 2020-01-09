package sort

func Quick(items []int, start, end int) []int {
	if start >= end {
		return items
	}

	pivot := partition(items, start, end)
	Quick(items, start, pivot-1)
	Quick(items, pivot+1, end)

	return items
}

func partition(items []int, start, end int) int {
	p := start
	for i := start; i < end; i++ {
		if items[i] <= items[end] {
			swap(items, i, p)
			p++
		}
	}
	swap(items, p, end)
	return p
}

func swap(items []int, a, b int) {
	tmp := items[a]
	items[a] = items[b]
	items[b] = tmp
}
