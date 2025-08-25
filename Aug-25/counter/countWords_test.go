package counter

import "testing"

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Normal sentence", "Go is fun", 3},
		{"Empty string", "", 0},
		{"Multiple spaces", "Go   is   fun", 3},
		{"Tabs", "Go\tis\tfun", 3},
		{"Newlines", "Go\nis\nfun", 3},
		{"single word", "Hello", 1},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			actual := CountWords(testcase.input)
			if actual != testcase.expected {
				t.Errorf("Expected %d, Actual %d", testcase.expected, actual)
			}
		})
	}
}

func TestCountWords2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Normal sentence", "Go is fun", 3},
		{"Empty string", "", 0},
		{"Multiple spaces", "Go   is   fun", 3},
		{"Normal sentence with space at end", "Go is fun ", 3},
		{"Multiple spaces", "Go   is   fun    ", 3},
		{"Tabs", "Go\tis\tfun", 3},
		{"Tabs 2", "Go\tis\tfun\t", 3},
		{"Newlines", "Go\nis\nfun", 3},
		{"Newlines 2", "Go\nis\nfun\n", 3},
		{"Edge case 1", "Go\tis\tfun\\", 3},
		{"single word", "Hello", 1},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			actual := CountWords2(testcase.input)
			if actual != testcase.expected {
				t.Errorf("Expected %d, Actual %d", testcase.expected, actual)
			}
		})
	}
}

func BenchmarkCountWords(b *testing.B) {
	largeText := ""
	for i := 0; i < 100000; i++ {
		largeText += "Go is fun "
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountWords(largeText)
	}
}

func BenchmarkCountWords2(b *testing.B) {
	largeText := ""
	for i := 0; i < 100000; i++ {
		largeText += "Go is fun "
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountWords(largeText)
	}
}
