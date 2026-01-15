package graph

// Graph: 문자열 키를 가진 노드들의 연결 관계를 저장
type Graph struct {
	AdjList map[string][]string
}

// NewGraph: 생성자
func NewGraph() *Graph {
	return &Graph{
		AdjList: make(map[string][]string),
	}
}

// AddEdge: 양방향 연결 (Undirected Graph)
func (g *Graph) AddEdge(src, dest string) {
	g.AdjList[src] = append(g.AdjList[src], dest)
	g.AdjList[dest] = append(g.AdjList[dest], src)
}

func (g *Graph) BFS(startNode string) []string {
	visited := make(map[string]bool) // 방문 기록
	var result []string

	// 큐 생성 (Go Slice 사용)
	queue := []string{startNode}
	visited[startNode] = true

	for len(queue) > 0 {
		// 1. Dequeue (맨 앞 꺼내기)
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// 2. 현재 노드의 이웃들을 본다
		neighbors := g.AdjList[current]

		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor) // Enqueue
			}
		}
	}
	return result
}

// DFS (깊이 우선 탐색): 재귀(Recursion) 사용
func (g *Graph) DFS(startNode string) []string {
	visited := make(map[string]bool)
	var result []string

	var dfsHelper func(node string)
	dfsHelper = func(node string) {
		current := node
		result = append(result, current)
		visited[current] = true

		neighbors := g.AdjList[current]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				dfsHelper(neighbor)
			}
		}
	}

	dfsHelper(startNode)
	return result
}
