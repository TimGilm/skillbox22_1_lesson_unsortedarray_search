// поиск в несортированном массиве, числа входящего в массив и не входящегов него.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10 //сколько элементов в массиве, в котором будем делать сортировку

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10 * n)
	}
	//найдем число находящееся в массиве:
	value := a[n/2]
	index := find(a, value)
	fmt.Printf("Индекс: %v\n", index)
	//теперь надем число не находящееся в массиве
	value = n * 20
	index = find(a, value)
	fmt.Printf("Индекс: %v\n", index)
}
func find(a [n]int, value int) (index int) {
	index = -1
	for i := 0; i < n; i++ {
		if a[i] == value {
			index = i
			break
		}
	}
	return
}
