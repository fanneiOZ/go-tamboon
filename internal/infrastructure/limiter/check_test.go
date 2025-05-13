package limiter

import (
	"fmt"
	"math"
	"testing"
)

func compareTriplets(a []int32, b []int32) []int32 {
	result := make([]int32, 2, 2)
	for i := range a {
		diffBToA := float64(b[i] - a[i])
		if diffBToA == 0 {
			continue
		}

		result[int((diffBToA/math.Abs(diffBToA)+1)/2)]++
	}

	return result
}

func sort(arr []int32) []int32 {
	size := len(arr)
	if size == 1 {
		return arr
	}

	if size == 2 {
		temp := arr[0]
		if temp > arr[1] {
			arr[0] = arr[1]
			arr[1] = temp
		}

		return arr
	}

	split := len(arr) / 2
	left := sort(arr[:split])
	right := sort(arr[split:])

	if right[0] > left[len(left)-1] {
		return append(left, right...)
	}

	return append(right, left...)
}

func TestMath(t *testing.T) {
	t.Run("test math", func(t *testing.T) {
		//result := compareTriplets([]int32{10, 3, 32}, []int32{1, 33, 2})

		//t.Logf("%v", result)
		//fmt.Printf("%*s\n", 2, strings.Repeat("#", 1))
		//
		//fmt.Println(math.Ceil(float64(0) / 2))

		fmt.Println(sort([]int32{10, 2, 9, 3, 1}))
	})
}
