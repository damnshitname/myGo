package main

import (
	"fmt"
	"time"
)

var cacheArray [50]uint64

func main() {
	//var arr [5]int
	//for i := 0;i<len(arr);i++ {
	//	fmt.Printf("array[%d] is: %d\n", i, arr[i])
	//}
	//for i, i2 := range arr {
	//	fmt.Printf("array[%d] is: %d\n", i, i2)
	//}

	//a := [...]string{"a", "b", "c", "d"}
	//for i := range a {
	//	fmt.Println("Array item", i, "is", a[i])
	//}
	//
	var arr [15]int
	//var arr1 = new([5]int)
	//fmt.Println(arr)
	//fmt.Println(arr1)

	//arr3 := &arr
	//arr4 := *arr1
	//arr3[2] = 100
	//arr4[2] = 111
	//fmt.Println(arr[2])
	//fmt.Println(arr1[2])
	//fmt.Println(arr3[2])
	//fmt.Println(arr4[2])

	//写一个循环并用下标给数组赋值（从 0 到 15）
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}
	fmt.Println(arr)

	//通过数组我们可以更快的计算出 Fibonacci 数
	fmt.Println("普通递归----->")
	start := time.Now()
	for i := 0; i < 45; i++ {
		fmt.Println(fibonacci(i))
	}
	cost := time.Now().Sub(start)
	fmt.Println("cost:", cost)

	fmt.Println("数组做缓存递归----->")
	start = time.Now()
	for i := 0; i < 45; i++ {
		fmt.Println(fibonacciWithArrayCache(i))
	}
	cost = time.Now().Sub(start)
	fmt.Println("cost:", cost)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacciWithArrayCache(n int) (res uint64) {
	if cacheArray[n] != 0 {
		res = cacheArray[n]
		return res
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciWithArrayCache(n-1) + fibonacciWithArrayCache(n-2)
	}
	cacheArray[n] = res
	return res
}
