package avltree

// Node는 AVL Tree의 노드를 나타냅니다
type Node struct {
	Value  int
	Height int
	Left   *Node
	Right  *Node
}

// AVLTree는 자가 균형 이진 탐색 트리입니다
type AVLTree struct {
	Root *Node
}

// New는 새로운 AVL Tree를 생성합니다
func New() *AVLTree {
	return &AVLTree{}
}

// height는 노드의 높이를 반환합니다
func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// max는 두 정수 중 큰 값을 반환합니다
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// updateHeight는 노드의 높이를 갱신합니다
func updateHeight(n *Node) {
	if n != nil {
		n.Height = 1 + max(height(n.Left), height(n.Right))
	}
}

// getBalance는 노드의 균형 인수를 반환합니다
func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

// rightRotate는 우회전을 수행합니다
func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	// 회전 수행
	x.Right = y
	y.Left = T2

	// 높이 갱신
	updateHeight(y)
	updateHeight(x)

	return x
}

// leftRotate는 좌회전을 수행합니다
func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	// 회전 수행
	y.Left = x
	x.Right = T2

	// 높이 갱신
	updateHeight(x)
	updateHeight(y)

	return y
}

// Insert는 값을 AVL Tree에 삽입합니다
func (t *AVLTree) Insert(value int) {
	t.Root = insert(t.Root, value)
}

// insert는 재귀적으로 노드를 삽입하고 균형을 유지합니다
func insert(node *Node, value int) *Node {
	// 1. 일반적인 BST 삽입
	if node == nil {
		return &Node{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	} else {
		// 중복된 값은 허용하지 않음
		return node
	}

	// 2. 노드의 높이 갱신
	updateHeight(node)

	// 3. 균형 인수 확인
	balance := getBalance(node)

	// 4. 불균형 처리

	// Left Left Case
	if balance > 1 && value < node.Left.Value {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && value > node.Right.Value {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && value > node.Left.Value {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && value < node.Right.Value {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// Search는 값을 AVL Tree에서 검색합니다
func (t *AVLTree) Search(value int) bool {
	return search(t.Root, value)
}

// search는 재귀적으로 값을 검색합니다
func search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	} else if value < node.Value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

// Delete는 값을 AVL Tree에서 삭제합니다
func (t *AVLTree) Delete(value int) {
	t.Root = delete(t.Root, value)
}

// delete는 재귀적으로 노드를 삭제하고 균형을 유지합니다
func delete(node *Node, value int) *Node {
	// 1. 일반적인 BST 삭제
	if node == nil {
		return node
	}

	if value < node.Value {
		node.Left = delete(node.Left, value)
	} else if value > node.Value {
		node.Right = delete(node.Right, value)
	} else {
		// 노드 발견: 삭제할 노드

		// 자식이 하나 이하인 경우
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// 자식이 둘인 경우: 오른쪽 서브트리의 최소값을 찾음
		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = delete(node.Right, minNode.Value)
	}

	// 2. 노드의 높이 갱신
	updateHeight(node)

	// 3. 균형 인수 확인
	balance := getBalance(node)

	// 4. 불균형 처리

	// Left Left Case
	if balance > 1 && getBalance(node.Left) >= 0 {
		return rightRotate(node)
	}

	// Left Right Case
	if balance > 1 && getBalance(node.Left) < 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && getBalance(node.Right) <= 0 {
		return leftRotate(node)
	}

	// Right Left Case
	if balance < -1 && getBalance(node.Right) > 0 {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// findMin은 서브트리에서 최소값을 가진 노드를 찾습니다
func findMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// InOrderTraversal은 중위 순회를 수행하여 값들을 반환합니다
func (t *AVLTree) InOrderTraversal() []int {
	var result []int
	inOrderTraversal(t.Root, &result)
	return result
}

// inOrderTraversal은 재귀적으로 중위 순회를 수행합니다
func inOrderTraversal(node *Node, result *[]int) {
	if node != nil {
		inOrderTraversal(node.Left, result)
		*result = append(*result, node.Value)
		inOrderTraversal(node.Right, result)
	}
}

// GetHeight는 트리의 높이를 반환합니다
func (t *AVLTree) GetHeight() int {
	return height(t.Root)
}

// IsBalanced는 트리가 균형 잡혀 있는지 확인합니다
func (t *AVLTree) IsBalanced() bool {
	return isBalanced(t.Root)
}

// isBalanced는 재귀적으로 모든 노드가 균형 잡혀 있는지 확인합니다
func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	balance := getBalance(node)
	if balance < -1 || balance > 1 {
		return false
	}

	return isBalanced(node.Left) && isBalanced(node.Right)
}
