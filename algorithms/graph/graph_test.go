package graph

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	// 1. 그래프 생성 및 간선 연결
	g := NewGraph()
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")

	// 2. BFS 테스트
	// A에서 시작하면: A -> B, C (순서 무관하지만 여기선 입력순) -> D, E
	t.Run("BFS 테스트", func(t *testing.T) {
		expected := []string{"A", "B", "C", "D", "E"}
		result := g.BFS("A")

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("BFS 순서 불일치. 기대값: %v, 실제값: %v", expected, result)
		}
	})

	// 3. DFS 테스트
	// A에서 시작 -> B로 감 -> B의 이웃 D로 감 -> 끝났으니 돌아와서 C -> E
	t.Run("DFS 테스트", func(t *testing.T) {
		expected := []string{"A", "B", "D", "C", "E"}
		result := g.DFS("A")

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DFS 순서 불일치. 기대값: %v, 실제값: %v", expected, result)
		}
	})
}
