package iteration

import "testing"

func Test(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("Expected %s, Got %s", expected, repeated)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++{
		Repeat("a", 3)
	}
}