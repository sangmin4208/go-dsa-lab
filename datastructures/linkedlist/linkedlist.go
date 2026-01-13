package linkedlist

import "fmt"

// Node: 데이터와 다음 노드의 주소(Next)를 가집니다.
type Node[T comparable] struct {
	Data T        // 저장할 값
	Next *Node[T] // 다음 노드를 가리키는 포인터 (이게 핵심입니다!)
}

// LinkedList: 첫 번째 노드(Head)만 알고 있으면 전체를 찾아갈 수 있습니다.
type LinkedList[T comparable] struct {
	Head *Node[T]
}

// Add: 리스트의 맨 끝에 새로운 값을 추가합니다.
// (l *LinkedList)는 리시버(Receiver)라고 하며, TypeScript의 클래스 메서드와 비슷합니다.
func (l *LinkedList[T]) Add(data T) {
	// 1. 새로운 노드를 메모리에 생성하고 그 주소(&)를 가져옵니다.
	newNode := &Node[T]{Data: data}

	// 2. 만약 리스트가 비어있다면(Head가 nil이면), 새 노드가 헤드가 됩니다.
	if l.Head == nil {
		l.Head = newNode
		return
	}

	// 3. 비어있지 않다면, 마지막 노드를 찾아야 합니다.
	current := l.Head
	for current.Next != nil { // 다음이 없을 때까지 계속 이동
		current = current.Next
	}

	// 4. 마지막 노드의 Next에 새 노드를 연결합니다.
	current.Next = newNode
}

// Remove: 리스트에서 특정 값을 삭제합니다.
func (l *LinkedList[T]) Remove(data T) {
	if l.Head == nil {
		return
	}
	// 삭제할 노드가 Head인 경우
	if l.Head.Data == data {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}

}

// PrintAll: 리스트의 모든 값을 출력합니다.
func (l *LinkedList[T]) PrintAll() {
	if l.Head == nil {
		fmt.Println("리스트가 비어있습니다.")
		return
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil") // 끝을 표시
}
