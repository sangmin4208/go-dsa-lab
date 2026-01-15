package hashtable

import "testing"

func TestHashTable(t *testing.T) {
	t.Run("해시테이블 기본 기능 테스트", func(t *testing.T) {
		// 1. 생성
		ht := NewHashTable(100)

		// 2. Put (데이터 저장)
		ht.Put("name", "Ironman")
		ht.Put("job", "Hero")
		ht.Put("age", 40)

		// 3. Get (데이터 조회)
		val, found := ht.Get("name")
		if !found || val != "Ironman" {
			t.Errorf("Get 실패. 기대값: Ironman, 실제값: %v", val)
		}

		val, found = ht.Get("age")
		if !found || val != 40 {
			t.Errorf("Get 실패. 기대값: 40, 실제값: %v", val)
		}

		// 4. Update 테스트 (같은 키로 저장하면 덮어써야 함)
		ht.Put("job", "CEO")
		val, found = ht.Get("job")
		if !found || val != "CEO" {
			t.Errorf("Update 실패. 기대값: CEO, 실제값: %v", val)
		}

		// 5. Delete 테스트
		ht.Delete("name")
		_, found = ht.Get("name")
		if found {
			t.Errorf("Delete 실패. 데이터가 여전히 존재합니다.")
		}

		// 6. 없는 데이터 조회
		_, found = ht.Get("ghost")
		if found {
			t.Errorf("없는 데이터를 찾았다고 나옵니다.")
		}
	})
}
