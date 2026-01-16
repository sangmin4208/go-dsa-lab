package sorting

// MergeSort: 외부에서 호출하는 함수
func MergeSort(arr []int) {
	// 배열의 길이가 1 이하면 이미 정렬된 상태
	if len(arr) <= 1 {
		return
	}
	
	// 임시 배열 생성 (병합에 사용)
	temp := make([]int, len(arr))
	mergeSortRecursive(arr, temp, 0, len(arr)-1)
}

// 재귀적으로 동작하는 내부 로직
func mergeSortRecursive(arr []int, temp []int, left, right int) {
	if left < right {
		// 중간 지점 찾기
		mid := left + (right-left)/2
		
		// 왼쪽 부분 정렬 (재귀)
		mergeSortRecursive(arr, temp, left, mid)
		
		// 오른쪽 부분 정렬 (재귀)
		mergeSortRecursive(arr, temp, mid+1, right)
		
		// 정렬된 두 부분 병합
		merge(arr, temp, left, mid, right)
	}
}

// merge: 정렬된 두 부분 배열을 병합하는 핵심 로직
func merge(arr []int, temp []int, left, mid, right int) {
	// 임시 배열에 현재 범위 복사
	for i := left; i <= right; i++ {
		temp[i] = arr[i]
	}
	
	i := left      // 왼쪽 부분 배열의 시작 인덱스
	j := mid + 1   // 오른쪽 부분 배열의 시작 인덱스
	k := left      // 병합될 배열의 인덱스
	
	// 두 부분 배열을 비교하며 병합
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}
	
	// 왼쪽 부분 배열에 남은 요소 복사
	for i <= mid {
		arr[k] = temp[i]
		i++
		k++
	}
	
	// 오른쪽 부분 배열에 남은 요소는 이미 제자리에 있으므로 복사 불필요
}
