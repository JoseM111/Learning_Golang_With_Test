package test
/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
import "testing"

func TestRepeated(test *testing.T) {
	// ™ ____________
	repeated := Repeat("a")
	expected := "aaaaa"
	
	if repeated != expected {
		test.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(benchmark *testing.B) {
	for i := 0; i < benchmark.N; i++ {
		Repeat("a")
	}
}
/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
