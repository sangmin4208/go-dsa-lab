package sorting

import (
	"reflect" // 슬라이스 비교를 위해 사용
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "일반적인 무작위 배열",
			input:    []int{9, 4, 1, 6, 7, 3, 8, 2, 5},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "이미 정렬된 배열",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "역순 배열",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "중복 값이 있는 배열",
			input:    []int{3, 1, 2, 3, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "빈 배열",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "단일 요소 배열",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "두 요소 배열 (정렬 필요)",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "음수 포함 배열",
			input:    []int{-5, 3, -1, 7, 0},
			expected: []int{-5, -1, 0, 3, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 입력값 복사 (원본 보존을 위해)
			arr := make([]int, len(tc.input))
			copy(arr, tc.input)

			MergeSort(arr)

			// reflect.DeepEqual로 슬라이스 내용 비교
			if !reflect.DeepEqual(arr, tc.expected) {
				t.Errorf("실패 %s: 기대값 %v, 실제값 %v", tc.name, tc.expected, arr)
			}
		})
	}
}
