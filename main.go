package main

import (
	"fmt"
	"go-dsa-lab/datastructures/hashtable"
)

func main() {
	fmt.Println("🚀 Go 해시 테이블 충돌 테스트")
	ht := hashtable.NewHashTable(100)

	// 해시 함수가 단순해서 충돌이 잘 일어납니다.
	// 예: "ABC" (65+66+67 = 198 % 100 = 98)
	//     "CBA" (67+66+65 = 198 % 100 = 98) -> 같은 98번 방에 들어감!

	fmt.Println("데이터 저장 중...")
	ht.Put("ABC", "첫번째 값")
	ht.Put("CBA", "두번째 값 (충돌!)")
	ht.Put("BCA", "세번째 값 (또 충돌!)")
	ht.Put("Simple", "충돌 없음")

	ht.Print()
}
