package main

import (
	"fmt"
	"sort"
)

func main() {
	getInputData()
}

func getInputData() {
	var n, want int

	fmt.Println("tentukan panjang array:")
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Println("masukan nilai bilangan ke", i)
		fmt.Scan(&a[i])
	}

	fmt.Println("masukan bilangan yang ingin dicari:")
	sort.Ints(a)
	fmt.Println(a)
	fmt.Scan(&want)
	sortAndSearch(n, want, a)

}

func sortAndSearch(n, want int, a []int) {
	start := 0
	end := n - 1
	for {
		mid := (start + end) / 2
		if want < a[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
		if !(end >= start) && (a[mid] != want) {
			fmt.Println("Data tidak ada dalam array")
			break
		} else {
			if a[mid] == want {
				fmt.Println("Data berada di index nomor", mid)
				break
			}
		}
	}
}
