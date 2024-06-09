package main

import (
	"fmt"
	"github.com/osavchenko/hsa-algorithms/balanced_tree"
	"github.com/osavchenko/hsa-algorithms/countring_sort"
	"math/rand"
	"slices"
)

func main() {
	btree()
	sort()
}

func btree() {
	fmt.Println("Balanced binary tree: number of operations on a custom sets")

	ins := make([]int, 100)
	find := make([]int, 100)
	del := make([]int, 100)

	for i := 0; i < 100; i++ {
		set := rand.Perm(100)
		root := &balanced_tree.TreeNode{}

		root, ins[i] = balanced_tree.Build(set)
		f := set[rand.Int()%100]
		_, find[i] = balanced_tree.Find(root, f, 0)
		_, del[i] = balanced_tree.Delete(root, f, 0)
	}

	fmt.Printf(
		"Min number of operations for insert: %d\nMax number of operations for insert: %d\nMean number of operations for insert: %.2f\n===\n",
		slices.Min(ins),
		slices.Max(ins),
		mean(ins),
	)

	fmt.Printf(
		"Min number of operations for search: %d\nMax number of operations for search: %d\nMean number of operations for search: %.2f\n===\n",
		slices.Min(find),
		slices.Max(find),
		mean(find),
	)

	fmt.Printf(
		"Min number of operations for delete: %d\nMax number of operations for search: %d\nMean number of operations for search: %.2f\n\n\n",
		slices.Min(del),
		slices.Max(del),
		mean(del),
	)
}

func mean(data []int) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0
	for _, d := range data {
		sum += d
	}
	return float64(sum) / float64(len(data))
}

func sort() {
	set := rand.Perm(100)
	result := countring_sort.CountingSort(set)

	fmt.Println("Counting sort: result on a custom set")
	fmt.Printf(
		"Initial set:\n%v\nResult:\n%v\n",
		set,
		result,
	)
}
