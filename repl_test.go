package main

import (
	"reflect"
	"testing"
)

func Test_cleanInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{
			name:     "Test tokenize and strip whitespace",
			args:     args{input: "  hello  world  "},
			expected: []string{"hello", "world"},
		},
		{
			name:     "Convert tokens to lower case",
			args:     args{input: "Pikachu BULBASAUR CharMander"},
			expected: []string{"pikachu", "bulbasaur", "charmander"},
		},
		{
			name:     "Test spaces is empty slice",
			args:     args{input: "     "},
			expected: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanInput(tt.args.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("cleanInput() = %v, want %v", got, tt.expected)
			}
		})
	}
}
