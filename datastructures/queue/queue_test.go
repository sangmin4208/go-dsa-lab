package queue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("Queue[string] 기능 테스트", func(t *testing.T) {
		q := Queue[string]{}

		if !q.IsEmpty() {
			t.Errorf("초기 큐는 비어있어야 합니다.")
		}

		q.Enqueue("A")
		q.Enqueue("B")
		q.Enqueue("C")

		if q.Size() != 3 {
			t.Errorf("크기가 다릅니다. 기대값: 3, 실제값: %d", q.Size())
		}

		// 3. Peek (맨 앞은 A여야 함)
		val, err := q.Peek()
		if err != nil || val != "A" {
			t.Errorf("Peek 실패. 기대값: A, 실제값: %v", val)
		}

		// 4. Dequeue (들어간 순서대로 A, B, C 나와야 함)
		val, err = q.Dequeue() // A 나옴
		if err != nil || val != "A" {
			t.Errorf("Dequeue 실패. 기대값: A, 실제값: %v", val)
		}

		val, err = q.Dequeue() // B 나옴
		if err != nil || val != "B" {
			t.Errorf("Dequeue 실패. 기대값: B, 실제값: %v", val)
		}

		val, err = q.Dequeue() // C 나옴
		if err != nil || val != "C" {
			t.Errorf("Dequeue 실패. 기대값: C, 실제값: %v", val)
		}

		// 5. 빈 큐에서 Dequeue 에러 확인
		_, err = q.Dequeue()
		if err == nil {
			t.Errorf("빈 큐에서 Dequeue 시 에러가 발생해야 합니다.")
		}

	})

}
