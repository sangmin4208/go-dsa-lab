package sorting

// QuickSort: 외부에서 호출하는 함수
func QuickSort(arr []int) {
	// 재귀 함수 호출 (전체 범위)
	quickSortRecursive(arr, 0, len(arr)-1)
}

// 재귀적으로 동작하는 내부 로직
func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		// pivot 위치를 잡고, 정렬된 pivot의 인덱스를 받아옴
		pivotIndex := partition(arr, low, high)

		// pivot 기준 왼쪽 부분 정렬 (재귀)
		quickSortRecursive(arr, low, pivotIndex-1)

		// pivot 기준 오른쪽 부분 정렬 (재귀)
		quickSortRecursive(arr, pivotIndex+1, high)
	}
}

// partition: 피벗을 기준으로 작은 값은 왼쪽, 큰 값은 오른쪽으로 보내는 핵심 로직
func partition(arr []int, low, high int) int {
	pivot := arr[high] // 가장 오른쪽 요소를 피벗으로 선택
	i := low - 1       // i는 "피벗보다 작은 값들이 저장될 마지막 위치"

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			// 교환 (Swap)
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 피벗을 자신의 자리(작은 값들과 큰 값들 사이)로 이동
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1 // 피벗의 최종 위치 반환
}
