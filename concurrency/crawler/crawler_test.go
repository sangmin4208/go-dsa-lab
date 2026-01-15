package crawler

import (
	"testing"
	"time"
)

// MockFetcher: 실제 네트워크 대신 잠시 멈췄다가 가짜 데이터를 주는 녀석
type MockFetcher struct{}

func (f *MockFetcher) Fetch(url string) string {
	time.Sleep(100 * time.Millisecond) // 0.1초 지연 (네트워크 대기 시뮬레이션)
	return "Result of " + url
}

func TestCrawl(t *testing.T) {
	urls := []string{
		"http://google.com",
		"http://naver.com",
		"http://daum.net",
		"http://github.com",
		"http://stackoverflow.com",
	}

	fetcher := &MockFetcher{}

	crawler := NewCrawler()

	// 1. 시작 시간 기록
	start := time.Now()

	// 2. 크롤링 실행 (우리가 만들 함수)
	results := crawler.Crawl(urls, fetcher)

	// 3. 걸린 시간 측정
	elapsed := time.Since(start)

	// 검증 1: 결과 개수
	if len(results) != len(urls) {
		t.Errorf("결과 개수가 다릅니다. 기대값: %d, 실제값: %d", len(urls), len(results))
	}

	// 검증 2: 성능 (동시성 체크)
	// 5개 * 0.1초 = 0.5초.
	// 동시성이 제대로 작동하면 이론상 0.1초 + 알파만 걸려야 함.
	// 넉넉하게 0.2초보다 적게 걸리면 성공으로 간주.
	if elapsed >= 250*time.Millisecond {
		t.Errorf("너무 느립니다! 동시성 처리가 안 된 것 같아요. 걸린 시간: %v", elapsed)
	} else {
		t.Logf("성공! 걸린 시간: %v (매우 빠름)", elapsed)
	}
}
