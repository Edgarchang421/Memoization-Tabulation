package main

import "fmt"

func main() {
	data := Constructer(10)
	fmt.Printf("11: %v\n", data.Get(10))
	fmt.Printf("10: %v\n", data.Data)
	fmt.Printf("11: %v\n", data.Get(11))
	fmt.Printf("10: %v\n", data.Data)

	fmt.Printf("tabulation Fibonacci: %v\n", tabulationFibonacci(10))
}

// memoization製作dp[]，便於往後查詢，針對time做優化
type Fibonacci struct {
	Data []int
}

func Constructer(n int) Fibonacci {
	data := make([]int, n+1)
	data[0], data[1] = 1, 1
	for i := 2; i <= n; i++ {
		data[i] = data[i-1] + data[i-2]
	}

	return Fibonacci{Data: data}
}
func (f *Fibonacci) Get(n int) int {
	if n < len(f.Data) {
		return f.Data[n]
	} else {
		originalLen := len(f.Data)
		f.Data = append(f.Data, make([]int, n-len(f.Data)+1)...)
		for i := originalLen; i < len(f.Data); i++ {
			f.Data[i] = f.Data[i-1] + f.Data[i-2]
		}
		return f.Data[n]
	}
}

// tabulation只暫存需要的資料，針對space做優化
func tabulationFibonacci(n int) int {
	last1 := 1
	last2 := 1
	var result int
	for i := 2; i <= n; i++ {
		result = last1 + last2
		last1, last2 = result, last1
	}
	return result
}
