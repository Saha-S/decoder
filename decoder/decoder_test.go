package decoder

import (
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "plain text passthrough",
			input: "ABC",
			want:  "ABC",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "single bracket group",
			input: "[3 A]",
			want:  "AAA",
		},
		{
			name:  "multiple bracket groups",
			input: "[3 A][2 B]",
			want:  "AAABB",
		},
		{
			name:  "mixed text and brackets",
			input: "ABC[10 D]EFG",
			want:  "ABCDDDDDDDDDDEFG",
		},
		{
			name:  "multi-character repeat string",
			input: "[5 -_]",
			want:  "-_-_-_-_-_",
		},
		{
			name:  "repeat count 1",
			input: "[1 X]",
			want:  "X",
		},
		{
			name:  "repeat with spaces in str",
			input: "[3 a b]",
			want:  "a ba ba b",
		},
		{
			name:  "large repeat",
			input: "[5 #][5 -_]-[5 #]",
			want:  "#####-_-_-_-_-_-#####",
		},
		{
			name:    "unbalanced bracket - missing close",
			input:   "[5 #",
			wantErr: true,
		},
		{
			name:    "unbalanced bracket - unexpected close",
			input:   "5 #]",
			wantErr: true,
		},
		{
			name:    "non-numeric count",
			input:   "[abc #]",
			wantErr: true,
		},
		{
			name:    "zero count",
			input:   "[0 #]",
			wantErr: true,
		},
		{
			name:    "negative count",
			input:   "[-1 #]",
			wantErr: true,
		},
		{
			name:    "empty repeat string",
			input:   "[5 ]",
			wantErr: true,
		},
		{
			name:    "missing space separator",
			input:   "[5#]",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Decode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDecodeLines(t *testing.T) {
	lines := []string{"[3 A]", "hello", "[2 BC]"}
	got, err := DecodeLines(lines)
	if err != nil {
		t.Fatalf("DecodeLines() unexpected error: %v", err)
	}
	want := []string{"AAA", "hello", "BCBC"}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("DecodeLines()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestDecodeLinesError(t *testing.T) {
	lines := []string{"[3 A]", "[bad]", "[2 BC]"}
	_, err := DecodeLines(lines)
	if err == nil {
		t.Error("DecodeLines() expected error for invalid line, got nil")
	}
}
