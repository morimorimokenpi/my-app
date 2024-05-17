package main

import "fmt"

// 同じ形式が続く場合は= iotaを省力できる
// ここでは本来 Orange = iota, Banana = iotaという意味
const	(
	Apple = iota
	Orange
	Banana
)

func main() {
	// fmt.Println(Apple)
	// fmt.Println(Orange)
	// fmt.Println(Banana)


	// 配列とスライス
	// 配列は固定長、スライスは可変長
	var a [3]int
	fmt.Println(a)

	b  := []int{1, 2, 3}
	fmt.Println(b)


	c := make([]int, 3)
	fmt.Println(c)
	arr1 := [2][3]int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(arr1)

	arr2 := [][]int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(arr2)

	arr3 := []int{1, 2, 3}
	fmt.Println(arr3)

	arr3 = append(arr3, 4, 5, 6)
	fmt.Println(arr3)
	fmt.Printf("aの長さは %d\n", len(arr3))

	// for i := 0; i < len(12); i++ {
	// 	fmt.Println(i)
	// }
		array1 := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			array1 = append(array1, i)
		}
		fmt.Println(array1)

		// アロケーションとはメモリの割当を指す
		// a2 := make([]int, 0 , len(array1) / 2)
		// for i := 0; i < len(array1); i++ {
		// 	if i % 2 == 0 {
		// 		a2 = append(a2, array1[i])
		// 	}
		// }
		// array1 = a2

		// fmt.Println(array1)


		n := 50
		// array1[:n]はindexの0からn-1(49)までを表す
		// array1[n+1:]はindexの51(52)から最後までを表す
		// 50を削除している
		// array1 = append(array1[:n], array1[n+1:]...)
		// fmt.Println(array1[:n])
		// fmt.Println(array1[n+1:])
		// fmt.Println(array1)

		// array1[:n]はindexの0からn-1(49)までを表す
		// copy(array1[n:], array1[n+1:])はarray1[n+1:]をarray1[n:]にコピーしている
		array1 = array1[:n+copy(array1[n:], array1[n+1:])]
		fmt.Println(array1)
}
