package hashtable

import "fmt"

type HashNode struct {
	Key   string
	Value any
	Next  *HashNode
}

type HashTable struct {
	// [수정 1] 배열([100])이 아니라 슬라이스([])로 변경
	buckets []*HashNode
}

// [수정 2] size를 받아서 실제로 사용
func NewHashTable(size int) *HashTable {
	return &HashTable{
		// make(타입, 길이)를 사용해 동적으로 공간 생성
		buckets: make([]*HashNode, size),
	}
}

func (h *HashTable) Hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	// 고정 상수 대신 현재 버킷의 길이를 사용
	return sum % len(h.buckets)
}

func (h *HashTable) Put(key string, value any) {
	index := h.Hash(key)
	node := h.buckets[index]

	if node == nil {
		h.buckets[index] = &HashNode{Key: key, Value: value}
		return
	}

	prev := node
	for node != nil {
		if node.Key == key {
			node.Value = value
			return
		}
		prev = node
		node = node.Next
	}

	prev.Next = &HashNode{Key: key, Value: value}
}

func (h *HashTable) Get(key string) (any, bool) {
	index := h.Hash(key)
	node := h.buckets[index]

	for node != nil {
		if node.Key == key {
			return node.Value, true // 찾음
		}
		node = node.Next
	}

	return nil, false // 못찾음
}

func (h *HashTable) Delete(key string) {
	index := h.Hash(key)
	node := h.buckets[index]

	if node == nil {
		return
	}

	// 첫 번째 노드가 삭제 대상인 경우
	if node.Key == key {
		h.buckets[index] = node.Next
		return
	}

	// 두 번째 노드부터 검색
	prev := node
	for node != nil {
		if node.Key == key {
			prev.Next = node.Next // 연결 끊기
			return
		}
		prev = node
		node = node.Next
	}
}

// Debug: 현재 해시 테이블 상태 출력 (학습용)
func (h *HashTable) Print() {
	for i, node := range h.buckets {
		if node != nil {
			fmt.Printf("Bucket [%d]: ", i)
			for node != nil {
				fmt.Printf("{%s: %v} -> ", node.Key, node.Value)
				node = node.Next
			}
			fmt.Println("nil")
		}
	}
}
