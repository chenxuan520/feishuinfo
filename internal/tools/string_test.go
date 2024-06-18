package tools

import (
	"testing"
)

func TestExtractFirstWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "test1",
			args:  args{s: "hello world-you ok"},
			want:  "hello",
			want1: "world-you ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExtractFirstWord(tt.args.s)
			if got != tt.want {
				t.Errorf("ExtractFirstWord() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractFirstWord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReplaceSpecialChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "hello\nworld"},
			want: "hello\\nworld",
		},
		{
			name: "test2",
			args: args{s: "hello\"world"},
			want: "hello\\\"world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpecialChar(tt.args.s); got != tt.want {
				t.Errorf("ReplaceSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
