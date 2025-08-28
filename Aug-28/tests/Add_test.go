package tests

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		wantError string
	}{
		{"empty string", "", 0, ""},
		{"single number", "7", 7, ""},
		{"two numbers", "1,2", 3, ""},
		{"three numbers", "5,10,15", 30, ""},
		{"newlines allowed", "1\n2,3", 6, ""},
		{"custom delimiter ;", "//;\n1;2", 3, ""},
		{"negatives error", "1,-2,3,-4", 0, "negatives not allowed: -2,-4"},
		{"ignore big numbers", "2,1001", 2, ""},
		{"custom long delimiter", "//[***]\n1***2***3", 6, ""},
		{"multiple delimiters", "//[*][%]\n1*2%3", 6, ""},
		{"multiple long delimiters", "//[**][%%]\n1**2%%3", 6, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Add(tc.input)

			if tc.wantError != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tc.wantError)
				}
				if !strings.Contains(err.Error(), tc.wantError) {
					t.Fatalf("expected error %q, got %q", tc.wantError, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Add(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
