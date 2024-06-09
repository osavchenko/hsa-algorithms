package countring_sort

func CountingSort(arr []int) []int {
	m := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
		}
	}

	cnt := make([]int, m+1)

	for _, v := range arr {
		cnt[v]++
	}

	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}

	out := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		cnt[arr[i]]--
		out[cnt[arr[i]]] = arr[i]
	}

	return out
}
