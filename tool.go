package gotool


// Permutations P(n) = n! => Get All Combination of int Array
// Example: input [1,2,3], out [[1,2,3],[2,1,3],[2,3,1],[3,1,2],[3,2,1],[1,3,2]] , it's len = 3!
func Permutations(arr ...int) [][]int {
	var permuteFunc func([]int, int)
	var res [][]int

	permuteFunc = func(arr []int, n int) {
		if n == 1 {
			// 宣告最後一個狀態資料
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				permuteFunc(arr, n-1)
				if n%2 == 1 {
					// swap for odd
					tmp := arr[i]
					arr[i] = arr[n-1] // n=3, i=0, [1,2,3,4] => [3,1,2,4]
					arr[n-1] = tmp
				} else {
					// swap n for even
					tmp := arr[0]
					arr[0] = arr[n-1] // n=4, [1,2,3,4] => [4,1,2,3]
					arr[n-1] = tmp
				}
			}
		}
	}
	permuteFunc(arr, len(arr))
	return res
}

