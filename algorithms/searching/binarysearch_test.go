package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	sortedData := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	tests := []struct {
		name     string
		target   int
		expected int // 찾은 인덱스 (못 찾으면 -1)
	}{
		{"중간값 찾기", 9, 4},       // 인덱스 4
		{"작은 값 찾기", 3, 1},      // 인덱스 1
		{"큰 값 찾기", 17, 8},      // 인덱스 8
		{"시작 값 찾기", 1, 0},      // 인덱스 0
		{"끝 값 찾기", 19, 9},      // 인덱스 9
		{"없는 값 찾기", 10, -1},    // 없음
		{"범위 밖 값 찾기", 100, -1}, // 없음
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(sortedData, tc.target)
			if result != tc.expected {
				t.Errorf("실패 %s: 목표값 %d, 기대인덱스 %d, 실제값 %d",
					tc.name, tc.target, tc.expected, result)
			}
		})
	}
}
