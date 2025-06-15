package db

import "testing"

func BenchmarkCompanies(b *testing.B) {
	for range b.N {
		_ = Companies()
	}
}

func BenchmarkCloneCompanies(b *testing.B) {
	for range b.N {
		_ = CloneCompanies()
	}
}
