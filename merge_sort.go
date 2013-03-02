/**
 * merge sort with gorutines
 *
 * @auther leo
 * 2013/3/2 14:20
 */
package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	a := []int{10, 31, 2, 30, 4, 132, 3, 12, 43, 8, 0, 7, 9}
	// a := []int{4, 3, 1, 2, 7, 0}
	fmt.Println("Before sort: ", a)
	mergeSort(a)
	fmt.Println("After  sort: ", a)
}

func merge(a []int, p, q, r int) error {
	if a == nil || p < 0 || p >= q || q >= r || len(a) < r {
		return errors.New("Invalid input!")
	}
	i, j, k := p, q, 0
	b := make([]int, r-p)
	for ; i < q && j < r; k++ {
		if a[i] <= a[j] {
			b[k] = a[i]
			i++
		} else {
			b[k] = a[j]
			j++
		}
	}
	for ; i < q; i, k = i+1, k+1 {
		b[k] = a[i]
	}
	for ; j < r; j, k = j+1, k+1 {
		b[k] = a[j]
	}
	for s, t := p, 0; s < r; s, t = s+1, t+1 {
		a[s] = b[t]
	}
	return nil
}

func mergeSort(a []int) {
	if a == nil || len(a) < 2 {
		return
	}
	var swait sync.WaitGroup
	for s, n := 2, len(a); (s >> 1) < n; s <<= 1 {
		for i := 0; i < n; i += s {
			swait.Add(1)
			go func(i int) {
				defer swait.Done()
				var r = i + s
				if r > n {
					r = n
				}
				merge(a, i, i+(s>>1), r)
			}(i)
		}
		swait.Wait()
	}
}
