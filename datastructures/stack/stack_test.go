package stack

import "testing"

func TestStack(t *testing.T) {

	t.Run("Stack[Int] 기능 테스트", func(t *testing.T) {
		s := Stack[int]{}

		if !s.IsEmpty() {
			t.Errorf("초기 스택은 비어있어야 합니다.")
		}

		s.Push(10)
		s.Push(20)
		s.Push(30)

		if s.IsEmpty() {
			t.Errorf("스택이 비어있지 않아야 합니다.")
		}

		val, err := s.Peek()
		if err != nil || val != 30 {
			t.Errorf("Peek 실패. 기대값; 30, 실제값 %v", val)
		}

		val, err = s.Pop()
		if err != nil || val != 30 {
			t.Errorf("Pop 실패. 기대값; 30, 실제값 %v", val)
		}
		val, err = s.Pop()
		if err != nil || val != 20 {
			t.Errorf("Pop 실패. 기대값; 20, 실제값 %v", val)
		}
		val, err = s.Pop()
		if err != nil || val != 10 {
			t.Errorf("Pop 실패. 기대값; 10, 실제값 %v", val)
		}

		_, err = s.Pop()
		if err == nil {
			t.Errorf("빈 스텍에서 Pop을 했는데 에러가 발생하지 않았습니다.")
		}
	})
}
