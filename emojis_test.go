package emojis

import "testing"

func BenchmarkLoadFromWebsite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadFromWebsite()
	}
}

func BenchmarkLoadFromApi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadFromApi()
	}
}
