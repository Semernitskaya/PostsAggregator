package main

type TwitterLoader interface {
	LoadTweetsByHashTag(hashTag string, count int) []string
}

type MockTwitterLoader struct {
	tweets map[string][]string
}

func (loader MockTwitterLoader) LoadTweetsByHashTag(hashTag string, count int) []string {
	tweets := loader.tweets[hashTag]
	if len(tweets) < count {
		return tweets
	}
	return tweets[0:count]
}

func New() *MockTwitterLoader {
	return &MockTwitterLoader{map[string][]string{
		"#java":   {"j1", "j2"},
		"#golang": {"g1", "g2", "g3"},
		"#devops": {"", ""},
		"#else":   {"e1", "e2"},
	}}
}
