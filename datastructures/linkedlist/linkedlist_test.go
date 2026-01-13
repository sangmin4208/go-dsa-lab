package linkedlist

import (
	"testing"
)

// ToSlice도 제네릭 적용
func (l *LinkedList[T]) ToSlice() []T {
	var result []T
	curr := l.Head
	for curr != nil {
		result = append(result, curr.Data)
		curr = curr.Next
	}
	return result
}

func TestLinkedList(t *testing.T) {
	// 1. 기존 정수형 테스트 (타입 명시 [int])
	t.Run("정수형(Int) 리스트 테스트", func(t *testing.T) {
		list := LinkedList[int]{} // [int] 명시 필수!
		list.Add(10)
		list.Add(20)
		list.Remove(10)

		result := list.ToSlice()
		if len(result) != 1 || result[0] != 20 {
			t.Errorf("Int 테스트 실패. 결과: %v", result)
		}
	})

	// 2. 문자열 테스트 (새로 추가!) - 제네릭의 힘
	t.Run("문자열(String) 리스트 테스트", func(t *testing.T) {
		list := LinkedList[string]{} // [string] 타입 사용
		list.Add("Apple")
		list.Add("Banana")
		list.Add("Cherry")

		list.Remove("Banana")

		result := list.ToSlice()
		expected := []string{"Apple", "Cherry"}

		if len(result) != len(expected) {
			t.Errorf("길이 오류")
		}
		if result[0] != "Apple" || result[1] != "Cherry" {
			t.Errorf("String 테스트 실패. 결과: %v", result)
		}
	})
}
