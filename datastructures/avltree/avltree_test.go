package avltree

import (
	"testing"
)

func TestAVLTree(t *testing.T) {

	t.Run("기본 삽입 및 검색 테스트", func(t *testing.T) {
		tree := New()

		tree.Insert(10)
		tree.Insert(20)
		tree.Insert(30)

		if !tree.Search(10) {
			t.Errorf("값 10을 찾을 수 없습니다.")
		}
		if !tree.Search(20) {
			t.Errorf("값 20을 찾을 수 없습니다.")
		}
		if !tree.Search(30) {
			t.Errorf("값 30을 찾을 수 없습니다.")
		}
		if tree.Search(40) {
			t.Errorf("존재하지 않는 값 40이 검색되었습니다.")
		}
	})

	t.Run("중위 순회 테스트", func(t *testing.T) {
		tree := New()

		values := []int{50, 30, 70, 20, 40, 60, 80}
		for _, v := range values {
			tree.Insert(v)
		}

		result := tree.InOrderTraversal()
		expected := []int{20, 30, 40, 50, 60, 70, 80}

		if len(result) != len(expected) {
			t.Errorf("중위 순회 결과 길이가 다릅니다. 기대값: %d, 실제값: %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("중위 순회 결과가 다릅니다. 인덱스 %d에서 기대값: %d, 실제값: %d", i, v, result[i])
			}
		}
	})

	t.Run("Right Rotation 테스트 (Left Left Case)", func(t *testing.T) {
		tree := New()

		// Left Left Case를 유발하는 삽입 순서
		tree.Insert(30)
		tree.Insert(20)
		tree.Insert(10)

		if tree.Root.Value != 20 {
			t.Errorf("Right Rotation 후 루트가 20이어야 합니다. 실제값: %d", tree.Root.Value)
		}

		if tree.Root.Left.Value != 10 {
			t.Errorf("Right Rotation 후 왼쪽 자식이 10이어야 합니다. 실제값: %d", tree.Root.Left.Value)
		}

		if tree.Root.Right.Value != 30 {
			t.Errorf("Right Rotation 후 오른쪽 자식이 30이어야 합니다. 실제값: %d", tree.Root.Right.Value)
		}

		if !tree.IsBalanced() {
			t.Errorf("트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("Left Rotation 테스트 (Right Right Case)", func(t *testing.T) {
		tree := New()

		// Right Right Case를 유발하는 삽입 순서
		tree.Insert(10)
		tree.Insert(20)
		tree.Insert(30)

		if tree.Root.Value != 20 {
			t.Errorf("Left Rotation 후 루트가 20이어야 합니다. 실제값: %d", tree.Root.Value)
		}

		if tree.Root.Left.Value != 10 {
			t.Errorf("Left Rotation 후 왼쪽 자식이 10이어야 합니다. 실제값: %d", tree.Root.Left.Value)
		}

		if tree.Root.Right.Value != 30 {
			t.Errorf("Left Rotation 후 오른쪽 자식이 30이어야 합니다. 실제값: %d", tree.Root.Right.Value)
		}

		if !tree.IsBalanced() {
			t.Errorf("트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("Left-Right Rotation 테스트", func(t *testing.T) {
		tree := New()

		// Left Right Case를 유발하는 삽입 순서
		tree.Insert(30)
		tree.Insert(10)
		tree.Insert(20)

		if tree.Root.Value != 20 {
			t.Errorf("Left-Right Rotation 후 루트가 20이어야 합니다. 실제값: %d", tree.Root.Value)
		}

		if !tree.IsBalanced() {
			t.Errorf("트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("Right-Left Rotation 테스트", func(t *testing.T) {
		tree := New()

		// Right Left Case를 유발하는 삽입 순서
		tree.Insert(10)
		tree.Insert(30)
		tree.Insert(20)

		if tree.Root.Value != 20 {
			t.Errorf("Right-Left Rotation 후 루트가 20이어야 합니다. 실제값: %d", tree.Root.Value)
		}

		if !tree.IsBalanced() {
			t.Errorf("트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("삭제 테스트 - 리프 노드", func(t *testing.T) {
		tree := New()

		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)

		tree.Delete(30)

		if tree.Search(30) {
			t.Errorf("값 30이 삭제되지 않았습니다.")
		}

		if !tree.Search(50) || !tree.Search(70) {
			t.Errorf("다른 값들이 유지되어야 합니다.")
		}

		if !tree.IsBalanced() {
			t.Errorf("삭제 후 트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("삭제 테스트 - 자식이 하나인 노드", func(t *testing.T) {
		tree := New()

		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(20)

		tree.Delete(30)

		if tree.Search(30) {
			t.Errorf("값 30이 삭제되지 않았습니다.")
		}

		if !tree.Search(50) || !tree.Search(20) {
			t.Errorf("다른 값들이 유지되어야 합니다.")
		}

		if !tree.IsBalanced() {
			t.Errorf("삭제 후 트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("삭제 테스트 - 자식이 둘인 노드", func(t *testing.T) {
		tree := New()

		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)
		tree.Insert(20)
		tree.Insert(40)

		tree.Delete(30)

		if tree.Search(30) {
			t.Errorf("값 30이 삭제되지 않았습니다.")
		}

		if !tree.Search(50) || !tree.Search(70) || !tree.Search(20) || !tree.Search(40) {
			t.Errorf("다른 값들이 유지되어야 합니다.")
		}

		if !tree.IsBalanced() {
			t.Errorf("삭제 후 트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("삭제 후 재균형 테스트", func(t *testing.T) {
		tree := New()

		// 균형 잡힌 트리 생성
		values := []int{50, 25, 75, 10, 30, 60, 80, 5, 15}
		for _, v := range values {
			tree.Insert(v)
		}

		// 삭제로 불균형 유발
		tree.Delete(80)
		tree.Delete(75)

		if !tree.IsBalanced() {
			t.Errorf("삭제 후 트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("대량 삽입 및 균형 유지 테스트", func(t *testing.T) {
		tree := New()

		// 순차적으로 값 삽입
		for i := 1; i <= 100; i++ {
			tree.Insert(i)
		}

		// 모든 값 검색 가능 확인
		for i := 1; i <= 100; i++ {
			if !tree.Search(i) {
				t.Errorf("값 %d를 찾을 수 없습니다.", i)
			}
		}

		// 균형 확인
		if !tree.IsBalanced() {
			t.Errorf("대량 삽입 후 트리가 균형을 유지해야 합니다.")
		}

		// 높이 확인 (균형 잡힌 트리의 높이는 log(n)에 가까워야 함)
		height := tree.GetHeight()
		if height > 10 {
			t.Errorf("100개 노드에 대한 트리 높이가 너무 큽니다. 높이: %d", height)
		}
	})

	t.Run("중복 값 삽입 테스트", func(t *testing.T) {
		tree := New()

		tree.Insert(50)
		tree.Insert(50)

		result := tree.InOrderTraversal()
		if len(result) != 1 {
			t.Errorf("중복 값은 무시되어야 합니다. 노드 개수: %d", len(result))
		}
	})

	t.Run("존재하지 않는 값 삭제 테스트", func(t *testing.T) {
		tree := New()

		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)

		tree.Delete(100)

		result := tree.InOrderTraversal()
		expected := []int{30, 50, 70}

		if len(result) != len(expected) {
			t.Errorf("노드 개수가 변하지 않아야 합니다. 기대값: %d, 실제값: %d", len(expected), len(result))
		}
	})

	t.Run("빈 트리에서 검색 테스트", func(t *testing.T) {
		tree := New()

		if tree.Search(10) {
			t.Errorf("빈 트리에서는 어떤 값도 찾을 수 없어야 합니다.")
		}
	})

	t.Run("빈 트리에서 삭제 테스트", func(t *testing.T) {
		tree := New()

		tree.Delete(10)

		if tree.Root != nil {
			t.Errorf("빈 트리는 여전히 비어있어야 합니다.")
		}
	})

	t.Run("루트 노드 삭제 테스트", func(t *testing.T) {
		tree := New()

		tree.Insert(50)
		tree.Insert(30)
		tree.Insert(70)

		tree.Delete(50)

		if tree.Search(50) {
			t.Errorf("루트 노드가 삭제되지 않았습니다.")
		}

		if !tree.Search(30) || !tree.Search(70) {
			t.Errorf("다른 노드들이 유지되어야 합니다.")
		}

		if !tree.IsBalanced() {
			t.Errorf("루트 삭제 후 트리가 균형을 유지해야 합니다.")
		}
	})

	t.Run("모든 노드 삭제 테스트", func(t *testing.T) {
		tree := New()

		values := []int{50, 30, 70, 20, 40}
		for _, v := range values {
			tree.Insert(v)
		}

		for _, v := range values {
			tree.Delete(v)
		}

		if tree.Root != nil {
			t.Errorf("모든 노드 삭제 후 트리는 비어있어야 합니다.")
		}

		result := tree.InOrderTraversal()
		if len(result) != 0 {
			t.Errorf("모든 노드 삭제 후 중위 순회 결과는 비어있어야 합니다.")
		}
	})
}
