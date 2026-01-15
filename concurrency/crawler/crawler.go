package crawler

type Fetcher interface {
	Fetch(url string) string
}

type Crawler struct{}

func NewCrawler() *Crawler {
	return &Crawler{}
}

func (c *Crawler) Crawl(urls []string, fetcher Fetcher) []string {
	var results []string

	// var wg sync.WaitGroup

	// var mu sync.Mutex

	// for _, url := range urls {
	// 	wg.Add(1)

	// 	go func(u string) {
	// 		defer wg.Done()

	// 		data := fetcher.Fetch(u)

	// 		mu.Lock()
	// 		results = append(results, data)
	// 		mu.Unlock()
	// 	}(url)
	// }
	// // 모든 작업자가 퇴근할 때까지 여기서 대기
	// wg.Wait()

	resultChan := make(chan string)

	for _, url := range urls {
		go func(u string) {
			data := fetcher.Fetch(u)
			resultChan <- data
		}(url)
	}

	// URL 개수만큼 채널에서 데이터 빼오기 (Receive)
	for range len(urls) {
		// 채널에서 데이터가 나올 때까지 대기함
		data := <-resultChan
		results = append(results, data)
	}

	return results

}
