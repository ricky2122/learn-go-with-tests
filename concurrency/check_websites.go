package concurrency

type (
	WebsiteChecker func(string) bool
	result         struct {
		string
		bool
	}
)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
