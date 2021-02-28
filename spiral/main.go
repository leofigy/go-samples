package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("(%s) missing number size , square matrix\n", os.Args[0])
		os.Exit(1)
	}

	size, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Panic(err)
	}

	data, err := spiral(size)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, array := range data {
		fmt.Printf("%+v\n", array)
	}
}

func spiral(n int) (result [][]int, err error) {

	if n < 1 {
		err = fmt.Errorf("invalid matrix size %d", n)
		return
	}

	result = make([][]int, n, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n, n)
	}

	// sequence generator
	val := 0
	incr := func() int {
		val++
		return val
	}

	// How to transverse the array
	populate(result, incr, 0, 0, n, n)
	return result, nil
}

func populate(data [][]int, incr func() int, i, j, m, n int) {
	if i >= m || j >= n {
		return
	}
	// first row
	for val := j; val < n; val++ {
		data[i][val] = incr()
	}

	// last column
	for val := i + 1; val < m; val++ {
		data[val][n-1] = incr()
	}

	// first reow and last not the same
	if (m - 1) != i {
		for val := n - 2; val >= j; val-- {
			data[m-1][val] = incr()
		}
	}

	if (n - 1) != j {
		for val := m - 2; val > i; val-- {
			data[val][j] = incr()
		}
	}

	populate(data, incr, i+1, j+1, m-1, n-1)
}
