package main

func AllIntPairs(a []int) [][2]int {
	var collate [][2]int

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			collate = append(collate, [2]int{a[i], a[j]})
		}
	}

	return collate
}
