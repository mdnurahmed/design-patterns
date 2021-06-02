package strategy

import "math"

type BubbleSortStrategy struct {
}

func (b *BubbleSortStrategy) sort(collection []int) {
	for i := 0; i < len(collection); i++ {
		for j := i + 1; j < len(collection); j++ {
			if collection[i] > collection[j] {
				//swap
				collection[i], collection[j] = collection[j], collection[i]
			}
		}
	}
}

type CountingSortStrategy struct {
}

func (c *CountingSortStrategy) sort(collection []int) {
	min := math.MaxInt32
	max := math.MinInt32
	freq := make(map[int]int)
	for i := 0; i < len(collection); i++ {
		freq[collection[i]]++
		if collection[i] > max {
			max = collection[i]
		}
		if collection[i] < min {
			min = collection[i]
		}
	}
	index := 0
	for i := min; i <= max; i++ {
		cnt := freq[i]
		for j := 1; j <= cnt; j++ {
			collection[index] = i
			index++
		}
	}
}
