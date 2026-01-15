package searching

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // 찾았다!
		} else if arr[mid] < target {
			// 중간값이 목표보다 작다면 -> 오른쪽(큰 쪽)을 봐야 함
			low = mid + 1
		} else {
			// 중간값이 목표보다 크다면 -> 왼쪽(작은 쪽)을 봐야 함
			high = mid - 1
		}
	}
	return -1 // 반복문이 끝날 때까지 못 찾음
}
