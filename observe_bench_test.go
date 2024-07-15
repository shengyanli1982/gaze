package gaze

import "testing"

func BenchmarkReactiveValue_IntGet(b *testing.B) {
	ov := NewNopReactiveValue(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ov.Get()
	}
}

func BenchmarkReactiveValue_IntSet(b *testing.B) {
	ov := NewNopReactiveValue(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ov.Set(i)
	}
}

func BenchmarkReactiveValue_IntSetGet(b *testing.B) {
	ov := NewNopReactiveValue(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ov.Set(i)
		ov.Get()
	}
}

func BenchmarkStd_IntSet(b *testing.B) {
	v := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = i
	}
	_ = v
}

func BenchmarkStd_IntGet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = i
	}
}

func BenchmarkStd_IntSetGet(b *testing.B) {
	value := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value = i
		_ = value
	}
}
