package encoder

import (
	"testing"

	"decoder/decoder"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "no repetition passthrough",
			input: "ABC",
			want:  "ABC",
		},
		{
			name:  "single repeated character",
			input: "AAAAA",
			want:  "[5 A]",
		},
		{
			name:  "mixed with repeated chars",
			input: "#####",
			want:  "[5 #]",
		},
		{
			name:  "alternating two-char pattern",
			input: "-_-_-_-_-_",
			want:  "[5 -_]",
		},
		{
			name:  "leftover after pattern",
			input: "-_-_-_-_-_-",
			want:  "[5 -_]-",
		},
		{
			name:  "single character no repeat",
			input: "X",
			want:  "X",
		},
		{
			name:  "two same characters",
			input: "AA",
			want:  "[2 A]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.input)
			if got != tt.want {
				t.Errorf("Encode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestEncodeRoundTrip(t *testing.T) {
	inputs := []string{
		"AAAAA",
		"#####-_-_-_-_-_-#####",
		"hello world",
		"ABABABAB",
		"AAABBBCCC",
	}
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			encoded := Encode(input)
			decoded, err := decoder.Decode(encoded)
			if err != nil {
				t.Errorf("round-trip Decode(%q) error: %v (encoded from %q)", encoded, err, input)
				return
			}
			if decoded != input {
				t.Errorf("round-trip: Encode(%q)=%q, Decode=%q, want original", input, encoded, decoded)
			}
		})
	}
}

func TestEncodeLines(t *testing.T) {
	lines := []string{"AAA", "XYZ", "BCBC"}
	got := EncodeLines(lines)
	want := []string{"[3 A]", "XYZ", "[2 BC]"}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("EncodeLines()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
