package arrays

func ArraySum(arr []int) int {
	sum := 0
	// for i:=0; i<len(arr);i++{
	// 	sum += arr[i]
	// }
	for _, number := range arr { // rangeでインデックス, 値
		sum += number
	}
	return sum
}

func AllSumArr(numbersToSum ...[]int) (sums []int) { // 可変長引数
	// length := len(numbersToSum)
	// sums = make([]int, length) // 第二引数ぶんスライスを作成

	// for i, numbers := range numbersToSum{
	// 	sums[i] = ArraySum(numbers)
	// }
	for _, numbers := range numbersToSum {
		sums = append(sums, ArraySum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 { // 空配列の場合
			sums = append(sums, 0)
		} else {
			sums = append(sums, ArraySum(numbers[1:]))
		}
	}
	return sums
}
